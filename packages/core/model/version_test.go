package model_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"pdfeg-core/model"
	"testing"
)

func Test_Version_String(t *testing.T) {
	version := model.Version{Major: 1, Minor: 7}
	assert.Equal(t, "1.7", version.String())
}

func Test_Version_AsFileHeaderToken(t *testing.T) {
	supportedVersions := []struct {
		version  model.Version
		expected string
	}{
		{model.Version{Major: 1, Minor: 0}, "%PDF-1.0"},
		{model.Version{Major: 1, Minor: 1}, "%PDF-1.1"},
		{model.Version{Major: 1, Minor: 2}, "%PDF-1.2"},
		{model.Version{Major: 1, Minor: 3}, "%PDF-1.3"},
		{model.Version{Major: 1, Minor: 4}, "%PDF-1.4"},
		{model.Version{Major: 1, Minor: 5}, "%PDF-1.5"},
		{model.Version{Major: 1, Minor: 6}, "%PDF-1.6"},
		{model.Version{Major: 1, Minor: 7}, "%PDF-1.7"},
		{model.Version{Major: 2, Minor: 0}, "%PDF-2.0"},
	}
	for _, testcase := range supportedVersions {
		t.Run(fmt.Sprintf("Testing Token for %s", testcase.version.String()),
			func(t *testing.T) {
				assert.Equal(t, []byte(testcase.expected), testcase.version.AsFileHeaderToken())
			},
		)
	}
}
