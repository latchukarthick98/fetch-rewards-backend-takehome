# fetch-rewards-backend-takehome
This Repository consists of the codebase for the take home assesment by Fetch Rewards.

# Assumptions Made
1. When an intial transaction (when there is no history of transactions) is negative , drop the transaction as it is of no use.
2. Deny the request if the points requested to spend is more than the available points. 

# Getting Started

## Installation

## Using Docker and Docker Compose (Recommended)
1. Install Docker and Docker Compose.
2. Build Docker image 
From the Project root run:

```
make docker_build
```
or
```
docker-compose build
```
3. Run Docker container
```
make run_docker
```
or
```
docker-compose up -d
```

4. Stopping the container
```
make stop_docker
```
or
```
docker-compose down
```

## Direct Install (Alternative)
## Install Golang

Make sure you have Go 1.13 or higher installed.

https://golang.org/doc/install

## Installing Dependences
From the project root directory run:

```
make dep
```

or

```
go mod download
```

## Building and Running the Project

```
make build_and_run
```
or
```
go build -o fetch-rewards-api main.go
./fetch-rewards-api

```

## Testing
From project root directory run:

```
make test
```

or

```
go test -v
```