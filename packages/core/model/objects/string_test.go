package objects_test

import (
	"github.com/stretchr/testify/assert"
	"pdfeg-core/model/objects"
	"testing"
)

func TestIndirectLiteralString(t *testing.T) {
	reference, err := objects.NewReference(1, 0)
	assert.Nil(t, err)
	literalString := objects.LiteralString{Reference: reference, Value: "any"}
	assert.Equal(t, "any", literalString.String())
	assert.Equal(t, "any", literalString.Value)
}

func TestNewHexadecimalString_HappyPath(t *testing.T) {
	reference, err := objects.NewReference(1, 0)
	assert.Nil(t, err)
	hexadecimalString := objects.HexadecimalString{Reference: reference, Value: "any"}
	assert.Nil(t, err)
	assert.Equal(t, "any", hexadecimalString.String())
	assert.Equal(t, "any", hexadecimalString.Value)
}

func TestLiteralString_TypeIndicators(t *testing.T) {
	literalString := objects.LiteralString{Value: "any"}
	assert.True(t, literalString.IsLiteral())
	assert.False(t, literalString.IsHexadecimal())
}

func TestHexadecimalString_TypeIndicators(t *testing.T) {
	hexString := objects.HexadecimalString{Value: "any"}
	assert.True(t, hexString.IsHexadecimal())
	assert.False(t, hexString.IsLiteral())
}

func testLiteralStringAsASCIIBytes(t *testing.T, s string, expectedBytes []byte) {
	literalString := objects.LiteralString{Value: s}
	actualBytes, err := literalString.AsASCIIBytes()
	assert.Nil(t, err)
	assert.Equal(t, expectedBytes, actualBytes)
}

func TestLiteralString_AsASCIIBytes(t *testing.T) {
	testLiteralStringAsASCIIBytes(t, "any", []byte("(any)"))
	testLiteralStringAsASCIIBytes(t, "\n", []byte("(\n)"))
}

func testHexadecimalStringAsASCIIBytes(t *testing.T, s string, expectedBytes []byte) {
	hexadecimalString := objects.HexadecimalString{Value: s}
	actualBytes, err := hexadecimalString.AsASCIIBytes()
	assert.Nil(t, err)
	assert.Equal(t, expectedBytes, actualBytes)
}

func TestHexadecimalString_AsASCIIBytes(t *testing.T) {
	testHexadecimalStringAsASCIIBytes(t, "any", []byte("<616e79>"))
	testHexadecimalStringAsASCIIBytes(t, " ", []byte("<20>"))
	testHexadecimalStringAsASCIIBytes(t, "\n", []byte("<0a>"))
	testHexadecimalStringAsASCIIBytes(t, "\r", []byte("<0d>"))
	testHexadecimalStringAsASCIIBytes(t, "\t", []byte("<09>"))
}
