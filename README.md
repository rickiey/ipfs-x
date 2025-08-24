# ipfs-x

[中文](README-zh.md)

Using the IPFS protocol, upload files to the IPFS network and return the IPFS hash value for file sharing.
It also interacts with various storage systems, including but not limited to:

+ Filecoin
+ Swarm
+ Pinata
+ Arweave
+ Walrus

Before running or building this project, please ensure the following conditions are met:

Install Go (version >= 1.18)
Install IPFS and ensure the IPFS node is running
(Optional) Configure access credentials for other storage systems (such as Pinata API keys, Walrus accounts, etc.)

## Basic Features

+ Get CID for local files
+ Upload local files to the IPFS network
+ Download files from the IPFS network

## Interactive Features

+ Upload files to other storage systems with IPFS CID

### Upload to Walrus

1. Get the file CID
2. Upload the file to Walrus with metadata {cid: CID}