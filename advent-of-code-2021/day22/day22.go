package day22

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pietrodll/aoc2021/utils/parse"
)

type Cuboid struct {
	MinX, MaxX, MinY, MaxY, MinZ, MaxZ int
}

type Instruction struct {
	Cuboid
	On bool
}

func parseInput(input string) []Instruction {
	pattern := regexp.MustCompile(
		`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`,
	)

	lines := strings.Split(input, "\n")
	cuboids := make([]Instruction, len(lines))

	for i, line := range lines {
		match := pattern.FindStringSubmatch(line)
		if match == nil {
			panic(fmt.Errorf("cannot match \"%s\"", line))
		}

		values := parse.StringsToIntegers(match[2:])

		cuboids[i] = Instruction{
			Cuboid{
				values[0],
				values[1],
				values[2],
				values[3],
				values[4],
				values[5],
			},
			match[1] == "on",
		}
	}

	return cuboids
}

func (i Instruction) InRegion(r Region) bool {
	return r.Contains(i.Cuboid)
}

type Region interface {
	Execute(i Instruction)
	Contains(c Cuboid) bool
	CountOn() int
}

type FiniteRegion struct {
	MinX, MaxX, MinY, MaxY, MinZ, MaxZ int
	grid                               [][][]bool
}

func InitFiniteRegion(minX, maxX, minY, maxY, minZ, maxZ int) FiniteRegion {
	grid := make([][][]bool, maxX-minX+1)

	for i := range grid {
		grid[i] = make([][]bool, maxY-minY+1)

		for j := range grid[i] {
			grid[i][j] = make([]bool, maxZ-minZ+1)
		}
	}

	return FiniteRegion{minX, maxX, minY, maxY, minZ, maxZ, grid}
}

func (r *FiniteRegion) Contains(c Cuboid) bool {
	return (c.MaxX <= r.MaxX &&
		c.MinX >= r.MinX &&
		c.MaxY <= r.MaxY &&
		c.MinY >= r.MinY &&
		c.MaxZ <= r.MaxZ &&
		c.MinZ >= r.MinZ)
}

func (r *FiniteRegion) setCube(x, y, z int, on bool) {
	r.grid[x-r.MinX][y-r.MinY][z-r.MinZ] = on
}

func (r *FiniteRegion) Execute(i Instruction) {
	for x := i.MinX; x <= i.MaxX; x++ {
		for y := i.MinY; y <= i.MaxY; y++ {
			for z := i.MinZ; z <= i.MaxZ; z++ {
				r.setCube(x, y, z, i.On)
			}
		}
	}
}

func (r *FiniteRegion) CountOn() int {
	count := 0

	for _, plane := range r.grid {
		for _, line := range plane {
			for _, val := range line {
				if val {
					count++
				}
			}
		}
	}

	return count
}

func (c *Cuboid) Includes(other Cuboid) bool {
	return (other.MaxX <= c.MaxX &&
		other.MinX >= c.MinX &&
		other.MaxY <= c.MaxY &&
		other.MinY >= c.MinY &&
		other.MaxZ <= c.MaxZ &&
		other.MinZ >= c.MinZ)
}

func disjoint(c1, c2 Cuboid) bool {
	return (c1.MaxX < c2.MinX ||
		c1.MaxY < c2.MinY ||
		c1.MaxZ < c2.MinZ) ||
		(c2.MaxX < c1.MinX ||
			c2.MaxY < c1.MinY ||
			c2.MaxZ < c1.MinZ)
}

func intersection(c1, c2 Cuboid) {}

func rebootReactor(r Region, instructions []Instruction) int {
	for _, instr := range instructions {
		if instr.InRegion(r) {
			r.Execute(instr)
		}
	}

	return r.CountOn()
}

func Run(input string) {
	instructions := parseInput(input)

	finite := InitFiniteRegion(-50, 50, -50, 50, -50, 50)
	fmt.Println(rebootReactor(&finite, instructions))
}
