package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	gtfs "gtfs-proxy"
	"log"
	"os"
)

// how to generate gtfs-realtime.pb.go
// wget https://raw.githubusercontent.com/google/transit/master/gtfs-realtime/proto/gtfs-realtime.proto
// add option go_package = "github.com/fyrtax/gtfs_proxy"; to gtfs-realtime.proto
// protoc --go_out=. --go_opt=paths=source_relative gtfs-realtime.proto

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Podaj nazwę pliku jako argument.")
		return
	}

	fileName := os.Args[1]

	// Otwarcie pliku
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Błąd podczas odczytu pliku:", err)
	}

	// Inicjalizacja struktury z pliku protobuffer
	tripUpdate := gtfs.FeedMessage{}

	// Deserializacja danych z pliku protobuffer
	if err := proto.Unmarshal(data, &tripUpdate); err != nil {
		log.Fatal("Błąd podczas deserializacji danych:", err)
	}

	// Tutaj można przekonwertować strukturę protobuffer na JSON używając MarshalJSON lub innych metod dostępnych w Twoich strukturach

	// fmt.Println("Deserializacja zakończona pomyślnie!")

	//enc := json.NewEncoder(os.Stdout)
	//enc.Encode(tripUpdate)
	//
	//log.Println(tripUpdate)

	//replace .pb with .json
	fileName = fileName[:len(fileName)-2] + "json"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	_ = file.Truncate(0)

	if err != nil {
		log.Fatal("Cannot open output file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Cannot close output file:", err)
		}
	}(file)

	// save to json file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(tripUpdate)
	if err != nil {
		log.Fatal("Cannot encode to JSON:", err)
	}

	fmt.Println("Zapisano do pliku", fileName)
}
