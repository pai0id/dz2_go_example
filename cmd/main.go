package main

import (
    "fmt"
    "dz4_go_example/internal/parser"
    "os"
    "io/ioutil"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Корректный вызов программы: go run <filename>")
        os.Exit(1)
    }

    filename := os.Args[1]

    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Printf("Ошибка чтения файла: %v\n", err)
        os.Exit(1)
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
}
