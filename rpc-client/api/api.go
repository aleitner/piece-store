package api

import (
  "io"
  "log"
  "os"

  "golang.org/x/net/context"

  "google.golang.org/grpc"

  pb "github.com/aleitner/piece-store/routeguide"
)

func StoreShardRequest(conn *grpc.ClientConn, hash string, data *os.File, length int64, ttl int64, offset int64) (error) {
  c := pb.NewRouteGuideClient(conn)

  stream, err := c.Store(context.Background())

  buffer := make([]byte, 4096)
  for {
    // Read data from read stream into buffer
    n, err := data.Read(buffer)
    if err == io.EOF {
      break
    }

    // Write the buffer to the stream we opened earlier
    if err := stream.Send(&pb.ShardStore{Hash: hash, Size: length, Ttl: ttl, StoreOffset: offset, Content: buffer[:n]}); err != nil {
  		log.Fatalf("%v.Send() = %v", stream, err)
  	}
  }

  reply, err := stream.CloseAndRecv()
  if err != nil {
    log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
  }
  log.Printf("Route summary: %v", reply)

  return nil
}

func DeleteShardRequest(conn *grpc.ClientConn, hash string) (error) {
  c := pb.NewRouteGuideClient(conn)

  reply, err := c.Delete(context.Background(), &pb.ShardDelete{Hash: hash})
  if err != nil {
    return err
  }
  log.Printf("Route summary : %v", reply)
  return nil
}
