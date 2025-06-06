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
	make build-handler-hello-get
	make build-handler-hello-post

# Zip all handlers
zip:
	make zip-handler-hello-get
	make zip-handler-hello-post