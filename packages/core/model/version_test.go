package model_test

import (
	"github.com/stretchr/testify/assert"
	"pdfeg-core/model"
	"testing"
)

func Test_Version_String(t *testing.T) {
	version := model.Version{Major: 1, Minor: 7}
	assert.Equal(t, "1.7", version.String())
}
