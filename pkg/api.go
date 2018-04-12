package piecestore // import "github.com/aleitner/piece-store/pkg"

import (
	"fmt"
  "bufio"
  "io"
	"os"
	"path"
)

type hashError struct {
  hash string
  msg string
}

func (e *hashError) Error() string {
  return fmt.Sprintf("HashError (%s): %s", e.hash, e.msg)
}

type fsError struct {
  path string
  msg string
}

func (e *fsError) Error() string {
  return fmt.Sprintf("FSError (%s): %s", e.path, e.msg)
}

func Store(hash string, r *bufio.Reader, dir string) (error) {
  fmt.Println("Storing...")
  if len(hash) < 20 {
    return &hashError{hash, "Hash is too short"}
  }
	if dir == "" {
    return &hashError{dir, "No path provided"}
  }

	// Folder structure
  folder1 := string(hash[0:2])
  folder2 := string(hash[2:4])
  fileName := string(hash[4:])

	// Create directory path string
	dirpath := path.Join(dir, folder1, folder2)

	// Create directory path on file system
	mkDirerr := os.MkdirAll(dirpath, 0700)
	if mkDirerr != nil {
		return mkDirerr
	}

	// Create File Path string
	filepath := path.Join(dirpath, fileName)

	// Create File on file system
	file, openErr := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
	if openErr != nil {
		return openErr
	}

	// Close when finished
	defer file.Close()

	// Buffer for reading data
  buffer := make([]byte, 4096)
	for {
		// Read data from read stream into buffer
		n, err := r.Read(buffer)
		if err == io.EOF {
			break
		}

		// Write to buffer to the file we opened earlier
		_, err = file.Write(buffer[:n])
	}

  return nil
}

func Retrieve(hash string, w interface{}, dir string) (error) {
  fmt.Println("Retrieving...")

  return nil
}

func Delete(hash string, dir string) (error) {
  fmt.Println("Deleting...")

	folder1 := string(hash[0:2])
  folder2 := string(hash[2:4])
  fileName := string(hash[4:])

	err := os.Remove(path.Join(dir, folder1, folder2, fileName))
	if err != nil {
		return err
	}

  return nil
}

func GetStoreInfo(dir string) {
  fmt.Println("Getting store info")
}
