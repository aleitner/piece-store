// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

// store, retrieve, and delete data from Storj farmers

package pstore // import "storj.io/storj/pkg/pstore"

import (
	"bufio"
	"fmt"
	"github.com/zeebo/errs"
	"io"
	"os"
	"path"
)

// Errors
var (
	ArgError = errs.Class("argError")
	FSError  = errs.Class("fsError")
)

/*
	Store

	Store data into piece store

	hash 		(string)		Hash of the data to be stored
	r 			(io.Reader)	File/Stream that containts the contents of the data to be stored
	length 	(length)		Size of the data to be stored
	offset 	(offset)		Offset of the data if you are writing. Useful for multiple connections to split the data transfer
	dir 		(string)		pstore directory containing all other data stored
*/
func Store(hash string, r io.Reader, length int64, offset int64, dir string) error {
	if len(hash) < 20 {
		return ArgError.New("Hash is too short. Must be atleast 20 bytes")
	}
	if dir == "" {
		return ArgError.New("No path provided")
	}

	// Folder structure
	folder1 := string(hash[0:2])
	folder2 := string(hash[2:4])
	fileName := string(hash[4:])

	// Create directory path string
	dirpath := path.Join(dir, folder1, folder2)

	// Create directory path on file system
	if err := os.MkdirAll(dirpath, 0700); err != nil {
		return err
	}

	// Create File Path string
	filepath := path.Join(dirpath, fileName)

	// Create File on file system
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	// Close when finished
	defer file.Close()

	file.Seek(offset, 0)

	if _, err := io.CopyN(file, r, length); err != nil {
		return err
	}

	return nil
}

func Retrieve(hash string, w *bufio.Writer, dir string) error {
	if len(hash) < 20 {
		return ArgError.New("Hash is too short. Must be atleast 20 bytes")
	}
	if dir == "" {
		return ArgError.New("No path provided")
	}

	folder1 := string(hash[0:2])
	folder2 := string(hash[2:4])
	fileName := string(hash[4:])

	filePath := path.Join(dir, folder1, folder2, fileName)

	file, openErr := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if openErr != nil {
		return openErr
	}
	// Close when finished
	defer file.Close()

	file.Seek(0, 0)
	buffer := make([]byte, 4096)
	for {
		// Read data from read stream into buffer
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}

		// Write to buffer to the file we opened earlier
		_, writeError := w.Write(buffer[:n])
		if writeError != nil {
			fmt.Println("Write Error:", writeError)
		}

	}

	w.Flush()

	return nil
}

func Delete(hash string, dir string) error {
	if len(hash) < 20 {
		return ArgError.New("Hash is too short. Must be atleast 20 bytes")
	}
	if dir == "" {
		return ArgError.New("No path provided")
	}

	folder1 := string(hash[0:2])
	folder2 := string(hash[2:4])
	fileName := string(hash[4:])

	err := os.Remove(path.Join(dir, folder1, folder2, fileName))
	if err != nil {
		return err
	}

	return nil
}
