FROM bufbuild/buf:1.28.1 as BUILDER
FROM golang:1.21-alpine

RUN apk add --no-cache \
  nodejs \
  npm \
  git \
  make

RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest

RUN git clone https://github.com/regen-network/protobuf.git; \
  cd protobuf; \
  go mod download; \
  make install; \
  cd ..

RUN git clone https://github.com/cosmos/gogoproto.git; \
  cd gogoproto/protoc-gen-gocosmos; \
  go install .; \
  cd ..


COPY --from=BUILDER /usr/local/bin /usr/local/bin
