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

func Test_NewIndirectInteger(t *testing.T) {
	actual, err := objects.NewIndirectInteger(1, 0, 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), actual.Value)
	assert.Equal(t, 1, actual.ObjectNumber)
	assert.Equal(t, 0, actual.GenerationNumber)
}

func Test_NewIndirectInteger_validates(t *testing.T) {
	actual, err := objects.NewIndirectInteger(0, 0, 1)
	assert.Nil(t, actual)
	assert.Error(t, err)
}

func assertIntegerAsBytes(t *testing.T, expected string, value int64) {
	bytes, err := (&objects.Integer{Value: value}).AsBytes()
	assert.Nil(t, err)
	assert.Equal(t, []byte(expected), bytes)
}

func TestIntegerAsBytes(t *testing.T) {
	assertIntegerAsBytes(t, "1", 1)
	assertIntegerAsBytes(t, "-1", -1)
	assertIntegerAsBytes(t, "0", 0)
	assertIntegerAsBytes(t, "4294967296", 1<<32)
	assertIntegerAsBytes(t, "4611686018427387904", 1<<62)
	assertIntegerAsBytes(t, "-4294967296", -1<<32)
	assertIntegerAsBytes(t, "-4611686018427387904", -1<<62)
}

func TestReal(t *testing.T) {
	r := objects.Real{Value: 1}
	assert.False(t, r.IsInteger())
	assert.True(t, r.IsReal())
}

func Test_NewIndirectReal(t *testing.T) {
	actual, err := objects.NewIndirectReal(1, 0, 1)
	assert.Nil(t, err)
	assert.Equal(t, float64(1), actual.Value)
	assert.Equal(t, 1, actual.ObjectNumber)
	assert.Equal(t, 0, actual.GenerationNumber)
}

func Test_NewIndirectReal_validates(t *testing.T) {
	actual, err := objects.NewIndirectReal(0, 0, 1)
	assert.Nil(t, actual)
	assert.Error(t, err)
}

func assertRealAsBytes(t *testing.T, expected string, real float64) {
	bytes, err := (&objects.Real{Value: real}).AsBytes()
	assert.Nil(t, err)
	assert.Equal(t, []byte(expected), bytes)
}

func TestRealAsBytes(t *testing.T) {
	assertRealAsBytes(t, "1", 1)
	assertRealAsBytes(t, "1", 1.0)
	assertRealAsBytes(t, "1.01", 1.010)
	assertRealAsBytes(t, "-1", -1)
	assertRealAsBytes(t, "-1", -1.0)
	assertRealAsBytes(t, "-1.09", -1.090)
	assertRealAsBytes(t, "0", 0)
	assertRealAsBytes(t, "0", 0.0)
	assertRealAsBytes(t, "9.9", 009.900)
	assertRealAsBytes(t, "4294967296", 1<<32)
	assertRealAsBytes(t, "4294967296.23", 4294967296.23)
	assertRealAsBytes(t, "18446744073709552000", 1<<64)
	assertRealAsBytes(t, "-4294967296", -1<<32)
	assertRealAsBytes(t, "-4294967296.23", -4294967296.23)
	assertRealAsBytes(t, "-18446744073709552000", -1<<64)
}
