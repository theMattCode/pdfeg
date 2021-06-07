package objects_test

import (
	"github.com/stretchr/testify/assert"
	o "pdfeg-core/model/objects"
	"testing"
)

var (
	expectedTrueBytes  = []byte{0x74, 0x72, 0x75, 0x65}
	expectedFalseBytes = []byte{0x66, 0x61, 0x6c, 0x73, 0x65}
)

func TestBooleanTrue(t *testing.T) {
	assertBoolean(t, &o.True, true)
}

func TestBooleanFalse(t *testing.T) {
	assertBoolean(t, &o.False, false)
}

func Test_NewIndirectBoolean_True(t *testing.T) {
	actual, err := o.NewIndirectBoolean(1, 0, true)
	assert.Nil(t, err)
	assertBoolean(t, actual, true)
	assert.Equal(t, 1, actual.ObjectNumber)
	assert.Equal(t, 0, actual.GenerationNumber)
}

func Test_NewIndirectBoolean_False(t *testing.T) {
	actual, err := o.NewIndirectBoolean(1, 0, false)
	assert.Nil(t, err)
	assertBoolean(t, actual, false)
	assert.Equal(t, 1, actual.ObjectNumber)
	assert.Equal(t, 0, actual.GenerationNumber)
}

func Test_NewIndirectBoolean_validates(t *testing.T) {
	actual, err := o.NewIndirectBoolean(0, 0, true)
	assert.Nil(t, actual)
	assert.Error(t, err)
}

func assertBoolean(t *testing.T, actual *o.Boolean, expectedNative bool) {
	assert.Equal(t, actual.Value, expectedNative)
	actualBytes, err := actual.AsBytes()
	assert.Nil(t, err)
	if expectedNative {
		assert.Equal(t, expectedTrueBytes, actualBytes)
	} else {
		assert.Equal(t, expectedFalseBytes, actualBytes)
	}
}
