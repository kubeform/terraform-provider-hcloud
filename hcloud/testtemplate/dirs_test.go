package testtemplate_test

import (
	"path/filepath"
	"testing"

	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/testsupport"
	"github.com/hetznercloud/terraform-provider-hcloud/hcloud/testtemplate"
	"github.com/stretchr/testify/assert"
)

func TestResourceTemplateDir(t *testing.T) {
	expected := filepath.Join(testsupport.ProjectRoot(t), "internal", "testdata", "r")
	actual := testtemplate.ResourceTemplateDir(t)
	assert.Equal(t, expected, actual)
}

func TestDataSourceTemplateDir(t *testing.T) {
	expected := filepath.Join(testsupport.ProjectRoot(t), "internal", "testdata", "d")
	actual := testtemplate.DataSourceTemplateDir(t)
	assert.Equal(t, expected, actual)
}
