module github.com/shuymn-sandbox/go-grpc-sample/go/blog/client

go 1.19

replace github.com/shuymn-sandbox/go-grpc-sample/go/protobuf/protoc/blog => ../../protobuf/protoc/blog

replace github.com/shuymn-sandbox/go-grpc-sample/go/protobuf/buf/blog => ../../protobuf/buf/blog

require (
	github.com/shuymn-sandbox/go-grpc-sample/go/protobuf/protoc/blog v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.49.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
