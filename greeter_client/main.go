package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpcTest/common"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

type Request struct {
	Clazz  string            `json:"clazz"`
	Method string            `json:"method"`
	Args   []string          `json:"args"`
	Kwargs map[string]string `json:"kwargs"`
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCommonServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	kwargs := map[string]string{"name": "wt"}
	data := Request{Clazz: "rpc_client.hello.test", Method: "User.get2", Args: nil, Kwargs: kwargs}
	body, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("json.Marshaler error: %v", err)
	}
	requestId := strings.Replace(uuid.NewString(), "-", "", -1)
	r, err := c.Handle(ctx, &pb.CommonRequest{RequestId: requestId, Request: string(body), Serialize: 3})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Response)
}
