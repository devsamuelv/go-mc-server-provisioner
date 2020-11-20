package util

import (
	"log"
	"fmt"
)

// PrintError Prints An Error Message
func PrintError(data string) {
	log.Fatalln("[ERROR]: " + data)
}

// PrintInfo Prints An Info Message
func PrintInfo(data string) {
	fmt.Println("[INFO]: " + data)
}
