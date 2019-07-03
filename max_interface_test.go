package max_interface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxInt(t *testing.T) {
	intSlice := []int{1, 2, 5, 100500, 6}

	interfaceSlice := make([]interface{}, len(intSlice))
	for i, v := range intSlice {
		interfaceSlice[i] = v
	}
	max := getMax(
		interfaceSlice,
		func(val1 interface{}, val2 interface{}) bool {
			return val1.(int) > val2.(int)
		},
	)
	assert.Equal(t, 100500, max)
}
