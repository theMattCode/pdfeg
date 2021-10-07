package io_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"pdfeg-core/io"
	"testing"
)

func Test_IsWhitespace(t *testing.T) {
	runes := []rune{io.SP, io.CR, io.LF, io.FF, io.NUL, io.HT}
	for _, r := range runes {
		t.Run(fmt.Sprintf("true for rune %v", r),
			func(t *testing.T) {
				assert.True(t, io.IsWhitespace(r))
			},
		)
	}

	runes = []rune{1, 2, 8, 14, 31, 33, 127}
	for _, r := range runes {
		t.Run(fmt.Sprintf("false for rune %v", r),
			func(t *testing.T) {
				assert.False(t, io.IsWhitespace(r))
			},
		)
	}
}
