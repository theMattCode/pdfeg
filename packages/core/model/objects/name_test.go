package objects_test

import (
	"github.com/stretchr/testify/assert"
	"pdfeg-core/model/objects"
	"testing"
)

func TestName_AsASCIIBytes(t *testing.T) {
	name := objects.Name{Value: "Name"}
	actualBytes, err := name.AsASCIIBytes()
	assert.Nil(t, err)
	assert.Equal(t, []byte("/Name"), actualBytes)
}

func TestName_Label(t *testing.T) {
	name := objects.Name{Value: "Name"}
	assert.Nil(t, name.Reference)
	assert.Nil(t, name.Label())
}

func TestIndirectName(t *testing.T) {
	reference, err := objects.NewReference(1,0)
	assert.Nil(t, err)
	name := objects.Name{Reference: reference, Value: "Name"}
	assert.Equal(t, 1, name.ObjectNumber)
	assert.Equal(t, 0, name.GenerationNumber)
}
