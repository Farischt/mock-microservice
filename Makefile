SERVICE_NAME=price-service

deps:
	@go mod download

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/service.proto

.PHONY: proto

build: 
	@go build -o bin/${SERVICE_NAME}

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/${SERVICE_NAME}

run:
	@./bin/${SERVICE_NAME} -port ${p}

start: 
	make build & make run p=${p}

docker-build:
	@docker build -t ${SERVICE_NAME} .

docker-run:
	@docker run -p ${p}:${p} -d ${SERVICE_NAME} 

clean:
	@rm -rf bin
