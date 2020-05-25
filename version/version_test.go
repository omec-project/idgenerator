package version_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/free5gc/idgenerator/version"
)

func TestVersion(t *testing.T) {
	assert.Equal(t, "2020-05-25-01", version.GetVersion())
}
