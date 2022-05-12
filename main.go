package main

import (
	"log"
	"os"
)

func main() {
}

func GetHelp() string {
	contents, err := os.ReadFile("help.txt")
	output := string(contents)
	if err != nil {
		log.Fatal(err)
	}
	return output
}
