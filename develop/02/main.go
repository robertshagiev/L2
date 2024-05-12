package main

/*
Задача на распаковку

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func stringUnpacking(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	var result strings.Builder
	var lastChar rune
	var multiplier string
	var escaped bool

	for _, char := range str {
		if char == '\\' && !escaped {
			escaped = true
			continue
		}

		if unicode.IsDigit(char) && !escaped {
			multiplier += string(char)
			continue
		}

		if len(multiplier) > 0 {
			count, _ := strconv.Atoi(multiplier)
			if count == 0 {
				resultString := result.String()
				resultString = resultString[:len(resultString)-1]
				result.Reset()
				result.WriteString(resultString)
			} else {
				result.WriteString(strings.Repeat(string(lastChar), count-1))
			}
			multiplier = ""
		}

		if !escaped || char != '\\' {
			result.WriteRune(char)
			lastChar = char
		}
		escaped = false
	}

	if len(multiplier) > 0 {
		return "", errors.New("incorrect string")
	}

	return result.String(), nil
}

func main() {
	tests := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`}

	for _, test := range tests {
		unpacked, err := stringUnpacking(test)
		if err != nil {
			fmt.Printf("Ошибка при распаковке строки '%s': %s\n", test, err)
		} else {
			fmt.Printf("Распакованная строка '%s': '%s'\n", test, unpacked)
		}
	}
}
