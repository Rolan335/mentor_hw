package unpack

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string, isRaw bool) (string, error) {
	if str == "" {
		return "", nil
	}
	if isRaw {
		str = strings.ReplaceAll(str, `\\`, `+`)
		str = strings.ReplaceAll(str, `\`, "_")
	}
	runes := []rune(str)
	if unicode.IsDigit(runes[0]) {
		return "", errors.New("error. String cannot start with digit")
	}
	var buf strings.Builder
	for i, v := range runes {
		if isRaw {
			//runes[i-1] panics if checks first letter
			if i != 0 && i != len(runes)-1 {
				if unicode.IsDigit(v) && runes[i-1] != '_' && unicode.IsDigit(runes[i+1]) {
					return "", errors.New("error. String cannot contain numbers >9")
				}
				if unicode.IsLetter(v) && runes[i-1] == '_' {
					return "", errors.New("error. cannot escape letter in escaping mode")
				}
			}
			if unicode.IsDigit(v) && runes[i-1] != '_' && int(v-49) >= 0 {
				buf.WriteString(strings.Repeat(string(runes[i-1]), int(v-49)))
				continue
			}
			if i != len(runes)-1 {
				if runes[i+1] == '0' {
					continue
				}
			}
			if runes[i] == '0' && runes[i-1] != '_' {
				continue
			}
			buf.WriteString(string(v))
			continue
		}
		if unicode.IsDigit(v) {
			if unicode.IsDigit(runes[i-1]) {
				return "", errors.New("error. String cannot contain numbers >9")
			}
			fmt.Println(string(v))
			if int(v-49) > 0 {
				buf.WriteString(strings.Repeat(string(runes[i-1]), int(v-49)))
				continue
			}
			continue
		}
		// Is there better way to check if out of range
		if i != len(runes)-1 {
			if runes[i+1] == '0' {
				continue
			}
		}
		if runes[i] == '0' {
			continue
		}
		buf.WriteString(string(v))
	}
	str = strings.ReplaceAll(buf.String(), "_", "")
	str = strings.ReplaceAll(str, "+", `\`)
	if isRaw {
		return str, nil
	}
	return strconv.Quote(str), nil
}
