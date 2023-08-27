BINARY_NAME=outdoorsy_app

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
 	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

run: build
	./${BINARY_NAME}