package cmd

import (
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/rickiey/ipfs-x/utils"
	"github.com/urfave/cli/v2"
)

func uploadToIPFS(c *cli.Context) error {
	filePath := c.String("file")
	return uploadFileToIPFS(filePath)
}

func uploadFileToIPFS(filePath string) error {
	// Check if file exists
	if !utils.FileExists(filePath) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	// Connect to IPFS
	sh := shell.NewShell("localhost:5001")

	// Upload file
	cid, err := sh.AddPath(filePath)
	if err != nil {
		return fmt.Errorf("failed to upload file to IPFS: %v", err)
	}

	fmt.Printf("File uploaded. CID: %s", cid)
	return nil
}
