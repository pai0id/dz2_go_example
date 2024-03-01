package main

import (
    "fmt"
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

    rmt.Println("CountParts")
    partsCount := CountParts(emailText)
    fmt.Println("Количество партов:", partsCount)

    rmt.Println("ParseEmail")
    partsCount, partContents := ParseEmail(emailText)
    fmt.Println("Количество партов:", partsCount)
    fmt.Println("Содержимое:")
    for i, content := range partContents {
        fmt.Printf("Парт %d: %s\n", i+1, content)
    }
}
