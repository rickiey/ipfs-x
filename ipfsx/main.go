package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-ipfs-api"
	"github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs"
	"github.com/ipfs/kubo/plugin/loader"
	mh "github.com/multiformats/go-multihash"
	"github.com/urfave/cli/v2"
)

func init() {
	// Load plugins
	plugins, err := loader.NewPluginLoader(".")
	if err != nil {
		log.Panicf("error loading plugins: %s", err)
	}
	if err := plugins.Initialize(); err != nil {
		log.Panicf("error initializing plugins: %s", err)
	}
	if err := plugins.Inject(); err != nil {
		log.Panicf("error injecting plugins: %s", err)
	}
}

func main() {
	app := &cli.App{
		Name:  "ipfs-x",
		Usage: "A tool for IPFS operations including CID calculation, uploading and downloading",
		Commands: []*cli.Command{
			{
				Name:  "cid",
				Usage: "Calculate CID v1 for a file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "Path to the file",
						Required: true,
					},
				},
				Action: calculateCID,
			},
			{
				Name:  "upload",
				Usage: "Upload a file to IPFS network",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "Path to the file",
						Required: true,
					},
				},
				Action: uploadToIPFS,
			},
			{
				Name:  "download",
				Usage: "Download a file from IPFS network",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "cid",
						Aliases:  []string{"c"},
						Usage:    "IPFS CID of the file",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "output",
						Aliases:  []string{"o"},
						Usage:    "Output file path",
						Required: true,
					},
				},
				Action: downloadFromIPFS,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// calculateCID calculates CID v1 for a file
func calculateCID(c *cli.Context) error {
	filePath := c.String("file")
	return calculateCIDv1(filePath)
}

func calculateCIDv1(filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Read file content
	reader := bufio.NewReader(file)
	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// Create a PBNode
	pbNode, err := unixfs.FSNodeFromBytes(data)
	if err != nil {
		return fmt.Errorf("failed to create PBNode: %v", err)
	}
	dagNode := merkledag.NewRawNode(pbNode.Data())

	// Calculate CID v1
	buidler := cid.V1Builder{Codec: cid.DagProtobuf, MhType: mh.SHA2_256}
	cidV1, err := buidler.Sum(dagNode.RawData())
	if err != nil {
		return fmt.Errorf("failed to calculate CID: %v", err)
	}

	fmt.Printf("CID v1: %s", cidV1.String())
	return nil
}

// uploadToIPFS uploads a file to IPFS network
func uploadToIPFS(c *cli.Context) error {
	filePath := c.String("file")
	return uploadFileToIPFS(filePath)
}

func uploadFileToIPFS(filePath string) error {
	// Connect to IPFS
	sh := shell.NewShell("localhost:5001")

	// Upload file
	cid, err := sh.AddDir(filePath)
	if err != nil {
		return fmt.Errorf("failed to upload file to IPFS: %v", err)
	}

	fmt.Printf("File uploaded. CID: %s", cid)
	return nil
}

// downloadFromIPFS downloads a file from IPFS network
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
