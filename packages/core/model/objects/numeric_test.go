package objects_test

import (
	"github.com/stretchr/testify/assert"
	"pdfeg-core/model/objects"
	"testing"
)

func TestInteger(t *testing.T) {
	i := objects.Integer{Value: 1}
	assert.True(t, i.IsInteger())
	assert.False(t, i.IsReal())
}

func TestIndirectInteger(t *testing.T) {
	reference, err := objects.NewReference(1,0)
	assert.Nil(t, err)
	actual := objects.Integer{Reference: reference, Value: 1}
	assert.Equal(t, int64(1), actual.Value)
	assert.NotNil(t, actual.Label())
	assert.Equal(t, 1, actual.ObjectNumber)
	assert.Equal(t, 0, actual.GenerationNumber)
}

func assertIntegerAsASCIIBytes(t *testing.T, expected string, value int64) {
	bytes, err := (&objects.Integer{Value: value}).AsASCIIBytes()
	assert.Nil(t, err)
	assert.Equal(t, []byte(expected), bytes)
}

func TestInteger_AsASCIIBytes(t *testing.T) {
	assertIntegerAsASCIIBytes(t, "1", 1)
	assertIntegerAsASCIIBytes(t, "-1", -1)
	assertIntegerAsASCIIBytes(t, "0", 0)
	assertIntegerAsASCIIBytes(t, "4294967296", 1<<32)
	assertIntegerAsASCIIBytes(t, "4611686018427387904", 1<<62)
	assertIntegerAsASCIIBytes(t, "-4294967296", -1<<32)
	assertIntegerAsASCIIBytes(t, "-4611686018427387904", -1<<62)
}

func TestReal(t *testing.T) {
	r := objects.Real{Value: 1}
	assert.False(t, r.IsInteger())
	assert.True(t, r.IsReal())
}

func TestIndirectReal(t *testing.T) {
	reference, err := objects.NewReference(1, 0)
	assert.Nil(t, err)
	actual := objects.Real{Reference: reference, Value: 1}
	assert.Equal(t, float64(1), actual.Value)
	assert.NotNil(t, actual.Label())
	assert.Equal(t, 1, actual.ObjectNumber)
	assert.Equal(t, 0, actual.GenerationNumber)
}

func assertRealAsASCIIBytes(t *testing.T, expected string, real float64) {
	bytes, err := (&objects.Real{Value: real}).AsASCIIBytes()
	assert.Nil(t, err)
	assert.Equal(t, []byte(expected), bytes)
}

func TestReal_AsASCIIBytes(t *testing.T) {
	assertRealAsASCIIBytes(t, "1", 1)
	assertRealAsASCIIBytes(t, "1", 1.0)
	assertRealAsASCIIBytes(t, "1.01", 1.010)
	assertRealAsASCIIBytes(t, "-1", -1)
	assertRealAsASCIIBytes(t, "-1", -1.0)
	assertRealAsASCIIBytes(t, "-1.09", -1.090)
	assertRealAsASCIIBytes(t, "0", 0)
	assertRealAsASCIIBytes(t, "0", 0.0)
	assertRealAsASCIIBytes(t, "9.9", 009.900)
	assertRealAsASCIIBytes(t, "4294967296", 1<<32)
	assertRealAsASCIIBytes(t, "4294967296.23", 4294967296.23)
	assertRealAsASCIIBytes(t, "18446744073709552000", 1<<64)
	assertRealAsASCIIBytes(t, "-4294967296", -1<<32)
	assertRealAsASCIIBytes(t, "-4294967296.23", -4294967296.23)
	assertRealAsASCIIBytes(t, "-18446744073709552000", -1<<64)
}
