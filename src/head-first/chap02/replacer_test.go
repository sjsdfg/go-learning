package chap02

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRegex(t *testing.T) {
	broken := "G# r#cks"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	assert.Equal(t, "Go rocks", fixed)
}
