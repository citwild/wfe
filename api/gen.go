package api

//go:generate go get github.com/golang/protobuf/protoc-gen-go
//go:generate go get google.golang.org/grpc
//go:generate sh -c "protoc --go_out=plugins=grpc:. *.proto"
