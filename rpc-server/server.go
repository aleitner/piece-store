// Package main implements a simple gRPC server that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It implements the route guide service whose definition can be found in routeguide/route_guide.proto.
package main

import (
	"flag"
	"fmt"
  "log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	// "github.com/golang/protobuf/proto"

	pb "github.com/aleitner/piece-store/routeguide"
)

var (
	port       = flag.Int("port", 10000, "The server port")
)

type routeGuideServer struct {
}

func (s *routeGuideServer) Store(stream pb.RouteGuide_StoreServer) error {
  fmt.Println("Storing data")

  return nil
}

func (s *routeGuideServer) Retrieve(rect *pb.ShardRetrieval, stream pb.RouteGuide_RetrieveServer) error {
  fmt.Println("Retrieving data")

  return nil
}

func (s *routeGuideServer) Delete(context.Context, *pb.ShardDelete) (*pb.Summary, error) {
  fmt.Println("Deleting data")

  return nil, nil
}

func main() {
  flag.Parse()
  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
  if err != nil {
          log.Fatalf("failed to listen: %v", err)
  }
  grpcServer := grpc.NewServer()
  pb.RegisterRouteGuideServer(grpcServer, &routeGuideServer{})
  grpcServer.Serve(lis)
}
