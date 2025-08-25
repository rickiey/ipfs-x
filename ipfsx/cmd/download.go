package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/urfave/cli/v2"
)

func downloadFromIPFS(c *cli.Context) error {
	cidStr := c.String("cid")
	outputPath := c.String("output")
	return downloadFileFromIPFS(cidStr, outputPath)
}

func downloadFileFromIPFS(cidStr, outputPath string) error {
	// Connect to IPFS
	sh := shell.NewShell("localhost:5001")

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Download file
	if err := sh.Get(cidStr, outputPath); err != nil {
		return fmt.Errorf("failed to download file from IPFS: %v", err)
	}

	fmt.Printf("File downloaded to: %s", outputPath)
	return nil
}
