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

## result

### list

```
[~/study/go/gacha-app/server/app] % grpcurl --plaintext localhost:8080 list
gacha.Gacha
grpc.reflection.v1alpha.ServerReflection
user.User
```

### User.Create

```
[~/work/myapp/study/go/gacha-app] % grpcurl --plaintext -emit-defaults -d '{"name": "test"}' localhost:8080 user.User.Create
{
  "id": "63beba5b6792a1c71564bc28",
  "name": "test",
  "createdAt": "2023-01-11T13:32:11.972707Z"
}
```

### User.ListUserItems

```
[~/work/myapp/study/go/gacha-app] % grpcurl --plaintext -emit-defaults -d '{"user_id": "63bcd84dfe44aca595b37536"}' localhost:8080 user.User.ListUserItems
{
  "items": [
    {
      "itemId": "item3",
      "name": "item3",
      "count": 3
    },
    {
      "itemId": "item2",
      "name": "item2",
      "count": 1
    }
  ]
}
```

### Gacha.Buy

```
[~/work/myapp/study/go/gacha-app] % grpcurl --plaintext -emit-defaults -d '{"user_id": "63bcd84dfe44aca595b37536", "gacha_id": "gacha1"}' localhost:8080 gacha.Gacha.Buy                                           
{
  "item": {
    "itemId": "item3",
    "name": "",
    "count": 4
  }
}
```