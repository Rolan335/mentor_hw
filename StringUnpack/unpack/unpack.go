package unpack

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// var buf strings.Builder
// runes := []rune(str)
// for i, v := range runes {
// 	if toInt, err := strconv.Atoi(string(v)); err == nil {
// 		if i != len(runes)-1 {
// 			if _, err := strconv.Atoi(string(runes[i+1])); err == nil {
// 				return ""
// 			}
// 		}
// 		buf.WriteString(strings.Repeat(string(runes[i-1]), toInt-1))
// 	} else {
// 		buf.WriteRune(v)
// 	}
// }
// return buf.String()

// type FormattedString struct{
// 	str string
// 	isRaw bool
// }

func Unpack(str string, isRaw bool) any {
	if str == "" {
		return ""
	}
	if isRaw {
		str = strings.ReplaceAll(str, `\\`, `+`)
		str = strings.ReplaceAll(str, `\`, "_")
	}
	fmt.Println(str)
	runes := []rune(str)
	if unicode.IsDigit(runes[0]) {
		return "wrong string"
	}
	var buf strings.Builder
	for i, v := range runes {
		if isRaw {
			if i != 0 {
				if unicode.IsLetter(v) && runes[i-1] == '_' {
					return "wrong string"
				}
			}
			if unicode.IsDigit(v) && runes[i-1] != '_' {
				buf.WriteString(strings.Repeat(string(runes[i-1]), int(v-49)))
				continue
			}
			//runes[i-1] panics if checks first letter
			buf.WriteString(string(v))
			continue
		}
		if unicode.IsDigit(v) {
			if unicode.IsDigit(runes[i-1]) {
				return "wrong string"
			}
			buf.WriteString(strings.Repeat(string(runes[i-1]), int(v-49)))
			continue
		}
		buf.WriteString(string(v))
		fmt.Println(buf.String())
	}
	str = strings.ReplaceAll(buf.String(), "_", "")
	str = strings.ReplaceAll(str, "+", `\`)
	if isRaw {
		return str
	}
	return strconv.Quote(str)
}

func unpackRaw(runes []rune) any {
	strSlice := []string{}
	for i, v := range runes {
		fmt.Print(string(v))
		if string(v) == `\` && (string(runes[i+1]) == `\` || unicode.IsDigit(runes[i+1])) {
			fmt.Print(true)
			fmt.Println()
			strSlice = append(strSlice, string(v), string(runes[i+1]))
		} else {
			strSlice = append(strSlice, string(v))
			fmt.Println()
		}
	}
	fmt.Println(strSlice)
	return strSlice
}

func ProcessString(s string) (string, error) {
	result := ""
	skipNext := false

	for i, r := range s {
		if skipNext {
			// Если предыдущий символ был слэш, добавляем текущий символ (если он не вне диапазона)
			if unicode.IsDigit(r) || r == '\\' {
				result += string(r)
				skipNext = false // Сбрасываем флаг
			} else {
				return "", fmt.Errorf("некорректная строка на позиции %d", i)
			}
			continue
		}

		if r == '\\' {
			// Устанавливаем флаг для пропуска следующего символа
			skipNext = true
			continue
		}

		result += string(r)
	}

	// Проверяем, остался ли флаг с экранированием
	if skipNext {
		return "", fmt.Errorf("некорректная строка: конечный слэш найден")
	}

	// Заменяем цифры на соответствующее количество цифр
	finalResult := ""
	for _, r := range result {
		if unicode.IsDigit(r) {
			count := int(r - '0')    // Преобразуем символ цифры в число
			finalResult += string(r) // Добавляем саму цифру
			for j := 1; j < count; j++ {
				finalResult += string(r) // Добавляем цифру столько же раз, сколько указано
			}
		} else {
			finalResult += string(r) // Добавляем обычные символы
		}
	}

	return finalResult, nil
}
