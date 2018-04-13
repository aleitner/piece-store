package piecestore // import "github.com/aleitner/piece-store/pkg"

import (
    "testing"
    "os"
    "path"
    "bufio"
)


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
  createTestFile()
  defer deleteTestFile()
  file, openTestFileErr := os.Open(testFile)
  if openTestFileErr != nil {
    t.Errorf("Could not open test file")
    return
  }

  // Close the file when we are done
  defer file.Close()

  reader := bufio.NewReader(file)

  hash := "0123456789ABCDEFGHIJ"
  Store(hash, reader, os.TempDir())
  // Run Store()
  folder1 := string(hash[0:2])
  folder2 := string(hash[2:4])
  fileName := string(hash[4:])

	// Create directory path string
  createdFilePath := path.Join(os.TempDir(), folder1, folder2, fileName)
  _, lStatErr := os.Lstat(createdFilePath)
  if lStatErr != nil {
    t.Errorf("No file was created from Store(): %s", lStatErr.Error())
    return
  }

  createdFile, openCreatedError := os.Open(createdFilePath)
  if openCreatedError != nil {
    t.Errorf("Error: %s opening created file %s", openCreatedError.Error() ,createdFilePath)
  }
  defer createdFile.Close()

  buffer := make([]byte, 5)
  _, _ = createdFile.Read(buffer)


  if string(buffer) != "butts" {
    t.Errorf("Expected data butts does not equal Actual data %s", string(buffer))
  }
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
