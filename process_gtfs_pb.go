package main

import (
	//gtfsrealtime "asdf/gtfs-realtime"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
	gtfsrealtime "gtfs-realtime/proto"
)

func mmain() {
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
	//tripUpdate := &gtfs_realtime.FeedMessage{}
	tripUpdate := &gtfsrealtime.FeedMessage{}

	// Deserializacja danych z pliku protobuffer
	if err := proto.Unmarshal(data, tripUpdate); err != nil {
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
