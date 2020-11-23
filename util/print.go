package util

import (
	"fmt"
	"log"
)

// PrintError Prints An Error Message
func PrintError(data string) {
	log.Fatalln("[ERROR]: " + data)
}

// PrintInfo Prints An Info Message
func PrintInfo(data string) {
	fmt.Println("[INFO]: " + data)
}

// PrintInit Prints Program Init Info
func PrintInit(name string, data string) {
	fmt.Println("[" + name + "]: " + data)
}
