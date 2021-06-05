package objects

import (
	"github.com/stretchr/testify/assert"
	model2 "pdfeg-core/model"
	"testing"
)

func Test_NewReference_HappyPath(t *testing.T) {
	actual, err := NewReference(1, 0)
	assert.Nil(t, err)
	assert.Equal(t, 1, actual.objectNumber)
	assert.Equal(t, 0, actual.generationNumber)

	actualBytes, err2 := actual.asBytes()
	assert.Nil(t, err2)
	assert.Equal(t, []byte("1 0 R"), actualBytes)
}

func Test_NewReference_validatesObjectNumber(t *testing.T) {
	reference, err := NewReference(0, 0)
	assert.Error(t, err)
	assert.Equal(t, "object number", err.Context)
	assert.Equal(t, model2.PositiveIntegerMessage, err.Message)
	assert.Equal(t, "0", err.Value)
	assert.Nil(t, reference)
}

func Test_NewReference_validatesGenerationNumber(t *testing.T) {
	reference, err := NewReference(1, -1)
	assert.Error(t, err)
	assert.Equal(t, "generation number", err.Context)
	assert.Equal(t, model2.NonNegativeIntegerMessage, err.Message)
	assert.Equal(t, "-1", err.Value)
	assert.Nil(t, reference)
}

func Test_NewNull_HappyPath(t *testing.T) {
	actual := NewNull()
	assert.Nil(t, actual.Reference)

	actualBytes, err := actual.asBytes()
	assert.Nil(t, err)
	assert.Equal(t, nullBytes, actualBytes)
}

func Test_NewIndirectNull_HappyPath(t *testing.T) {
	actual, err := NewIndirectNull(1, 0)
	assert.Nil(t, err)
	assert.Equal(t, 1, actual.objectNumber)
	assert.Equal(t, 0, actual.generationNumber)

	actualBytes, err := actual.asBytes()
	assert.Nil(t, err)
	assert.Equal(t, nullBytes, actualBytes)
}

func Test_NewIndirectNull_ValidatesObjectNumber(t *testing.T) {
	actual, err := NewIndirectNull(0, 0)
	assert.Nil(t, actual)
	assert.Error(t, err)
}
