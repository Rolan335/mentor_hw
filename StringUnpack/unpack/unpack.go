package unpack

import (
	"errors"
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
	}
	runes := []rune(str)
	if unicode.IsDigit(runes[0]) {
		return "", errors.New("error. String cannot start with digit")
	}
	var buf strings.Builder
	for i, v := range runes {
		vInt := int(v - '0' - 1)
		if isRaw {
			//runes[i-1] panics if checks first letter
			if i != 0 && i != len(runes)-1 {
				if unicode.IsDigit(v) && string(runes[i-1]) != `\` && unicode.IsDigit(runes[i+1]) {
					return "", errors.New("error. String cannot contain numbers >9 or 00, 01, 02, etc")
				}
				if (unicode.IsLetter(v) && string(runes[i-1]) == `\`) || (string(v) == `\` && unicode.IsLetter(runes[i+1])){
					return "", errors.New("error. cannot escape letter in escaping mode")
				}
			}
			if unicode.IsDigit(v) && string(runes[i-1]) != `\` && vInt >= 0 {
				buf.WriteString(strings.Repeat(string(runes[i-1]), vInt))
				continue
			}
			if i != len(runes)-1 {
				if runes[i+1] == '0' {
					continue
				}
			}
			if runes[i] == '0' && string(runes[i-1]) != `\` {
				continue
			}
			buf.WriteRune(v)
			continue
		}
		if unicode.IsDigit(v) {
			if unicode.IsDigit(runes[i-1]) {
				return "", errors.New("error. String cannot contain numbers >9 or 00, 01, 02, etc")
			}
			if vInt > 0 {
				buf.WriteString(strings.Repeat(string(runes[i-1]), vInt))
				continue
			}
			continue
		}
		//? Is there better way to check if out of range
		if i != len(runes)-1 {
			if runes[i+1] == '0' {
				continue
			}
		}
		buf.WriteRune(v)
	}
	str = strings.ReplaceAll(buf.String(), `\`, ``)
	str = strings.ReplaceAll(str, "+", `\`)
	if isRaw {
		return str, nil
	}
	return strconv.Quote(str), nil
}
