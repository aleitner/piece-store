package piecestore // import "github.com/aleitner/piece-store/pkg"

import "testing"
import "os"
import "path"

var testFile = path.Join(os.TempDir(), "test.txt")

func createTestFile() (error) {
  file, createFileErr := os.Create(testFile)

  if createFileErr != nil {
    return createFileErr
  }

  b := []byte{'b', 'u', 't', 't', 's'}
  _, writeToFileErr := file.Write(b)

  if writeToFileErr != nil {
    return writeToFileErr
  }

  return nil
}

func deleteTestFile() (error) {
  err := os.Remove(testFile)
  if err != nil {
    return err
  }

  return nil
}

func TestStore(t *testing.T) {
  // Run Store()
  // Check if folder structure and file were created
  // delete the folders and file
}

func TestRetrieve(t *testing.T) {
  // Run Store()
  // Run Retrieve()
  // Verify that the contents of the retrieve match what was stored
  // delete the folders and file
}

func TestDelete(t *testing.T) {
  // Run Store()
  // Verify files exist
  // Run Delete()
  // Verify that the files are no longer existing
}

func TestGetStoreInfo(t *testing.T) {
    GetStoreInfo("/tmp/")
}
