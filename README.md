# Nats communication with Go
This is an example script that is able to send and receive data via the nats protocol using Go.


## Prerequisites

- Golang
- Docker

## Installation

Nats server (docker container)
1. Pull the latest docker image from nats: ```docker pull nats```
2. Build a container via the following command: ```sudo docker run --network host -p 4222:4222 nats -js```

Go dependencies
1. in the root of the project run ```go mod tidy```.

## Running the script

you can run the script by executing ```go run ./nats_jetsteam.go```