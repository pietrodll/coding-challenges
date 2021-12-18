package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	toCheck := []string{
		"[1,2]",
		"[[1,2],3]",
		"[9,[8,7]]",
		"[[6,[5,[4,[3,2]]]],1]",
	}

	for _, numberStr := range toCheck {
		parsed := parseSnailfishNumber(numberStr)
		assert.Equal(t, numberStr, parsed.String())
	}
}

func TestMagnitude(t *testing.T) {
	magnitude := map[string]int{
		"[[1,2],[[3,4],5]]":                                     143,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]":                     1384,
		"[[[[1,1],[2,2]],[3,3]],[4,4]]":                         445,
		"[[[[3,0],[5,3]],[4,4]],[5,5]]":                         791,
		"[[[[5,0],[7,4]],[5,5]],[6,6]]":                         1137,
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]": 3488,
	}

	for numberStr, expectedMagnitude := range magnitude {
		number := parseSnailfishNumber(numberStr)
		assert.Equal(t, expectedMagnitude, number.Magnitude())
	}
}

func TestExplode(t *testing.T) {
	beforeAfter := map[string]string{
		"[[[[[9,8],1],2],3],4]":                 "[[[[0,9],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]":                 "[7,[6,[5,[7,0]]]]",
		"[[6,[5,[4,[3,2]]]],1]":                 "[[6,[5,[7,0]]],3]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]": "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]":     "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	}

	for beforeStr, afterStr := range beforeAfter {
		before := parseInput(beforeStr)[0]
		hasExploded := explodeStep(before)
		assert.True(t, hasExploded)
		assert.Equal(t, parseSnailfishNumber(afterStr).String(), before.String())
	}
}

func TestSum(t *testing.T) {
	assert.Equal(
		t,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		sum(
			parseSnailfishNumber("[[[[4,3],4],4],[7,[[8,4],9]]]"),
			parseSnailfishNumber("[1,1]"),
		).String(),
	)
}

func TestSumAll(t *testing.T) {
	numbers := parseInput(`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`)

	assert.Equal(
		t,
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		sumAll(numbers).String(),
	)

	numbers = parseInput(`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`)

	assert.Equal(
		t,
		"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
		sumAll(numbers).String(),
	)
}

func TestFindHighestMagnitudeSum(t *testing.T) {
	numbers := []string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	}

	assert.Equal(
		t,
		"[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]",
		findHighestMagnitudeSum(numbers).String(),
	)
}
