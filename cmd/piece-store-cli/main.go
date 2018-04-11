package main // import "github.com/aleitner/piece-store"

import (
	"fmt"
	"github.com/aleitner/piece-store/pkg"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello")
	piecestore.Touchbutts()
}
