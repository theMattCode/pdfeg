package objects_test

import (
	"github.com/stretchr/testify/assert"
	"pdfeg-core/model"
	"pdfeg-core/model/objects"
	"testing"
)

type objMock struct {
	labelMock func() *objects.Reference
	asASCIIBytesMock func() ([]byte, error)
}

func (r objMock) Label() *objects.Reference {
	return r.labelMock()
}

func (r objMock) AsASCIIBytes() ([]byte, error) {
	return r.asASCIIBytesMock()
}

func TestNewReference_HappyPath(t *testing.T) {
	actual, err := objects.NewReference(1, 0)
	assert.Nil(t, err)
	assert.Equal(t, 1, actual.ObjectNumber)
	assert.Equal(t, 0, actual.GenerationNumber)

	actualBytes, err2 := actual.AsASCIIBytes()
	assert.Nil(t, err2)
	assert.Equal(t, []byte("1 0 R"), actualBytes)
}

func TestNewReference_validatesObjectNumber(t *testing.T) {
	reference, err := objects.NewReference(0, 0)
	assert.Error(t, err)
	assert.Equal(t, "object number", err.Context)
	assert.Equal(t, model.PositiveIntegerMessage, err.Message)
	assert.Equal(t, "0", err.Value)
	assert.Nil(t, reference)
}

func TestNewReference_validatesGenerationNumber(t *testing.T) {
	reference, err := objects.NewReference(1, -1)
	assert.Error(t, err)
	assert.Equal(t, "generation number", err.Context)
	assert.Equal(t, model.NonNegativeIntegerMessage, err.Message)
	assert.Equal(t, "-1", err.Value)
	assert.Nil(t, reference)
}

func TestNull_AsASCIIBytes(t *testing.T) {
	actual := objects.Null{}
	assert.Nil(t, actual.Reference)
	assert.Nil(t, actual.Label())

	actualBytes, err := actual.AsASCIIBytes()
	assert.Nil(t, err)
	assert.Equal(t, objects.NullBytes, actualBytes)
}

func TestIndirectNull(t *testing.T) {
	reference, err := objects.NewReference(1,0)
	assert.Nil(t, err)
	actual := objects.Null{Reference: reference}
	assert.NotNil(t, actual.Label())
	assert.Equal(t, 1, actual.ObjectNumber)
	assert.Equal(t, 0, actual.GenerationNumber)
}
