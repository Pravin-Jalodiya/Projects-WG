package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

// StreamJSONObjects takes a channel, a filename, and a reflect.Type representing the struct type,
// reads the JSON file, unmarshals it into a slice of the given struct type, and sends each item over the channel.
func StreamJSONObjects(ch chan<- any, fileName string, structType reflect.Type) {
	// Read the entire file into memory
	file, err := os.Open(fileName)
	content, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
		close(ch)
		return
	}
	//fmt.Println(structType)
	// Unmarshal the JSON content into a slice of the specified struct type
	var items []Course
	err = json.Unmarshal(content, &items)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		close(ch)
		return
	}

	// Send each item over the channel
	for _, item := range items {
		ch <- item
	}

	close(ch)
}

func mainT() {
	// Example usage
	ch := make(chan any, 10) // Buffered channel to hold 10 items
	structType := reflect.TypeOf(Course{})
	go StreamJSONObjects(ch, "C:/Projects-WG/json/courses.json", structType)

	// Receive items from the channel
	for item := range ch {
		fmt.Println(item)
	}
}

// MyStruct represents the struct type expected in the JSON file.
// Replace this with the actual struct type you're working with.
type Course struct {
	CID     int
	Title   string
	Modules []Module
}

type Module struct {
	MID       float32
	Title     string
	Weightage int
}
