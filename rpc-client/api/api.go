package api

import (
  "io"
  "log"
  "os"

  "golang.org/x/net/context"

  "google.golang.org/grpc"

  ps "github.com/aleitner/piece-store/routeguide"
)

func StoreShardRequest(conn *grpc.ClientConn, hash string, data *os.File, length int64, ttl int64, offset int64) (error) {
  c := ps.NewRouteGuideClient(conn)

  stream, err := c.Store(context.Background())

  buffer := make([]byte, 4096)
  for {
    // Read data from read stream into buffer
    n, err := data.Read(buffer)
    if err == io.EOF {
      break
    }

    // Write the buffer to the stream we opened earlier
    stream.Send(&ps.ShardStore{Hash: hash, Size: length, Ttl: ttl, StoreOffset: offset, Content: buffer[:n]})
  }

  reply, err := stream.CloseAndRecv()
  if err != nil {
    log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
  }
  log.Printf("Route summary: %v", reply)

  return nil
}
