# Mock microservice

## Description

This is a mock pricing microservice that can be used to test the communication with any other microservice.

## Tech stack

- Golang 1.19
- GRPC
- Protobuf
- JSON

## How to run

### Run locally

1. Clone the repository
2. Run `make deps` to install dependencies
3. Run `make start p=3000` to run the binary

### Run with Docker

1. Clone the repository
2. Run `make docker-build` to build the docker image
3. Run `make docker-run p=3000` to run the docker image
