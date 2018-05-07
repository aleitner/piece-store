
// Package main implements a simple gRPC client that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It interacts with the route guide service whose definition can be found in routeguide/route_guide.proto.
package main

import (
	// "flag"
	// "io"
	"log"
	"os"
	// "math/rand"
	// "time"

	"google.golang.org/grpc"

	"github.com/aleitner/piece-store/rpc-client/api"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()

	// Set up connection with rpc server
  var conn *grpc.ClientConn
  conn, err := grpc.Dial(":7777", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()

	app.Commands = []cli.Command{
    {
      Name:    "upload",
      Aliases: []string{"u"},
      Usage:   "upload data",
      Action:  func(c *cli.Context) error {
				api.StoreRequest(conn)
        return nil
      },
    },
    {
      Name:    "download",
      Aliases: []string{"d"},
      Usage:   "download data",
      Action:  func(c *cli.Context) error {
        return nil
      },
    },
  }

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
