# Set Go environment
GOOS=linux
GOARCH=amd64

# Building the handlers
build-handler-hello-get:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/handler_hello_get/bootstrap ./cmd/hello/get

build-handler-hello-post:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/handler_hello_post/bootstrap ./cmd/hello/post

# Ziping the handlers
zip-handler-hello-get:
	zip handler_hello_get.zip -j bin/handler_hello_get/bootstrap

zip-handler-hello-post:
	zip handler_hello_post.zip -j bin/handler_hello_post/bootstrap

# Build all handlers
build:
  build-handler-hello-get \
  build-handler-hello-post

# Zip all handlers
zip:
  zip-handler-hello-get \
  zip-handler-hello-post

# Run all tests
# test:
# go test ./...

# Format code
# fmt:
# go fmt ./...

# Clean output binaries
# clean:
# rm -rf bin/
