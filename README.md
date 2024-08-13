# webservice-prototype

A prototype building out with GraphQL, Golang, and MongoDB

## Setup

### Go work

```txt
go 1.22.5

use ./web
```

## Running

From the root of the application...

1. Start Docker mongo database instance, and mongo express instance

```sh
docker compose up -d
```

2. Start the main application

```sh
go run main.go
```

3. test the application is up and healthy

```sh
curl http://localhost:8080/
```
