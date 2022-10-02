package day25

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`

func TestFindFirstStoppingStep(t *testing.T) {
	g := parseInput(input)
	assert.Equal(t, 58, findFirstStoppingStep(g))
}
