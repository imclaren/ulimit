package ulimit

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestSetMax(t *testing.T) {
	err := SetMax(0)
	assert.Nil(t, err)
}
