package task_two

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLock(t *testing.T) {
	i := counter(1000)
	assert.Equal(t, 10000, i)

	j := counterAtomic(1000)
	assert.Equal(t, 10000, j)
}
