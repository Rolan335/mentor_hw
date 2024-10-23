package unpack

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func Pack(str string) (string, error) {
	runes := []rune(str)
	var buf strings.Builder
	count := 1
	for i, v := range runes {
		if unicode.IsDigit(v) {
			return "", errors.New("error. Cannot use digits")
		}
		if i != len(runes)-1 {
			if count >= 9 {
				buf.WriteString(fmt.Sprintf("%c%v", v, count))
				count = 1
				continue
			}
			if v == runes[i+1] {
				count++
				continue
			}
		}
		buf.WriteString(fmt.Sprintf("%c%v", v, count))
		count = 1
	}
	return buf.String(), nil
}
