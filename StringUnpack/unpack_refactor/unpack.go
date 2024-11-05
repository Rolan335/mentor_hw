package unpack_refactor

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrStartWithDigit = errors.New("cannot start with digit")
var ErrContainNum = errors.New("cannot contain numbers >9 or 00, 01, 02, etc")
var ErrEscLetters = errors.New("cannot escape letter in escaping mode")

var vInt int

func Unpack(str string, isRaw bool) (string, error) {
	if str == "" {
		return "", nil
	}
	if rune(str[0]) > '0' && rune(str[0]) < '9' {
		return "", ErrStartWithDigit
	}
	if isRaw {
		return rawMode(str)
	}
	return basicMode(str)
}

func basicMode(str string) (string, error) {
	var buf strings.Builder
	runes := []rune(str)
	for i, v := range runes {
		if i != len(runes)-1 && runes[i+1] == '0' || runes[i] == '0' {
			continue
		}
		//utf8 get rune by index check.
		if unicode.IsDigit(v) && unicode.IsDigit(runes[i-1]) {
			return "", ErrContainNum
		}
		if unicode.IsDigit(v) && v != '0' {
			vInt = int(v - '0' - 1)
			buf.WriteString(strings.Repeat(string(runes[i-1]), vInt))
			continue
		}
		buf.WriteRune(v)
	}
	return strconv.Quote(buf.String()), nil
}

func rawMode(str string) (string, error) {
	str = strings.ReplaceAll(str, `\\`, `+`)
	runes := []rune(str)
	var buf strings.Builder
	for i, v := range runes {
		if i > 0 && runes[i-1] == '\\' && unicode.IsLetter(v) {
			return "", ErrEscLetters
		}
		if i > 1 && runes[i-2] != '\\' && unicode.IsDigit(runes[i-1]) && unicode.IsDigit(v) {
			return "", ErrContainNum
		}
		if i != len(runes)-1 && runes[i+1] == '0' || runes[i] == '0' && runes[i-1] != '\\' {
			continue
		}
		if unicode.IsDigit(v) && runes[i-1] != '\\' {
			vInt = int(v - '0' - 1)
			buf.WriteString(strings.Repeat(string(runes[i-1]), vInt))
			continue
		}
		buf.WriteRune(v)
	}
	str = strings.ReplaceAll(buf.String(), `\`, ``)
	str = strings.ReplaceAll(str, "+", `\`)
	return str, nil
}
