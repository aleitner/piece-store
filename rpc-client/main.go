
// Package main implements a simple gRPC client that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It interacts with the route guide service whose definition can be found in routeguide/route_guide.proto.
package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/aleitner/piece-store/rpc-client/api"
	"github.com/aleitner/piece-store/rpc-client/utils"
	"github.com/urfave/cli"
	"github.com/zeebo/errs"
)

var ArgError = errs.Class("argError")

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
				if c.Args().Get(0) == "" {
					return ArgError.New("No input file specified")
				}

				file, err := os.Open(c.Args().Get(0))
				if err != nil {
					return err
				}
				// Close the file when we are done
				defer file.Close()

				fileInfo, err := os.Stat(c.Args().Get(0))

				if fileInfo.IsDir() {
					return ArgError.New(fmt.Sprintf("Path (%s) is a directory, not a file", c.Args().Get(0)))
				}

				var offset int64 = 0
				var length int64 = fileInfo.Size()
				var ttl int64 = 86400

				hash, err := utils.DetermineHash(file, offset, length)
				if err != nil {
					return err
				}

				err = api.StoreShardRequest(conn, hash, file, length, ttl, offset)
        return err
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
		{
      Name:    "delete",
      Aliases: []string{"x"},
      Usage:   "delete data",
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
