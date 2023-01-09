# demo gacha app

go + gRPC + mongo + redis + [franela/goblin](https://github.com/franela/goblin)

## proto
https://github.com/JY8752/gacha-app-proto

## setup

### init
```
go mod init JY8752/gacha-app
```

### gRPC
```
go get -u google.golang.org/grpc
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### mongo
```
go get -u go.mongodb.org/mongo-driver/mongo
```

### config
```
go get -u github.com/BurntSushi/toml@latest
go get -u github.com/joho/godotenv
```

## gRPC

### コード生成

```
protoc --go_out=./pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=./pkg/grpc --go-grpc_opt=paths=source_relative\
  -I=./proto \
	proto/**/*.proto
```

## test

### dockertest

```
go get -u github.com/ory/dockertest/v3
```

### testcontainers

```
go get -u github.com/testcontainers/testcontainers-go
```

### goblin
mochaライクなBDDテスティングフレームワーク

```
go get -u github.com/franela/goblin
```