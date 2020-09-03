# go-grpc
====

GRPC server and client with golang based on grpc lib

- - - - 

## Necessary Technology Versions

Technology  | Version
------------- | -------------
Go | go1.14.3 linux/amd64
Docker | 18.09.6
docker-compose | 1.24.1
libprotoc | 3.13.0

## Pre-running

If you're going to start the application manually

    $ cd ssl && ./generate-keys.sh

If you're going to start the application with docker

    $ cd ssl && ./generate-keys-docker.sh

## Regenerate protobuf files

If any changes were made to proto files 

    $ cd internal/store && protoc --go_out=. --go-grpc_out=. store.proto

## Running

To run the chat server we create a docker container for it

    $ docker-compose up -d

## Configurations

### Server Environment Variables

| Name | Description | Default |
| ---- | ----------- | ------- |
| PORT | Server Port | 4000 |

### Client Environment Variables

| Name | Description | Default |
| ---- | ----------- | ------- |
| ADDR | Server IP | localhost |
| PORT | Server Port | 4000 |