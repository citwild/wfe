package api

//go:generate go get github.com/gogo/protobuf/protoc-gen-gogo
//go:generate go get google.golang.org/grpc
//go:generate protoc --gogo_out=plugins=grpc:. wfe.proto

