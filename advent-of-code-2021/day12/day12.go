package day12

import (
	"fmt"
	"strings"

	"github.com/pietrodll/aoc2021/utils/collections"
)

func parseInput(input string) map[string][]string {
	graph := make(map[string][]string)

	for _, pathStr := range strings.Split(input, "\n") {
		path := strings.Split(pathStr, "-")
		src, dst := path[0], path[1]

		graph[src] = append(graph[src], dst)
		graph[dst] = append(graph[dst], src)
	}

	return graph
}

type VisitTracker interface {
	canVisit(string) bool
	register(string)
	unregister(string)
}

func findAllPathsRec(
	g *map[string][]string,
	curr string,
	destination string,
	currentPath *[]string,
	pathContainer *collections.StringSet,
	tracker VisitTracker,
) {
	if curr == destination {
		pathContainer.Add(strings.Join(*currentPath, ",") + "," + destination)
		return
	}

	tracker.register(curr)

	*currentPath = append(*currentPath, curr)

	for _, neighbor := range (*g)[curr] {
		if tracker.canVisit(neighbor) {
			findAllPathsRec(
				g,
				neighbor,
				destination,
				currentPath,
				pathContainer,
				tracker,
			)
		}
	}

	*currentPath = (*currentPath)[:len(*currentPath)-1]

	tracker.unregister(curr)
}

type AtMostOnceVisitTracker struct {
	visited collections.StringSet
}

func (v *AtMostOnceVisitTracker) canVisit(cave string) bool {
	return strings.ToUpper(cave) == cave || !v.visited.Contains(cave)
}

func (v *AtMostOnceVisitTracker) register(cave string) {
	if strings.ToLower(cave) == cave {
		v.visited.Add(cave)
	}
}

func (v *AtMostOnceVisitTracker) unregister(cave string) {
	if v.visited.Contains(cave) {
		v.visited.Remove(cave)
	}
}

func findAllPathsWithoutVisitingSmallCavesTwice(g map[string][]string) []string {
	currentPath := make([]string, 0)
	pathContainer := collections.NewStringSet()
	tracker := AtMostOnceVisitTracker{collections.NewStringSet()}

	findAllPathsRec(&g, "start", "end", &currentPath, &pathContainer, &tracker)

	return pathContainer.ToArray()
}

func countPathsWithoutVisitingSmallCavesTwice(g map[string][]string) int {
	return len(findAllPathsWithoutVisitingSmallCavesTwice(g))
}

type AtMostTwiceVisitTracker struct {
	smallCave        string
	smallCaveVisited bool
	visited          collections.StringSet
}

func (v *AtMostTwiceVisitTracker) canVisit(cave string) bool {
	return !v.visited.Contains(cave)
}

func (v *AtMostTwiceVisitTracker) register(cave string) {
	if strings.ToUpper(cave) == cave {
		return
	}

	if cave == v.smallCave && !v.smallCaveVisited {
		v.smallCaveVisited = true
	} else {
		v.visited.Add(cave)
	}
}

func (v *AtMostTwiceVisitTracker) unregister(cave string) {
	if v.visited.Contains(cave) {
		v.visited.Remove(cave)
	} else if v.smallCave == cave {
		v.smallCaveVisited = false
	}
}

func findAllPathsVisitOneSmallCaveTwice(g map[string][]string) []string {
	pathContainer := collections.NewStringSet()

	for cave := range g {
		if cave != "start" && cave != "end" && strings.ToLower(cave) == cave {
			currentPath := make([]string, 0)
			tracker := AtMostTwiceVisitTracker{cave, false, collections.NewStringSet()}
			findAllPathsRec(&g, "start", "end", &currentPath, &pathContainer, &tracker)
		}
	}

	return pathContainer.ToArray()
}

func countPathsVisitOneSmallCaveTwice(g map[string][]string) int {
	return len(findAllPathsVisitOneSmallCaveTwice(g))
}

func Run(input string) {
	graph := parseInput(input)
	fmt.Println(
		"Number of paths without visiting twice small caves:",
		countPathsWithoutVisitingSmallCavesTwice(graph),
	)
	fmt.Println(
		"Number of paths by visiting a single small cave twice:",
		countPathsVisitOneSmallCaveTwice(graph),
	)
}
