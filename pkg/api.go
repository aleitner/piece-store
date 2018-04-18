// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

// store, retrieve, and delete data from Storj farmers

package pstore // import "storj.io/storj/pkg/pstore"

import (
	"github.com/zeebo/errs"
	"io"
	"os"
	"path"
	"path/filepath"
)

// Errors
var (
	ArgError = errs.Class("argError")
	FSError  = errs.Class("fsError")
)

// creates datapath from hash and dir
func pathByHash(hash, dir string) (string, error) {
	if len(hash) < 20 {
		return "", ArgError.New("Invalid hash length")
	}

	folder1 := string(hash[0:2])
	folder2 := string(hash[2:4])
	fileName := string(hash[4:])

	return path.Join(dir, folder1, folder2, fileName), nil
}

/*
	Store

	Store data into piece store

	hash 		(string)		Hash of the data to be stored
	r 			(io.Reader)	File/Stream that contains the contents of the data to be stored
	length 	(length)		Size of the data to be stored
	offset 	(offset)		Offset of the data that you are writing. Useful for multiple connections to split the data transfer
	dir 		(string)		pstore directory containing all other data stored
	returns (error) if failed and nil if successful
*/
func Store(hash string, r io.Reader, length int64, offset int64, dir string) error {
	if len(hash) < 20 {
		return ArgError.New("Hash is too short. Must be atleast 20 bytes")
	}

	if dir == "" {
		return ArgError.New("No path provided")
	}

	dataPath, err := pathByHash(hash, dir)
	if err != nil {
		return err
	}

	// Create directory path on file system
	if err = os.MkdirAll(filepath.Dir(dataPath), 0700); err != nil {
		return err
	}

	// Create File on file system
	dataFile, err := os.OpenFile(dataPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	// Close when finished
	defer dataFile.Close()

	dataFile.Seek(offset, 0)

	if _, err := io.CopyN(dataFile, r, length); err != nil {
		if err != io.EOF {
			return err
		}
	}

	return nil
}

/*
	Retrieve

	Retrieve data from pstore directory

	hash 		(string)		Hash of the stored data
	w 			(io.Writer)	File/Stream that recieves the stored data
	length 	(length)		Amount of data to read
	offset 	(offset)		Offset of the data that you are reading. Useful for multiple connections to split the data transfer
	dir 		(string)		pstore directory containing all other data stored
	returns (error) if failed and nil if successful
*/
func Retrieve(hash string, w io.Writer, length int64, offset int64, dir string) error {
	if len(hash) < 20 {
		return ArgError.New("Hash is too short. Must be atleast 20 bytes")
	}

	if dir == "" {
		return ArgError.New("No path provided")
	}

	dataPath, err := pathByHash(hash, dir)
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(dataPath)

	// If offset is greater than file size return
	if offset >= fileInfo.Size() || offset < 0 {
		return ArgError.New("Invalid offset: %v", offset)
	}

	// If length is 0 read the entire file
	if length <= 0 {
		length = fileInfo.Size()
	}

	// If trying to read past the end of the file, just read to the end
	if fileInfo.Size() < offset+length {
		length = fileInfo.Size() - offset
	}

	dataFile, err := os.OpenFile(dataPath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	// Close when finished
	defer dataFile.Close()

	dataFile.Seek(offset, 0)

	if _, err := io.CopyN(w, dataFile, length); err != nil {
		if err != io.EOF {
			return err
		}
	}

	return nil
}

/*
	Delete

	Delete data from farmer

	hash (string) Hash of the data to be stored
	dir (string) pstore directory containing all other data stored
	returns (error) if failed and nil if successful
*/
func Delete(hash string, dir string) error {
	if len(hash) < 20 {
		return ArgError.New("Hash is too short. Must be atleast 20 bytes")
	}
	if dir == "" {
		return ArgError.New("No path provided")
	}

	dataPath, err := pathByHash(hash, dir)
	if err != nil {
		return err
	}

	if err = os.Remove(dataPath); err != nil {
		return err
	}

	return nil
}
