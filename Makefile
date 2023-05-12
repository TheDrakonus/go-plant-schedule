# Define the name of the output binary
BINARY_NAME := go-plant-scheduler

# Define the directory for the output binary
BUILD_DIR := ./build

# Define the Go compiler command
GO := go

# Define the Go linker flags
LDFLAGS := -w -s

# Define the build command
BUILD_CMD := cd api && $(GO) build -ldflags "$(LDFLAGS)" -o ../$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: all build clean

all: clean build

build:
	mkdir -p $(BUILD_DIR)
	$(BUILD_CMD)

clean:
	rm -rf $(BUILD_DIR)
