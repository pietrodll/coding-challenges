package day19

import (
	"testing"

	"github.com/pietrodll/aoc2021/utils/geometry"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	input := `--- scanner 0 ---
-1,-1,1
-2,-2,2
-3,-3,3
-2,-3,1
5,6,-4
8,0,7`
	expected := [][]*geometry.Vector{
		{
			geometry.MakeVector(-1, -1, 1),
			geometry.MakeVector(-2, -2, 2),
			geometry.MakeVector(-3, -3, 3),
			geometry.MakeVector(-2, -3, 1),
			geometry.MakeVector(5, 6, -4),
			geometry.MakeVector(8, 0, 7),
		},
	}

	assert.Equal(t, expected, parseInput(input))
}

func TestGenerateRotations(t *testing.T) {
	rotations := generateRotations()
	assert.Len(t, rotations, 24)
}

func TestInCommon(t *testing.T) {
	setA := []*geometry.Vector{geometry.MakeVector(1, 2), geometry.MakeVector(3, 3), geometry.MakeVector(4, 2)}
	setB := []*geometry.Vector{geometry.MakeVector(1, 2), geometry.MakeVector(3, 4)}

	assert.Equal(t, []*geometry.Vector{geometry.MakeVector(1, 2)}, findCommonVectors(setA, setB))
}

func TestComputeDifferences(t *testing.T) {
	vectors := []*geometry.Vector{
		geometry.MakeVector(1, 2),
		geometry.MakeVector(2, 3),
		geometry.MakeVector(4, 1),
	}

	expected := []*geometry.Vector{
		geometry.MakeVector(-1, -1),
		geometry.MakeVector(-3, 1),
		geometry.MakeVector(1, 1),
		geometry.MakeVector(-2, 2),
		geometry.MakeVector(3, -1),
		geometry.MakeVector(2, -2),
	}

	assert.ElementsMatch(t, expected, computeDifferences(vectors))
}
