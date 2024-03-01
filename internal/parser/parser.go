package parser

import (
    "strings"
)

// CountParts принимает текст письма в формате SMTP и возвращает количество партов
func CountParts(emailText string) int {
    // Разделяем письмо на части по разделителю "--Boundary_"
    parts := strings.Split(emailText, "--Boundary_")
    // Первая часть письма часто пустая, поэтому проверяем, является ли первая часть пустой строкой
    if len(parts[0]) == 0 {
        return len(parts) - 1 // Количество частей равно количеству разделителей минус один
    }
    return len(parts) // Если первая часть не пустая, количество частей равно количеству разделителей
}

// ParseEmail принимает текст письма в формате SMTP и парсит его, возвращает количество партов и их содержимое
func ParseEmail(emailText string) (int, []string) {
    parts := strings.Split(emailText, "--Boundary_")
    // Инициализируем массив для содержимого частей письма
    partContents := make([]string, 0)
    // Первая часть письма часто пустая, поэтому начинаем с индекса 1
    for i := 1; i < len(parts); i++ {
        // Разбиваем каждую часть письма на строки
        lines := strings.Split(strings.TrimSpace(parts[i]), "\n")
        // Если в части письма есть хотя бы одна строка, которая не является пустой, то это содержимое части
        if len(lines) > 0 && len(strings.TrimSpace(lines[0])) > 0 {
            // Содержимое части - все строки после первой пустой строки
            content := ""
            for j := 1; j < len(lines); j++ {
                if len(strings.TrimSpace(lines[j])) > 0 {
                    content += lines[j] + "\n"
                }
            }
            // Добавляем содержимое части в массив
            partContents = append(partContents, content)
        }
    }
    // Количество частей равно количеству элементов в массиве содержимого
    return len(partContents), partContents
}
