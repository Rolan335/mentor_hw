package unpack_refactor

import (
	"errors"
	"strconv"
	"strings"
)

var ErrUseDigits = errors.New("Cannot use digits")

func Pack(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	var buf strings.Builder
	count := 0
	lastSymbol := rune(str[0])
	for _, v := range str {
		if v >= '0' && v <= '9' {
			return "", ErrUseDigits
		}
		if count >= 9 {
			buf.WriteRune(lastSymbol)
			buf.WriteString(strconv.Itoa(count))
			count = 1
			continue
		}
		if lastSymbol == v {
			count++
			continue
		}
		buf.WriteRune(lastSymbol)
		buf.WriteString(strconv.Itoa(count))
		lastSymbol = v
		count = 1
	}
	buf.WriteRune(lastSymbol)
	buf.WriteString(strconv.Itoa(count))
	return buf.String(), nil
}