# IPFS-X

A Go-based tool for IPFS operations including CID calculation, uploading and downloading.

## Features

- Calculate CID v1 for files
- Upload files to IPFS network
- Download files from IPFS network

## Requirements

- Go (version >= 1.18)
- IPFS node running on localhost:5001

## Installation

### Build from source

```bash
git clone <repository-url>
cd go
make build
```

### Install globally

```bash
cd go
make install
```

## Usage

### Calculate CID v1

```bash
# Using the binary directly
./ipfs-x cid -f /path/to/file

# Using go run
go run main.go cid -f /path/to/file

# Using make
make cid file=/path/to/file
```

### Upload to IPFS

```bash
# Using the binary directly
./ipfs-x upload -f /path/to/file

# Using go run
go run main.go upload -f /path/to/file

# Using make
make upload file=/path/to/file
```

### Download from IPFS

```bash
# Using the binary directly
./ipfs-x download -c <CID> -o /path/to/output

# Using go run
go run main.go download -c <CID> -o /path/to/output

# Using make
make download=<CID> output=/path/to/output
```

## Command Line Options

### cid command
- `-f, --file`: Path to the file (required)

### upload command
- `-f, --file`: Path to the file (required)

### download command
- `-c, --cid`: IPFS CID of the file (required)
- `-o, --output`: Output file path (required)

## Examples

1. Calculate CID for a file:
```bash
./ipfs-x cid -f example.txt
```

2. Upload a file to IPFS:
```bash
./ipfs-x upload -f example.txt
```

3. Download a file from IPFS:
```bash
./ipfs-x download -c QmXoypizjW3WknFiJnKLwHCnL72vedAjJPduzAisXMarbw -o downloaded_example.txt
```

## License

[MIT License](LICENSE)
