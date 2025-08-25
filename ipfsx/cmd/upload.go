package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/urfave/cli/v2"
)

func uploadToIPFS(c *cli.Context) error {
	filePath := c.String("file")
	recursive := c.Bool("recursive")
	return uploadFileToIPFS(filePath, recursive)
}

func uploadFileToIPFS(filePath string, recursive bool) error {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	// Get file info
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	// Connect to IPFS
	sh := shell.NewShell("localhost:5001")

	// Test connection
	if _, err := sh.ID(); err != nil {
		return fmt.Errorf("failed to connect to IPFS daemon: %v", err)
	}

	var cid string

	if fileInfo.IsDir() {
		if !recursive {
			return fmt.Errorf("path is a directory, use --recursive flag to upload directories")
		}

		// Upload directory
		cid, err = sh.AddDir(filePath)
		if err != nil {
			return fmt.Errorf("failed to upload directory to IPFS: %v", err)
		}
		fmt.Printf("Directory uploaded. CID: %s\n", cid)
	} else {
		// Upload single file
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file: %v", err)
		}
		defer file.Close()

		// Upload file with options
		cid, err = sh.Add(
			file,
			shell.Pin(true),           // Pin the file
			shell.OnlyHash(false),     // Actually upload, not just calculate hash
			shell.CidVersion(1),       // Use CID v1
		)
		if err != nil {
			return fmt.Errorf("failed to upload file to IPFS: %v", err)
		}
		fmt.Printf("File uploaded. CID: %s\n", cid)
	}

	// Get file size
	size, err := getFileSize(filePath)
	if err != nil {
		fmt.Printf("Warning: could not get file size: %v\n", err)
	} else {
		fmt.Printf("File size: %d bytes\n", size)
	}

	return nil
}

// getFileSize calculates the total size of a file or directory
func getFileSize(path string) (int64, error) {
	var size int64

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	return size, err
}

// uploadWithProgress uploads file with progress indication
func uploadWithProgress(sh *shell.Shell, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Get file size for progress
	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %v", err)
	}

	// Create progress reader
	progressReader := &progressReader{
		reader:   file,
		total:    fileInfo.Size(),
		callback: func(bytes int64) { fmt.Printf("\rUploading: %d/%d bytes", bytes, fileInfo.Size()) },
	}

	// Upload with progress
	cid, err := sh.Add(
		progressReader,
		shell.Pin(true),
		shell.OnlyHash(false),
		shell.CidVersion(1),
	)
	if err != nil {
		return "", err
	}

	fmt.Println() // New line after progress
	return cid, nil
}

// progressReader wraps an io.Reader to track progress
type progressReader struct {
	reader   io.Reader
	total    int64
	read     int64
	callback func(int64)
}

func (pr *progressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	pr.read += int64(n)
	if pr.callback != nil {
		pr.callback(pr.read)
	}
	return n, err
}
