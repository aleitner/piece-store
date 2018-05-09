// Package main implements a simple gRPC server that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It implements the route guide service whose definition can be found in routeguide/route_guide.proto.
package api

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"golang.org/x/net/context"

	pb "github.com/aleitner/piece-store/routeguide"
	"github.com/aleitner/piece-store/src"
)

type Server struct {
  PieceStoreDir string
}

func (s *Server) Store(stream pb.RouteGuide_StoreServer) error {
  fmt.Println("Storing data...")
	startTime := time.Now()
	var total int64 = 0
	for {
		shardData, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Successfully stored data...")
			endTime := time.Now()
			return stream.SendAndClose(&pb.ShardStoreSummary{
				Status:   0,
				Message: "OK",
				TotalReceived: total,
				ElapsedTime:  int64(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}

		length := int64(len(shardData.Content))

		// Write chunk received to disk
		err = pstore.Store(shardData.Hash, bytes.NewReader(shardData.Content), length, total + shardData.StoreOffset, s.PieceStoreDir)

		if err != nil {
			fmt.Println("Store data Error: ", err.Error())
			endTime := time.Now()
			return stream.SendAndClose(&pb.ShardStoreSummary{
				Status:   -1,
				Message: err.Error(),
				TotalReceived: total,
				ElapsedTime:  int64(endTime.Sub(startTime).Seconds()),
			})
		}

		total += length
	}
  return nil
}

func (s *Server) Retrieve(shardMeta *pb.ShardRetrieval, stream pb.RouteGuide_RetrieveServer) error {
  fmt.Println("Retrieving data...")

	buffer := make([]byte, 4096)
	var total int64 = 0
	for {
		n, retrieveErr := pstore.Retrieve(shardMeta.Hash, bytes.NewBuffer(buffer), int64(cap(buffer)), shardMeta.StoreOffset + total, s.PieceStoreDir)
		if retrieveErr != nil {
			if retrieveErr != io.EOF {
				return retrieveErr
			}
		}

		fmt.Println("Sending data...")

		// Write the buffer to the stream we opened earlier
		if err := stream.Send(&pb.ShardRetrievalStream{Content: buffer[:len(buffer)]}); err != nil {
			fmt.Println("%v.Send() = %v", stream, err)
			return err
		}

		if retrieveErr == io.EOF {
			break
		}

		total += n
	}

  return nil
}

func (s *Server) Delete(context.Context, *pb.ShardDelete) (*pb.ShardDeleteSummary, error) {
  fmt.Println("Deleting data")

  return nil, nil
}
