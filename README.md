
# GoMemoryServer
* Host gRPC Server with Go to remember at max 3 numbers.

### commands
```

go mod init github.com/<github-username>/<module-name> -- this will initiate the mod file.

go get go.uber.org/zap -- this will add the dependancy in go.mod file

go mod tidy --> add/remove dependancy from the mod file

go mod why  

-- execute this from project directory
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/ms.proto
```

### Server execution
➜  server go run main.go

### Client Execution
➜  client go run main.go 1


