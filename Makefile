REPO_ROOT ?= $(shell pwd)

all: clean build

all_debug: clean build_debug_service run-debug

rebuild:
	rm -rf bin/
	mkdir bin
	build

build:
	rm -rf bin/task_queue_system
	# dep ensure
	go mod download && go mod verify
	go build -a --ldflags '-extldflags "-static"' -o bin/task_queue_system ${REPO_ROOT}/cmd/main.go

build_debug_service:
	rm -rf bin/task_queue_system
	# dep ensure
	go mod download && go mod verify
	GOOS=linux go build -gcflags="all=-N -l" -o bin/task_queue_system-debug ${REPO_ROOT}/cmd/main.go

# Run lint
lint:
	golangci-lint run ./cmd/...

clean:
	rm -rf bin/

run:
	./bin/task_queue_system

run_debug:
	./bin/task_queue_system-debug

test:
	go get github.com/stretchr/testify
	go test -v -cover ./...

.PHONY: clean all build run