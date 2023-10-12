package samples

import (
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestTrim(t *testing.T) {
	assert.Equal(t, strings.TrimSpace("  a  "), "a")
	assert.Equal(t, strings.TrimSpace("a "), "a")
	assert.Equal(t, strings.TrimSpace("  a"), "a")
	assert.Equal(t, strings.TrimSpace(" a b "), "a b")
}
