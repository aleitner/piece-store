package piecestore // import "github.com/aleitner/piece-store/pkg"

import (
	"fmt"
  "bufio"
  "io"
)

type hashError struct {
  hash string
  msg string
}

func (e *hashError) Error() string {
  return fmt.Sprintf("HashError (%s): %s", e.hash, e.msg)
}


func Store(hash string, r *bufio.Reader) (int, error) {
  fmt.Println("Storing...")
  if len(hash) < 20 {
    return 1, &hashError{hash, "Hash is too short"}
  }

  folder1 := string(hash[0:2])
  folder2 := string(hash[2:4])
  fileName := string(hash[4:])
  fmt.Println(fmt.Sprintf("%s/%s/%s", folder1, folder2, fileName))

  buffer := make([]byte, 4096)
	for {
		n, err := r.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Println(string(buffer[:n]))
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
