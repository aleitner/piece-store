// Package main implements a simple gRPC server that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It implements the route guide service whose definition can be found in routeguide/route_guide.proto.
package main

import (
	"fmt"
  "log"
	"net"

	"google.golang.org/grpc"

  "github.com/aleitner/piece-store/rpc-server/api"
  pb "github.com/aleitner/piece-store/routeguide"

)

func main() {
  // create a listener on TCP port 7777
  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  // create a server instance
  s := api.Server{}

  // create a gRPC server object
  grpcServer := grpc.NewServer()

  // attach the api service to the server
  pb.RegisterRouteGuideServer(grpcServer, &s)

  // start the server
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}
