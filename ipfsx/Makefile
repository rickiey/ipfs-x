# Go build flags
GO_BUILD=go build
GO_FLAGS=-ldflags="-s -w"

# Binary name
BINARY_NAME=ipfs-x

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	$(GO_BUILD) $(GO_FLAGS) -o $(BINARY_NAME) .

# Run the binary
.PHONY: run
run: build
	./$(BINARY_NAME)

# Install the binary
.PHONY: install
install:
	$(GO_BUILD) $(GO_FLAGS) -o $(BINARY_NAME) .
	sudo mv $(BINARY_NAME) /usr/local/bin/

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

# Calculate CID for a file
.PHONY: cid
cid:
	go run main.go cid -f $(file)

# Upload a file to IPFS
.PHONY: upload
upload:
	go run main.go upload -f $(file)

# Download a file from IPFS
.PHONY: download
download:
	go run main.go download -c $(cid) -o $(output)
