package cmd

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs"
	"github.com/multiformats/go-multihash"
	"github.com/rickiey/ipfs-x/utils"
	"github.com/urfave/cli/v2"
)

func calculateCID(c *cli.Context) error {
	filePath := c.String("file")
	return calculateCIDv1(filePath)
}

func calculateCIDv1(filePath string) error {
	// Read file content
	data, err := utils.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// Create a PBNode
	pbNode := unixfs.FSNodeFromBytes(data)
	dagNode := merkledag.NewRawNode(pbNode.Data())

	// Calculate CID v1
	buidler := cid.V1Builder{Codec: cid.DagProtobuf, MhType: multihash.SHA2_256}
	cidV1, err := buidler.Sum(dagNode.RawData())
	if err != nil {
		return fmt.Errorf("failed to calculate CID: %v", err)
	}

	fmt.Printf("CID v1: %s", cidV1.String())
	return nil
}
