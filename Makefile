# List of architectures to build
ARCHES := linux/amd64 linux/386 linux/arm64 linux/arm darwin/amd64 darwin/arm64 windows/386 windows/amd64 windows/arm windows/arm64

# Output directory
OUTPUT_DIR := ./build

# Binary name
BINARY_NAME := wggen
VERSION := ${VERSION_NUM}
DATE := $$(date +'%F %T %Z')
GITREVISION = $$(git show -s --format="%H" HEAD)


BUILDVARS = -X 'github.com/xaner4/wggen/cmd.version=$(VERSION)' \
			-X 'github.com/xaner4/wggen/cmd.buildDate=$(DATE)' \
			-X 'github.com/xaner4/wggen/cmd.gitRevision=$(GITREVISION)'

all: $(ARCHES)

$(ARCHES):
	go mod download
	@echo "Building $(BINARY_NAME) for $@..."
	GOARCH=$(word 2,$(subst /, ,$@)) GOOS=$(word 1,$(subst /, ,$@)) go build -ldflags="-s -w $(BUILDVARS) " -trimpath -gcflags="-trimpath -m" -o $(OUTPUT_DIR)/$(BINARY_NAME)_$(VERSION)_$(subst /,_,$@) .

clean:
	@echo "Cleaning build directory..."
	rm -rf $(OUTPUT_DIR)/$(BINARY_NAME)*

.PHONY: all $(ARCHES) clean
