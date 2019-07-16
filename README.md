# Get Started

## Installation

[Install golang](https://golang.org/dl/)

[Install Postgres SQL with **brew**](https://gist.github.com/ibraheem4/ce5ccd3e4d7a65589ce84f2a3b7c23a3)
```$ brew install postgres```

Database Configration

```UNIX
// Create a user for the database
$ sudo -u postgres createuser user_name
// Create a DB for the app
$ sudo -u postgres createdb tool_db -o user_name
// Go to the DB
$ sudo -u postgres psql tool_db
```

Install golang-migrate

```$ brew install golang-migrate```

Create schema & tables for the app

```$ migrate -source file://migrations -database "postgres://hank:password@localhost:5432/tool_db?sslmode=disable" up 2```

Install golang packages

```$ go get package_name```

## Run

### Start app

```$ go run main.go // localhost:8001```

### Run in development  

```$ realize start --path="." // dev-tool: localhost:5001```

## Migration

```$ migrate -source file://migrations -database "postgres://hank:password@localhost:5432/tool_db?sslmode=disable" up 2```

## Packages

1. [database/sql](https://golang.org/pkg/database/sql/)
2. [golang-migrate](https://github.com/golang-migrate/migrate)
3. [golang-realize](https://github.com/oxequa/realize)
4. [scheduler](https://godoc.org/github.com/carlescere/scheduler)
