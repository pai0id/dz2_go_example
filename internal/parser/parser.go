package parser

import (
    "strings"
)

func CountParts(emailText string) int {
    if emailText == "" {
        return 0
    }

    // Разделители между партами - строка Boundary_...
    boundaries := strings.Count(emailText, "--Boundary_")

    // Если не нашли ни одного разделителя, значит, только одна часть
    if boundaries == 0 {
        return 1
    }

    // Иначе количество партов - количество разделителей плюс один (последняя часть после последнего разделителя)
    return boundaries + 1
}

func ParseEmail(emailText string) (int, []string) {
    if emailText == "" {
        return 0, []string{}
    }

    // Разделители между партами - строка Boundary_...
    boundaries := strings.Split(emailText, "--Boundary_")

    // Счетчик партов
    partsCount := 0

    // Массив содержимого партов
    partContents := make([]string, 0)

    // Проходимся по разделителям
    for _, boundary := range boundaries {
        // Очищаем строку от пробельных символов и символов переноса строки
        cleanBoundary := strings.TrimSpace(boundary)

        // Если строка не пустая, добавляем ее содержимое в массив
        if cleanBoundary != "" {
            partsCount++
            partContents = append(partContents, cleanBoundary)
        }
    }

    return partsCount, partContents
}

