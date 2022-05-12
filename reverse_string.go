package reverse_string

import (
	"github.com/kyokomi/emoji/v2"
	"strings"
)

func ReverseString(input string) (output string) {
	var res strings.Builder
	input = emoji.Sprint(input)
	input = strings.TrimSpace(input)
	inputLen := len(input)
	resRunes := make([]rune, inputLen)

	for i, v := range []rune(input) {
		resRunes[inputLen-1-i] = v
	}

	for _, v := range resRunes {
		res.WriteRune(v)
	}

	output = res.String()

	return output
}
