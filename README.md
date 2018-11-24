*HealthChecker on Golang*

run app `go run main.go`

run migrations

```
cd migrations
go run *.go
```

reset `go run *.go reset`

up `go run *.go up 2`

down `go run *.go down 2`

check migration version `go run *.go version`

*Turn on Module mod*

```
go mod init
go mod tidy
go get ./...
go mod vendor // if needed
```
