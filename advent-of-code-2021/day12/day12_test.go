package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertGraphsEqual(t *testing.T, expected map[string][]string, actual map[string][]string) {
	assert.Equal(t, len(expected), len(actual))

	for source, destinations := range expected {
		actualDestinations, present := actual[source]
		assert.Truef(t, present, "Key %s is expected to be present", source)
		assert.ElementsMatch(t, destinations, actualDestinations)
	}
}

func TestParseInput(t *testing.T) {
	graph := parseInput("start-A\nstart-b\nA-c\nA-b\nb-d\nA-end\nb-end")
	expected := map[string][]string{
		"start": {"A", "b"},
		"A":     {"start", "b", "c", "end"},
		"c":     {"A"},
		"b":     {"start", "A", "d", "end"},
		"d":     {"b"},
		"end":   {"A", "b"},
	}
	assertGraphsEqual(t, expected, graph)
}

func TestFindAllPathsWithoutVisitingSmallCavesTwice(t *testing.T) {
	graph := parseInput("start-A\nstart-b\nA-c\nA-b\nb-d\nA-end\nb-end")
	expected := []string{
		"start,A,b,A,c,A,end",
		"start,A,b,A,end",
		"start,A,b,end",
		"start,A,c,A,b,A,end",
		"start,A,c,A,b,end",
		"start,A,c,A,end",
		"start,A,end",
		"start,b,A,c,A,end",
		"start,b,A,end",
		"start,b,end",
	}
	assert.ElementsMatch(t, expected, findAllPathsWithoutVisitingSmallCavesTwice(graph))

	graph = parseInput("dc-end\nHN-start\nstart-kj\ndc-start\ndc-HN\nLN-dc\nHN-end\nkj-sa\nkj-HN\nkj-dc")
	expected = []string{
		"start,HN,dc,HN,end",
		"start,HN,dc,HN,kj,HN,end",
		"start,HN,dc,end",
		"start,HN,dc,kj,HN,end",
		"start,HN,end",
		"start,HN,kj,HN,dc,HN,end",
		"start,HN,kj,HN,dc,end",
		"start,HN,kj,HN,end",
		"start,HN,kj,dc,HN,end",
		"start,HN,kj,dc,end",
		"start,dc,HN,end",
		"start,dc,HN,kj,HN,end",
		"start,dc,end",
		"start,dc,kj,HN,end",
		"start,kj,HN,dc,HN,end",
		"start,kj,HN,dc,end",
		"start,kj,HN,end",
		"start,kj,dc,HN,end",
		"start,kj,dc,end",
	}
	assert.ElementsMatch(t, expected, findAllPathsWithoutVisitingSmallCavesTwice(graph))
}

func TestCountPathsWithoutVisitingSmallCavesTwice(t *testing.T) {
	graph := parseInput("fs-end\nhe-DX\nfs-he\nstart-DX\npj-DX\nend-zg\nzg-sl\nzg-pj\npj-he\nRW-he\nfs-DX\npj-RW\nzg-RW\nstart-pj\nhe-WI\nzg-he\npj-fs\nstart-RW")
	assert.Equal(t, 226, countPathsWithoutVisitingSmallCavesTwice(graph))

}

func TestFindAllPathsVisitOneSmallCaveTwice(t *testing.T) {
	graph := parseInput("start-A\nstart-b\nA-c\nA-b\nb-d\nA-end\nb-end")
	expected := []string{
		"start,A,b,A,b,A,c,A,end",
		"start,A,b,A,b,A,end",
		"start,A,b,A,b,end",
		"start,A,b,A,c,A,b,A,end",
		"start,A,b,A,c,A,b,end",
		"start,A,b,A,c,A,c,A,end",
		"start,A,b,A,c,A,end",
		"start,A,b,A,end",
		"start,A,b,d,b,A,c,A,end",
		"start,A,b,d,b,A,end",
		"start,A,b,d,b,end",
		"start,A,b,end",
		"start,A,c,A,b,A,b,A,end",
		"start,A,c,A,b,A,b,end",
		"start,A,c,A,b,A,c,A,end",
		"start,A,c,A,b,A,end",
		"start,A,c,A,b,d,b,A,end",
		"start,A,c,A,b,d,b,end",
		"start,A,c,A,b,end",
		"start,A,c,A,c,A,b,A,end",
		"start,A,c,A,c,A,b,end",
		"start,A,c,A,c,A,end",
		"start,A,c,A,end",
		"start,A,end",
		"start,b,A,b,A,c,A,end",
		"start,b,A,b,A,end",
		"start,b,A,b,end",
		"start,b,A,c,A,b,A,end",
		"start,b,A,c,A,b,end",
		"start,b,A,c,A,c,A,end",
		"start,b,A,c,A,end",
		"start,b,A,end",
		"start,b,d,b,A,c,A,end",
		"start,b,d,b,A,end",
		"start,b,d,b,end",
		"start,b,end",
	}
	assert.ElementsMatch(t, expected, findAllPathsVisitOneSmallCaveTwice(graph))
}

func TestCountPathsVisitOneSmallCaveTwice(t *testing.T) {
	graph := parseInput("dc-end\nHN-start\nstart-kj\ndc-start\ndc-HN\nLN-dc\nHN-end\nkj-sa\nkj-HN\nkj-dc")
	assert.Equal(t, 103, countPathsVisitOneSmallCaveTwice(graph))

	graph = parseInput("fs-end\nhe-DX\nfs-he\nstart-DX\npj-DX\nend-zg\nzg-sl\nzg-pj\npj-he\nRW-he\nfs-DX\npj-RW\nzg-RW\nstart-pj\nhe-WI\nzg-he\npj-fs\nstart-RW")
	assert.Equal(t, 3509, countPathsVisitOneSmallCaveTwice(graph))

}
