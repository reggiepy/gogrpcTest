module grpcTest

go 1.16

require (
	github.com/google/uuid v1.3.0
	golang.org/x/tools v0.1.12
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.51.0
