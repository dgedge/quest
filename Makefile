.PHONY: compile
compile: ## Compile the proto file.
	protoc -I pkg/proto/question/ pkg/proto/question/question.proto --go-grpc_out=.
	protoc -I pkg/proto/question/ pkg/proto/question/question.proto --go_out=.

.PHONY: server
server: ## Build and run server.
	go build -o cmd/server cmd/server/main.go
	chmod +x cmd/server/main
	cmd/server/main

.PHONY: client
client: ## Build and run client.
	go build -o cmd/client cmd/client/main.go
	chmod +x cmd/client/main
	cmd/client/main

.PHONY: rest
rest: ## Build and run resful interface
	go build -o cmd/rest cmd/rest/main.go
	chmod +x cmd/rest/main
	cmd/rest/main
