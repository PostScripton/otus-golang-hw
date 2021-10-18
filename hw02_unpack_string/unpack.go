package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var (
		newStr  strings.Builder
		temp    string
		escaped bool
	)

	for _, s := range str {
		if temp == `\` && !escaped {
			temp, escaped = string(s), true
			continue
		}

		if unicode.IsDigit(s) {
			if temp == "" {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(string(s))
			newStr.WriteString(strings.Repeat(temp, count))
			temp, escaped = "", false
		} else {
			if temp != "" {
				newStr.WriteString(temp)
				escaped = false
			}

			temp = string(s)
		}
	}

	newStr.WriteString(temp)
	return newStr.String(), nil
}
