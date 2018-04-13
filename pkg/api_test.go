package piecestore // import "github.com/aleitner/piece-store/pkg"

import "testing"

func TestGetStoreInfo(t *testing.T) {
    GetStoreInfo("/tmp/")
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
