package day19

import (
	"sort"
	"strings"

	"github.com/pietrodll/aoc2021/utils/geometry"
	"github.com/pietrodll/aoc2021/utils/parse"
)

func parseInput(input string) [][]*geometry.Vector {
	blocks := strings.Split(input, "\n\n")
	result := make([][]*geometry.Vector, len(blocks))

	for i, block := range blocks {
		lines := strings.Split(block, "\n")
		points := make([]*geometry.Vector, len(lines)-1)

		for j, coordStr := range lines[1:] {
			points[j] = geometry.MakeVector(parse.ParseIntegers(coordStr, ",")...)
		}

		result[i] = points
	}

	return result
}

func generateRotations() []*geometry.Matrix {
	rotations := make([]*geometry.Matrix, 0, 24)

	cos := []int{1, 0, -1, 0}
	sin := []int{0, 1, 0, -1}

	for yaw := range cos {
		yawMatrix := geometry.MakeMatrix([][]int{
			{cos[yaw], -sin[yaw], 0},
			{sin[yaw], cos[yaw], 0},
			{0, 0, 1},
		})

		for roll := range cos {
			rollMatrix := geometry.MakeMatrix([][]int{
				{1, 0, 0},
				{0, cos[roll], -sin[roll]},
				{0, sin[roll], cos[roll]},
			})

			rotations = append(rotations, yawMatrix.Mult(rollMatrix))
		}
	}

	for pitch := 1; pitch <= 3; pitch += 2 {
		pitchMatrix := geometry.MakeMatrix([][]int{
			{cos[pitch], 0, -sin[pitch]},
			{0, 1, 0},
			{sin[pitch], 0, cos[pitch]},
		})

		for roll := range cos {
			rollMatrix := geometry.MakeMatrix([][]int{
				{1, 0, 0},
				{0, cos[roll], -sin[roll]},
				{0, sin[roll], cos[roll]},
			})

			rotations = append(rotations, pitchMatrix.Mult(rollMatrix))
		}
	}

	return rotations
}

func computeDifferences(vects []*geometry.Vector) []*geometry.Vector {
	differences := make([]*geometry.Vector, 0, len(vects)*(len(vects)-1))

	for i, vec1 := range vects {
		for j, vec2 := range vects {
			if i != j {
				differences = append(differences, vec1.Sub(vec2))
			}
		}
	}

	return differences
}

type VecSlice []*geometry.Vector

func (v VecSlice) Len() int {
	return len(v)
}

func (v VecSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func lessLexOrder(u, v *geometry.Vector) bool {
	for k := 0; k < u.Size(); k++ {
		if u.Get(k) < v.Get(k) {
			return true
		}

		if u.Get(k) > v.Get(k) {
			return false
		}
	}

	return true
}

func (v VecSlice) Less(i, j int) bool {
	return lessLexOrder(v[i], v[j])
}

// the arrays must be sorted in lexicographic order
func findCommonVectors(vectors, otherVectors []*geometry.Vector) []*geometry.Vector {
	inCommon := make([]*geometry.Vector, 0)

	i, j := 0, 0
	for i < len(vectors) && j < len(otherVectors) {
		if vectors[i].Eq(otherVectors[j]) {
			inCommon = append(inCommon, vectors[i])
			i++
			j++
		} else if lessLexOrder(vectors[i], otherVectors[j]) {
			i++
		} else {
			j++
		}
	}

	return inCommon
}

func multiplyAll(matrix *geometry.Matrix, vectors []*geometry.Vector) []*geometry.Vector {
	multiplied := make([]*geometry.Vector, len(vectors))

	for i, vec := range vectors {
		multiplied[i] = matrix.MultVec(vec)
	}

	return multiplied
}

func findScannerPosition(points, otherPoints []*geometry.Vector) *geometry.Vector {
	diff := computeDifferences(points)
	sort.Sort(VecSlice(diff))

	rotations := generateRotations()
	maxInCommon := 0
	var translated, otherDiff []*geometry.Vector

	for _, rot := range rotations {
		rotatedCandidate := multiplyAll(rot, otherPoints)
		otherDiffCandidate := computeDifferences(translated)
		sort.Sort(VecSlice(otherDiff))

		if inCommon := findCommonVectors(diff, otherDiffCandidate); len(inCommon) > maxInCommon {
			maxInCommon = len(inCommon)
			translated = rotatedCandidate
			otherDiff = otherDiffCandidate
		}
	}

	sort.Sort(VecSlice(translated))

	orderedPoints := make([]*geometry.Vector, len(points))
	copy(orderedPoints, points)
	sort.Sort(VecSlice(orderedPoints))

	return nil
}

func Run(input string) {}
