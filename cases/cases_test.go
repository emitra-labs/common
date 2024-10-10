package cases_test

import (
	"testing"

	"github.com/emitra-labs/common/cases"
	"github.com/stretchr/testify/assert"
)

func TestToSentence(t *testing.T) {
	res := cases.ToSentence("foo-bar")

	assert.Equal(t, "Foo bar", res)
}
