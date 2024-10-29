package unpack

import (
	"errors"
	"fmt"
	"strconv"
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
				buf.WriteString(fmt.Sprintf("%c%d", v, count))
				count = 1
				continue
			}
			if v == runes[i+1] {
				count++
				continue
			}
		}
		//fmt.sprintf slow better use strconv
		buf.WriteRune(v)
		buf.WriteString(strconv.Itoa(count))
		count = 1
	}
	return buf.String(), nil
}
