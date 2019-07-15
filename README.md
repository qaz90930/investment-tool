## Run
* start app 
> go run main.go // localhost:8001

* run in development  
> realize start --path="." // dev-tool: localhost:5001 

## Migration
> migrate -source file://migrations -database "postgres://hank:password@localhost:5432/tool_db?sslmode=disable" up 2

## Packages
1. [database/sql](https://golang.org/pkg/database/sql/)
2. [golang-migrate](https://github.com/golang-migrate/migrate)
3. [golang-realize](https://github.com/oxequa/realize)
