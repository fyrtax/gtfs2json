package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// FileSystemGtfsStore stores gtfs files details in the filesystem.
type FileSystemGtfsStore struct {
	database *json.Encoder
	files    map[string]GtfsFile
}

var gtfsFiles = GtfsServerFiles{
	"TripUpdates_A":      {Type: "protobuf", Filename: "TripUpdates_T.pb", LastChecked: time.Unix(0, 0), ModTime: time.Unix(0, 0), Vehicle: "A", FormattedTime: time.Unix(0, 0).Format(time.DateTime)},
	"TripUpdates_T":      {Type: "protobuf", Filename: "TripUpdates_T.pb", LastChecked: time.Unix(0, 0), ModTime: time.Unix(0, 0), Vehicle: "T", FormattedTime: time.Unix(0, 0).Format(time.DateTime)},
	"ServiceAlerts_A":    {Type: "protobuf", Filename: "ServiceAlerts_A.pb", LastChecked: time.Unix(0, 0), ModTime: time.Unix(0, 0), Vehicle: "A", FormattedTime: time.Unix(0, 0).Format(time.DateTime)},
	"ServiceAlerts_T":    {Type: "protobuf", Filename: "ServiceAlerts_T.pb", LastChecked: time.Unix(0, 0), ModTime: time.Unix(0, 0), Vehicle: "T", FormattedTime: time.Unix(0, 0).Format(time.DateTime)},
	"VehiclePositions_A": {Type: "protobuf", Filename: "VehiclePositions_A.pb", LastChecked: time.Unix(0, 0), ModTime: time.Unix(0, 0), Vehicle: "A", FormattedTime: time.Unix(0, 0).Format(time.DateTime)},
	"VehiclePositions_T": {Type: "protobuf", Filename: "VehiclePositions_T.pb", LastChecked: time.Unix(0, 0), ModTime: time.Unix(0, 0), Vehicle: "T", FormattedTime: time.Unix(0, 0).Format(time.DateTime)},
	"GTFS_KRK_A":         {Type: "zip", Filename: "GTFS_KRK_A.zip", LastChecked: time.Unix(0, 0), ModTime: time.Unix(0, 0), Vehicle: "A", FormattedTime: time.Unix(0, 0).Format(time.DateTime)},
	"GTFS_KRK_T":         {Type: "zip", Filename: "GTFS_KRK_T.zip", LastChecked: time.Unix(0, 0), ModTime: time.Unix(0, 0), Vehicle: "T", FormattedTime: time.Unix(0, 0).Format(time.DateTime)},
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
		// write gtfsFiles as json to file
		err = json.NewEncoder(file).Encode(gtfsFiles)

		if err != nil {
			return fmt.Errorf("problem encoding gtfs files to file %s, %v", file.Name(), err)
		}

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
	// add close func to close file

	return nil, errors.New("not implemented")
}
