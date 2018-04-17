// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package pstore // import "storj.io/storj/pkg/pstore"

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"testing"
)

func createTestFile() (os.File, error) {
	content := []byte("butts")
	tmpfile, err := ioutil.TempFile("", "api_test")
	if err != nil {
		return nil, err
	}

	if _, err := tmpfile.Write(content); err != nil {
		return nil, err
	}

	return tmpfile, nil
}

func deleteTestFile(tmpfile *os.File) error {
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}

	err := os.Remove(tmpfile.Name())
	if err != nil {
		return err
	}

	return nil
}

func TestStore(t *testing.T) {
	tmpFile, err := createTestFile()
	if err != nil {
		t.Errorf("Failed to generate test data: %s", err.Error())
		return
	}
	defer deleteTestFile(tmpFile)

	reader := bufio.NewReader(tmpFile)

	hash := "0123456789ABCDEFGHIJ"
	Store(hash, reader, os.TempDir())

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
	_, _ = createdFile.Read(buffer)

	if string(buffer) != "butts" {
		t.Errorf("Expected data butts does not equal Actual data %s", string(buffer))
	}
}

func TestRetrieve(t *testing.T) {
	tmpFile, err := createTestFile()
	if err != nil {
		t.Errorf("Failed to generate test data: %s", err.Error())
		return
	}
	defer deleteTestFile(tmpFile)

	reader := bufio.NewReader(tmpFile)

	hash := "0123456789ABCDEFGHIJ"
	Store(hash, reader, os.TempDir())

	// Create file for retrieving data into
	retrievalFilePath := path.Join(os.TempDir(), "retrieved.txt")
	retrievalFile, retrievalFileError := os.OpenFile(retrievalFilePath, os.O_RDWR|os.O_CREATE, 0777)
	if retrievalFileError != nil {
		t.Errorf("Error creating file: %s", retrievalFileError.Error())
		return
	}
	defer retrievalFile.Close()

	writer := bufio.NewWriter(retrievalFile)

	retrieveErr := Retrieve(hash, writer, os.TempDir())

	if retrieveErr != nil {
		t.Errorf("Retrieve Error: %s", retrieveErr.Error())
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
	tmpFile, err := createTestFile()
	if err != nil {
		t.Errorf("Failed to generate test data: %s", err.Error())
		return
	}
	defer deleteTestFile(tmpFile)

	reader := bufio.NewReader(tmpFile)

	hash := "0123456789ABCDEFGHIJ"
	Store(hash, reader, os.TempDir())

	folder1 := string(hash[0:2])
	folder2 := string(hash[2:4])
	fileName := string(hash[4:])

	_, existErr := os.Stat(path.Join(os.TempDir(), folder1, folder2, fileName))
	if existErr != nil {
		t.Errorf("Failed to Store test file")
		return
	}

	Delete(hash, os.TempDir())
	_, deletedExistErr := os.Stat(path.Join(os.TempDir(), folder1, folder2, fileName))
	if deletedExistErr == nil {
		t.Errorf("Failed to Delete test file")
		return
	}
}
