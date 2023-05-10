# Build go module into a binary
# Usage: make build


build/go-plant-schedule: main.go
	go build -o build/go-plant-schedule main.go

build: build/go-plant-schedule
.PHONY: build

clean:
	rm -rf build
.PHONY: clean