package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_isWhitespace(t *testing.T) {
	runes := []rune{SP, CR, LF, FF, NUL, HT}
	for _, r := range runes {
		t.Run(fmt.Sprintf("true for rune %v", r),
			func(t *testing.T) {
				assert.True(t, isWhitespace(r))
			},
		)
	}

	runes = []rune{1, 2, 8, 14, 31, 33, 127}
	for _, r := range runes {
		t.Run(fmt.Sprintf("false for rune %v", r),
			func(t *testing.T) {
				assert.False(t, isWhitespace(r))
			},
		)
	}
}

func TestScanner_FindStart(t *testing.T) {
	type TestCase struct {
		input             string
		expectedByteCount int
		expectedVersion   string
		additionalInfo    string
	}

	testCases := []TestCase{
		{"%PDF-1.0\n%\x80\x81\x7F\x7F\n", 15, "%PDF-1.0", "at beginning"},
		{"%PDF-1.1\n%\x80\x81\x7E\x7F\n", 15, "%PDF-1.1", "at beginning"},
		{"%PDF-1.2\n%\x80\x81\x7E\x7F\n", 15, "%PDF-1.2", "at beginning"},
		{"%PDF-1.3\n%\x80\x81\x7E\x7F\n", 15, "%PDF-1.3", "at beginning"},
		{"%PDF-1.4\n%\x80\x81\x7E\x7F\n", 15, "%PDF-1.4", "at beginning"},
		{"%PDF-1.5\n%\x80\x81\x7E\x7F\n", 15, "%PDF-1.5", "at beginning"},
		{"%PDF-1.6\n%\x80\x81\x7E\x7F\n", 15, "%PDF-1.6", "at beginning"},
		{"%PDF-1.7\n%\x80\x81\x7E\x7F\n", 15, "%PDF-1.7", "at beginning"},
		{"%PDF-2.0\n%\x80\x81\x7E\x7F\n", 15, "%PDF-2.0", "at beginning"},
		{"there \n are\n some\r\n inappropriate lines\n%PDF-2.0\n%\x80\x81\x7E\x7F\n", 55, "%PDF-2.0", "with prepended lines"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v %v", testCase.expectedVersion, testCase.additionalInfo),
			func(t *testing.T) {
				rawReader := strings.NewReader(testCase.input)
				scanner := NewScanner(rawReader)
				version, pos, err := scanner.FindStart()
				assert.NoError(t, err)
				assert.Equal(t, uint64(testCase.expectedByteCount), pos)
				assert.Equal(t, testCase.expectedVersion, version)
			},
		)
	}
}
