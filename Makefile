# List of architectures to build
ARCHES := linux/amd64 linux/386 linux/arm64 linux/arm darwin/amd64 darwin/arm64 windows/386 windows/amd64 windows/arm windows/arm64

# Output directory
OUTPUT_DIR := ./build

# Binary name (change this to your binary name)
BINARY_NAME := wggen

all: $(ARCHES)

$(ARCHES):
	@echo "Building $(BINARY_NAME) for $@..."
	GOARCH=$(word 2,$(subst /, ,$@)) GOOS=$(word 1,$(subst /, ,$@)) go build -ldflags="-s -w" -trimpath -gcflags="-trimpath -m" -o $(OUTPUT_DIR)/$(BINARY_NAME)_$(subst /,_,$@) .

clean:
	@echo "Cleaning build directory..."
	rm -rf $(OUTPUT_DIR)

.PHONY: all $(ARCHES) clean
