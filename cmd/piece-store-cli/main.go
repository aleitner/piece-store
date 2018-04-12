package main // import "github.com/aleitner/piece-store"

import (
	"fmt"
	"sort"
	"github.com/aleitner/piece-store/pkg"
	"github.com/urfave/cli"
	"log"
	"bufio"
	"os"
)

type argError struct {
    msg string
}

func (e *argError) Error() string {
    return fmt.Sprintf("ArgError: %s", e.msg)
}

func main() {
	app := cli.NewApp()

  app.Name = "Piece Store CLI"
  app.Usage = "Store data in hash folder structure"
	app.Version = "1.0.0"

  app.Flags = []cli.Flag {}

  app.Commands = []cli.Command{
    {
      Name:    "store",
      Aliases: []string{"s"},
      Usage:   "Store data by hash",
      Action:  func(c *cli.Context) error {
				if c.Args().Get(0) == "" {
					return &argError{"Missing data Hash"}
				}

				// NB: Use stdin if no file specified
				// files that don't exist cause infinite loop
				if c.Args().Get(1) == "" {
					return &argError{"No input file specified"}
				}

				file, err := os.Open(c.Args().Get(1))
				if err != nil {
					return err
				}

				// Close the file when we are done
				defer file.Close()

				reader := bufio.NewReader(file)

				_, err = piecestore.Store(c.Args().Get(0), reader)

        return err
      },
    },
    {
      Name:    "retrieve",
      Aliases: []string{"r"},
      Usage:   "Retrieve data by hash",
      Action:  func(c *cli.Context) error {
				if c.Args().Get(0) == "" {
					return &argError{"Missing data Hash"}
				}

				_, err := piecestore.Retrieve(c.Args().Get(0), nil)

        return err
      },
    },
  }

  sort.Sort(cli.FlagsByName(app.Flags))
  sort.Sort(cli.CommandsByName(app.Commands))

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
