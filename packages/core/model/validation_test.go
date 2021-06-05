package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ValidationError_Message(t *testing.T) {
	actual := ValidationError{"ctx", "msg", "val"}
	assert.Equal(t, "validation error: ctx: msg: val", actual.Error())
}

func Test_ValidatePositiveInteger_OneIsValid(t *testing.T) {
	assert.Nil(t, ValidatePositiveInteger("positive", 1))
	assert.Nil(t, ValidatePositiveInteger("positive", (1<<32)-1))
}

func Test_ValidatePositiveInteger_ZeroIsInvalid(t *testing.T) {
	context := "zero"
	err := ValidatePositiveInteger(context, 0)
	assert.Error(t, err)
	assert.Equal(t, PositiveIntegerMessage, err.Message)
	assert.Equal(t, "0", err.Value)
	assert.Equal(t, context, err.Context)
}

func Test_ValidatePositiveInteger_NegativeIsInvalid(t *testing.T) {
	context := "negative"
	err := ValidatePositiveInteger(context, -1)
	assert.Error(t, err)
	assert.Equal(t, PositiveIntegerMessage, err.Message)
	assert.Equal(t, "-1", err.Value)
	assert.Equal(t, context, err.Context)
}

func Test_ValidateNonNegativeInteger_ZeroIsValid(t *testing.T) {
	assert.Nil(t, ValidateNonNegativeInteger("zero", 0))
}

func Test_ValidateNonNegativeInteger_NegativeIsInvalid(t *testing.T) {
	context := "negative"
	err := ValidateNonNegativeInteger(context, -1)
	assert.Error(t, err)
	assert.Equal(t, NonNegativeIntegerMessage, err.Message)
	assert.Equal(t, "-1", err.Value)
	assert.Equal(t, context, err.Context)
}
