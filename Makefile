SERVICE_NAME=price-service

deps:
	@go mod download

build: 
	@go build -o bin/${SERVICE_NAME}

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/${SERVICE_NAME}

start:
	@./bin/${SERVICE_NAME} -port ${p}

docker-build:
	@docker build -t ${SERVICE_NAME} .

docker-run:
	@docker run -p ${p}:${p} -d ${SERVICE_NAME} 

clean:
	@rm -rf bin