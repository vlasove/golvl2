package helper

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	errInvalidString = errors.New("invalid string literal")
)

const (
	backslashCode = 92
)

// Unpack ...
func Unpack(sample string) (string, error) {
	var res strings.Builder
	runeSlice := []rune(sample)
	for i := 0; i < len(runeSlice); i++ {
		// если i-ый элемент - \
		// то в билдер помещаем следующий за ним
		if runeSlice[i] == backslashCode {
			i++
			res.WriteRune(runeSlice[i])
		} else if n, err := strconv.Atoi(string(runeSlice[i])); err == nil {
			// если i-ый элемент - цифра
			if i == 0 ||
				(i > 0 && unicode.IsDigit(runeSlice[i-1]) &&
					(i > 1 && runeSlice[i-2] != backslashCode)) {
				// если это первая цифра в строке или
				// предыдущий символ - цифра
				//                          и
				//                            пред-предыдущий элемент - не экран
				// q65
				return "", errInvalidString
			}
			// если i-ый элемент цифра и все ок
			// записываем n-1 раз предыдущий элемент
			// n-1 т.к. исходник уже записан
			res.WriteString(strings.Repeat(string(runeSlice[i-1]), n-1))
		} else {
			// если это не цифра или не \
			// просто записываем этот элемент
			res.WriteRune(runeSlice[i])
		}
	}
	return res.String(), nil
}
