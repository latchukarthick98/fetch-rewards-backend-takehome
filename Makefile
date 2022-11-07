BINARY_NAME=fetch-rewards-api

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go
	go build -o ${BINARY_NAME} main.go
run:
	./${BINARY_NAME}

build_and_run: build run

build_docker:
	docker-compose build

run_docker:
	docker-compose up -d

stop_docker:
	docker-compose down

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

test:
	go test -v

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all