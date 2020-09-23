GO_FILES    = env.go eval.go lis.go operators.go token.go tokenize.go
BINARY_NAME = lisgo

.PHONY: build

build: ${GO_FILES} 
	go build -o ${BINARY_NAME} $^
