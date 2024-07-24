package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

type ServerSettings struct {
	WorkingDirectory string
	GtfsFilesLogFile string
}

var settings = ServerSettings{
	WorkingDirectory: "gtfs_rf_files",
	GtfsFilesLogFile: "gtfs_files.log",
}

func main() {
	println(time.Now().Format(time.RFC1123))

	var nowTime = time.Now()

	println(nowTime.Unix())
	println(time.Unix(nowTime.Unix(), 0).Format(time.RFC1123))

	// create directory server_files if it does not exist
	err := os.MkdirAll(settings.WorkingDirectory, 0755)

	if err != nil {
		log.Fatalf("problem creating %q directory, %v", settings.WorkingDirectory, err)
	}

	store, closeStore, err := FileSystemGtfsStoreFromFile(settings.WorkingDirectory + "/" + settings.GtfsFilesLogFile)

	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()

	server, err := NewProxyServer(store)

	if err != nil {
		log.Fatalf("problem creating gtfs proxy server %v", err)
	}

	// print server link
	println("Server is running on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
