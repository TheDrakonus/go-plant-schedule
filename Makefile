# Build go module into a binary
# Usage: make build

# Glob files in src/ 
SOURCES := $(shell find . -name '*.go' -path "./*")

build/go-plant-schedule:
	go build -o build/go-plant-schedule .

build: build/go-plant-schedule
.PHONY: build

clean:
	rm -rf build
.PHONY: clean