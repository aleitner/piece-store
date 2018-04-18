// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package pstore // import "storj.io/storj/pkg/pstore"

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)

var tmpfile string

func TestStore(t *testing.T) {
	file, err := os.Open(tmpfile)
	if err != nil {
		t.Errorf("Error opening tmp file: %s", err.Error())
		return
	}

	fi, err := file.Stat()
	if err != nil {
		t.Errorf("Could not stat test file: %s", err.Error())
		return
	}

	defer file.Close()

	hash := "0123456789ABCDEFGHIJ"
	Store(hash, file, int64(fi.Size()), 0, os.TempDir())

	folder1 := string(hash[0:2])
	folder2 := string(hash[2:4])
	fileName := string(hash[4:])

	createdFilePath := path.Join(os.TempDir(), folder1, folder2, fileName)
	defer os.RemoveAll(path.Join(os.TempDir(), folder1))
	_, lStatErr := os.Lstat(createdFilePath)
	if lStatErr != nil {
		t.Errorf("No file was created from Store(): %s", lStatErr.Error())
		return
	}

	createdFile, openCreatedError := os.Open(createdFilePath)
	if openCreatedError != nil {
		t.Errorf("Error: %s opening created file %s", openCreatedError.Error(), createdFilePath)
	}
	defer createdFile.Close()

	buffer := make([]byte, 5)
	createdFile.Seek(0, 0)
	_, _ = createdFile.Read(buffer)

	if string(buffer) != "butts" {
		t.Errorf("Expected data butts does not equal Actual data %s", string(buffer))
	}
}

func TestRetrieve(t *testing.T) {
	file, err := os.Open(tmpfile)
	if err != nil {
		t.Errorf("Error opening tmp file: %s", err.Error())
		return
	}

	fi, err := file.Stat()
	if err != nil {
		t.Errorf("Could not stat test file: %s", err.Error())
		return
	}

	defer file.Close()

	hash := "0123456789ABCDEFGHIJ"
	Store(hash, file, int64(fi.Size()), 0, os.TempDir())

	// Create file for retrieving data into
	retrievalFilePath := path.Join(os.TempDir(), "retrieved.txt")
	retrievalFile, err := os.OpenFile(retrievalFilePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		t.Errorf("Error creating file: %s", err.Error())
		return
	}
	defer retrievalFile.Close()

	err = Retrieve(hash, retrievalFile, int64(fi.Size()), 0, os.TempDir())

	if err != nil {
		t.Errorf("Retrieve Error: %s", err.Error())
	}

	buffer := make([]byte, 5)

	retrievalFile.Seek(0, 0)
	_, _ = retrievalFile.Read(buffer)

	fmt.Printf("Retrieved data: %s", string(buffer))

	if string(buffer) != "butts" {
		t.Errorf("Expected data butts does not equal Actual data %s", string(buffer))
	}
}

func TestDelete(t *testing.T) {
	file, err := os.Open(tmpfile)
	if err != nil {
		t.Errorf("Error opening tmp file: %s", err.Error())
		return
	}

	fi, err := file.Stat()
	if err != nil {
		t.Errorf("Could not stat test file: %s", err.Error())
		return
	}

	defer file.Close()

	hash := "0123456789ABCDEFGHIJ"
	Store(hash, file, int64(fi.Size()), 0, os.TempDir())

	folder1 := string(hash[0:2])
	folder2 := string(hash[2:4])
	fileName := string(hash[4:])

	if _, err := os.Stat(path.Join(os.TempDir(), folder1, folder2, fileName)); err != nil {
		t.Errorf("Failed to Store test file")
		return
	}

	Delete(hash, os.TempDir())
	_, err = os.Stat(path.Join(os.TempDir(), folder1, folder2, fileName))
	if err == nil {
		t.Errorf("Failed to Delete test file")
		return
	}
}

func TestMain(m *testing.M) {
	content := []byte("butts")
	tmpfilePtr, err := ioutil.TempFile("", "api_test")
	if err != nil {
		log.Fatal(err)
	}

	// defer os.Remove(tmpfile.Name()) // clean up
	tmpfile = tmpfilePtr.Name()

	if _, err := tmpfilePtr.Write(content); err != nil {
		log.Fatal(err)
	}

	if err := tmpfilePtr.Close(); err != nil {
		log.Fatal(err)
	}

	m.Run()

	os.Exit(0)
}
