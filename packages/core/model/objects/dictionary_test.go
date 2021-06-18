package objects_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"pdfeg-core/model/objects"
	"strings"
	"testing"
)

var entryKey0 = "K0"
var entryValue0 = objects.Integer{Value: 1}
var entryKey1 = "K1"
var entryValue1 = objects.Real{Value: 2.3}
var entryKey2 = "K2"
var entryValue2 = objects.Name{Value: "Name"}
var entries = map[string]objects.Object{
	entryKey0: entryValue0,
	entryKey1: entryValue1,
	entryKey2: entryValue2,
}

func TestDictionary_Label_Direct(t *testing.T) {
	dict := objects.Dictionary{Entries: entries}
	assert.Nil(t, dict.Label())
}

func TestDictionary_Label_Indirect(t *testing.T) {
	reference, err := objects.NewReference(1, 0)
	assert.Nil(t, err)
	dict := objects.Dictionary{Reference: reference, Entries: entries}
	assert.Equal(t, len(entries), len(dict.Entries))
	assert.NotNil(t, dict.Label())
	assert.Equal(t, 1, dict.ObjectNumber)
	assert.Equal(t, 0, dict.GenerationNumber)
}

func TestDictionary_AsASCIIBytes(t *testing.T) {
	dict := objects.Dictionary{Entries: entries}
	bytes, err := dict.AsASCIIBytes()
	assert.Nil(t, err)
	bytesAsString := string(bytes)
	assert.True(t, strings.HasPrefix(bytesAsString, "<<"))
	assert.True(t, strings.HasSuffix(bytesAsString, ">>"))
	assert.Contains(t, bytesAsString, "/K2 /Name")
	assert.Contains(t, bytesAsString, "/K0 1 /K1 2.3")
	assert.Contains(t, bytesAsString, "/K1 2.3")
}

func TestDictionary_AsASCIIBytes_ErrorInValue(t *testing.T) {
	expectedError := errors.New("expected")
	entryValueError := objMock{asASCIIBytesMock: func() ([]byte, error) {
		return nil, expectedError
	}}
	array := objects.Dictionary{Entries: map[string]objects.Object{entryKey0: entryValueError}}
	bytes, err := array.AsASCIIBytes()
	assert.Error(t, err)
	assert.Nil(t, bytes)
}
