package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ipfs/kubo/plugin/loader"
	"github.com/urfave/cli/v2"
)

// init initializes IPFS plugins
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

// New creates a new CLI application
func New() *cli.App {
	return &cli.App{
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
					&cli.BoolFlag{
						Name:     "recursive",
						Aliases:  []string{"r"},
						Usage:    "Upload recursively",
						Required: false,
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
}

// Run runs the CLI application
func Run() {
	if err := New().Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
