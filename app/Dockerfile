FROM --platform=linux/amd64 golang:1.18

RUN apt-get update && apt-get install -y unzip curl 

# protocolbufのインストール
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip && \
  unzip protoc-3.19.4-linux-x86_64.zip -d protoc3 && \
  mv protoc3/bin/* /usr/local/bin/ && \
  mv protoc3/include/* /usr/local/include/

# protoc-gen-goのインストール
RUN go install github.com/golang/protobuf/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /app

COPY . .
RUN go mod download

# protoのコード生成
RUN protoc --go_out=./pkg/grpc --go_opt=paths=source_relative \
  --go-grpc_out=./pkg/grpc --go-grpc_opt=paths=source_relative \
  -I=./proto \
  proto/**/*.proto

RUN go build -o /gacha-demo-app

EXPOSE 8080

CMD [ "/gacha-demo-app" ]