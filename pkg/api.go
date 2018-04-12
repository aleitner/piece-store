package piecestore // import "github.com/aleitner/piece-store/pkg"

import (
	"fmt"
  "bufio"
  "io"
)

func Store(hash string, r *bufio.Reader) (int, error) {
  fmt.Println("Storing...")

  p := make([]byte, 4)
	for {
		n, err := r.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
  return 0, nil
}

func Retrieve(hash string, w interface{}) (int, error) {
  fmt.Println("Retrieving...")

  return 0, nil
}

func Delete(hash string) (int, error) {
  fmt.Println("Deleting...")

  return 0, nil
}

func GetStoreInfo() {
  fmt.Println("Getting store info")
}
