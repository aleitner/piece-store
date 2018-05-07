// Package main implements a simple gRPC server that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It implements the route guide service whose definition can be found in routeguide/route_guide.proto.
package api

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/aleitner/piece-store/routeguide"
	"github.com/aleitner/piece-store/src"
)

type Server struct {
  // Put resources here that you want available inside RPC calls
}

func (s *Server) Store(stream pb.RouteGuide_StoreServer) error {
  fmt.Println("Storing data")

  return nil
}

func (s *Server) Retrieve(rect *pb.ShardRetrieval, stream pb.RouteGuide_RetrieveServer) error {
  fmt.Println("Retrieving data")

  return nil
}

func (s *Server) Delete(context.Context, *pb.ShardDelete) (*pb.Summary, error) {
  fmt.Println("Deleting data")

  return nil, nil
}
