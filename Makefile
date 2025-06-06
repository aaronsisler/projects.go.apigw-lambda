# Set Go environment
GOOS=linux
GOARCH=amd64

# Target to build handler1 Lambda binary
build-handler-hello-get: GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/handler_hello_get/bootstrap ./cmd/hello/get

zip-handler-hello-get: zip handler_hello_get.zip -j bin/handler_hello_get/bootstrap

# # Target to build handler2 Lambda binary
# build-handler2:
# 	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/handler2 ./cmd/handler2

# Build all handlers
build: build-handler-hello-get
zip: zip-handler-hello-get

# Run all tests
# test:
	# go test ./...

# Format code
# fmt:
	# go fmt ./...

# Clean output binaries
# clean:
	# rm -rf bin/
