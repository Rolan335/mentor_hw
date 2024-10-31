package unpack_refactor

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrUseDigits = errors.New("error. Cannot use digits")

func Pack(str string) (string, error) {
	runes := []rune(str)
	var buf strings.Builder
	count := 1
	for i, v := range runes {
		if unicode.IsDigit(v) {
			return "", ErrUseDigits
		}
		if i != len(runes)-1 && count >= 9 {
			buf.WriteString(fmt.Sprintf("%c%d", v, count))
			count = 1
			continue
		}
		if i != len(runes)-1 && v == runes[i+1] {
			count++
			continue
		}
		buf.WriteRune(v)
		buf.WriteString(strconv.Itoa(count))
		count = 1
	}
	return buf.String(), nil
}
