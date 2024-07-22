package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// FileSystemGtfsStore stores gtfs files details in the filesystem.
type FileSystemGtfsStore struct {
	database *json.Encoder
	files    map[string]GtfsFile
}

// FileSystemGtfsStoreFromFile creates a GtfsProxyStore from the contents of a JSON file found at path.
func FileSystemGtfsStoreFromFile(path string) (*FileSystemGtfsStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemGtfsStore(db)

	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system gtfs store, %v ", err)
	}

	return store, closeFunc, nil
}

// NewFileSystemGtfsStore creates a FileSystemGtfsStore initialising the store if needed.
func NewFileSystemGtfsStore(file *os.File) (*FileSystemGtfsStore, error) {
	err := initialiseGtfsDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initialising gtfs db file, %v", err)
	}

	return &FileSystemGtfsStore{
		database: json.NewEncoder(&Tape{file}),
	}, nil
}

func initialiseGtfsDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}

	return nil
}

// UpdateFile updates the time of the file
func (f *FileSystemGtfsStore) UpdateFile(name string, time int64) {
	//TODO implement me
	panic("implement me")
}

// GetFileDetails returns the details of the file
func (f *FileSystemGtfsStore) GetFileDetails(name string) GtfsFile {
	//details := f.files.Find(name)
	//// do some
	//f.files.Encode(f.files)

	//TODO implement me
	panic("implement me")
}

// SaveFile saves the file
func (f *FileSystemGtfsStore) SaveFile(name string, file io.Writer) error {
	//TODO implement me
	panic("implement me")
}

// GetFile returns the file
func (f *FileSystemGtfsStore) GetFile(name string) (io.Reader, error) {
	//TODO implement me
	panic("implement me")
}
