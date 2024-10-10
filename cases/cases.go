package cases

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func ToSentence(s string) string {
	sentence := strcase.ToDelimited(s, ' ')
	return strings.ToUpper(sentence[:1]) + sentence[1:]
}
