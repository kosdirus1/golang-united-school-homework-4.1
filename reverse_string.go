package reverse_string

import (
	"github.com/kyokomi/emoji/v2"
	"strings"
)

func ReverseString(input string) (output string) {
	var res strings.Builder
	input = emoji.Sprint(input)
	inputLen := len(input)
	resRunes := make([]string, inputLen)

	for i, v := range []rune(input) {
		resRunes[inputLen-1-i] = string(v)
	}

	for _, v := range resRunes {
		res.WriteString(v)
	}

	output = res.String()

	return output
}
