package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

func TestParseInput(t *testing.T) {
	expected := Entry{[]string{"be", "abcdefg", "bcdefg", "acdefg", "bceg", "cdefg", "abdefg", "bcdef", "abcdf", "bde"}, []string{"abcdefg", "bcdef", "bcdefg", "bceg"}}

	assert.Equal(t, expected, parseInput(input)[0])
}

func TestCountUniqueLengthDigits(t *testing.T) {
	entries := parseInput(input)

	assert.Equal(t, 26, countUniqueSegmentDigits(entries))
}

func TestDecodeDigits(t *testing.T) {
	expected := []int{8394, 9781, 1197, 9361, 4873, 8418, 4548, 1625, 8717, 4315}

	for i, entry := range parseInput(input) {
		assert.Equal(t, expected[i], decodeDigits(entry))
	}
}
