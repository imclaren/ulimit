package ulimit

import (
	"testing"
	"fmt"

	assert "github.com/stretchr/testify/require"
)

func TestSetMax(t *testing.T) {
	oldLimit, err := Get()
	assert.Nil(t, err)
	err = SetMax()
	assert.Nil(t, err)
	newLimit, err := Get()
	assert.Nil(t, err)
	assert.True(t, newLimit > oldLimit)
	fmt.Println(oldLimit, newLimit)
}
