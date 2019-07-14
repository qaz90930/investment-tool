## Run
> go run main.go

## Migration
> migrate -source file://migrations -database "postgres://hank:password@localhost:5432/tool_db?sslmode=disable" up 2

## Packages
1. database/sql
2. golang-migrate
