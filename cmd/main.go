package main

import (
	"dz4_go_example/internal/parser"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileName, err := parseFlags()
	if err != nil {
		log.Fatal("parse flags: ", err)
	}

	if err := run(fileName); err != nil {
		log.Fatal("run: ", err)
	}
}

func run(fileName string) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	emailText := string(content)

	fmt.Println("CountParts")
	partsCount := parser.CountParts(emailText)
	fmt.Println("Количество партов:", partsCount)

	fmt.Println("ParseEmail")
	partsCount, partContents := parser.ParseEmail(emailText)
	fmt.Println("Количество партов:", partsCount)
	fmt.Println("Содержимое:")
	for i, content := range partContents {
		fmt.Printf("Парт %d: %s\n", i+1, content)
	}
	return nil
}

func parseFlags() (fileName string, err error) {
	flag.StringVar(&fileName, "FILE", "", "filenmae (required)")

	flag.Parse()

	if fileName == "" {
		return "", fmt.Errorf("missing required argument FILE")
	}

	return
}
