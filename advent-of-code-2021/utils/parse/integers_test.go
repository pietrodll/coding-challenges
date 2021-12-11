package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIntegers(t *testing.T) {
	data := "1,2,3,4,5,6"
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, ParseIntegers(data, ","))

	assert.Panics(t, func() {
		ParseIntegers("1,2,3,x,5,6", ",")
	})
}
