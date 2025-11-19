package task_two

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	num := 10
	plusTen(&num)
	assert.Equal(t, 20, num)

	plusPointerValue(&num, 5)
	assert.Equal(t, 25, num)

	numArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	multiply(&numArr, 2)
	assert.Equal(t, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, numArr)

}
