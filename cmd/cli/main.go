package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/protobuf/proto"
	gtfs "gtfs-proxy"
	"os"
)

// how to generate gtfs-realtime.pb.go
// wget https://raw.githubusercontent.com/google/transit/master/gtfs-realtime/proto/gtfs-realtime.proto
// add option go_package = "github.com/fyrtax/gtfs_proxy"; to gtfs-realtime.proto
// protoc --go_out=. --go_opt=paths=source_relative gtfs-realtime.proto

func main() {
	help := flag.Bool("h", false, "Display help")
	flag.Usage = printUsage
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	args := flag.Args()
	if len(args) < 1 || len(args) > 2 {
		fmt.Println("Invalid number of arguments")
		flag.Usage()
		os.Exit(1)
	}

	inputFile, outputFile := args[0], ""
	if len(args) == 2 {
		outputFile = args[1]
	}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		exitWithError("Error reading input file", err)
	}

	tripUpdate := &gtfs.FeedMessage{}
	if err := proto.Unmarshal(data, tripUpdate); err != nil {
		exitWithError("Error deserializing", err)
	}

	if outputFile == "" {
		if err := json.NewEncoder(os.Stdout).Encode(tripUpdate); err != nil {
			exitWithError("Error encoding to JSON", err)
		}
		return
	}

	file, err := os.Create(outputFile)
	if err != nil {
		exitWithError("Error creating output file", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			exitWithError("Error closing output file", err)
		}
	}(file)

	if err := json.NewEncoder(file).Encode(tripUpdate); err != nil {
		exitWithError("Error encoding to JSON", err)
	}
}

func printUsage() {
	fmt.Println("Converts a GTFS-realtime Protocol Buffer file to JSON")
	fmt.Printf("Usage: %s [options] <input_file> [output_file]\n", os.Args[0])
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
	fmt.Printf("Arguments:\n")
	fmt.Printf("\tinput_file\tPath to the .pb GTFS file to convert\n")
	fmt.Printf("\toutput_file\t(Optional) Path to the output JSON file. If not provided, output is printed to stdout\n")
}

func exitWithError(message string, err error) {
	fmt.Printf("%s: %v\n", message, err)
	os.Exit(1)
}
