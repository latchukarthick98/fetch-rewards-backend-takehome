# fetch-rewards-backend-takehome
This Repository consists of the codebase for the take home assesment by Fetch Rewards.

# Assumptions Made
1. When an intial transaction (when there is no history of transactions) is negative , drop the transaction as it is of no use.
2. Deny the request if the points requested to spend is more than the available points. 

# Getting Started

## Environment Variables
Make sure the `.env` file is present. It is important for the API to start, as it has the PORT number to which it has to listen.

If `.env` file is not present, copy it from the template `.env.example` by running:
```
cp .env.example .env
```

### Example
`.env`
```
PORT=3001
```
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

# API Documentation

## Routes

## Get Balance
Path: `/balance`
Method: `GET`
Content-Type: `application/json`
Response-Type: `application/json`
### Example
Request: `http://localhost:${PORT}/balance`
curl: `curl -X GET http://localhost:${PORT}/balance`
Response:
```
{
    "DANNON": 1000,
    "UNILEVER": 0
}
```

## Post Transaction
Path: `/transaction`
Method: `POST`
Content-Type: application/json
Response-Type: `application/json`

### Example
Request: `http://localhost:${PORT}/transaction`

curl: `curl -d '{ "payer": "DANNON", "points": 300, "timestamp": "2022-10-31T10:00:00Z" }' -H "Content-Type: application/json" -X POST http://localhost:${PORT}/transaction`

Response:
```
{
	"payer": "DANNON",
	"timestamp": "2022-11-02T14:00:00Z",
	"points": -200
}
```


## Spend points
Path: `/spend`
Method: `POST`
Content-Type: application/json
Response-Type: `application/json`

### Example
Request: `http://localhost:${PORT}/spend`

curl: `curl -d '{"points": 5000}' -H "Content-Type: application/json" -X POST http://localhost:${PORT}/spend`

Response:
```
[
	{
		"payer": "DANNON",
		"points": -100
	},
	{
		"payer": "UNILEVER",
		"points": -200
	},
	{
		"payer": "MILLER COORS",
		"points": -4700
	}
]
```