package api

import (
  "log"

  "golang.org/x/net/context"

  "google.golang.org/grpc"

  ps "github.com/aleitner/piece-store/routeguide"
)

func StoreRequest(conn *grpc.ClientConn) {
  c := ps.NewRouteGuideClient(conn)

  storemsg := &ps.ShardStore{Hash: "ABCD", Size: 1024, Ttl: 86400, StoreOffset: 0, Content: "Hello World!"}
   stream, err := c.Store(context.Background())
  stream.Send(storemsg)
  reply, err := stream.CloseAndRecv()
  if err != nil {
    log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
  }
  log.Printf("Route summary: %v", reply)
}
