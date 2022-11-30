package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Basin struct {
	Positions []*Pos
}

func NewBasin() *Basin {
	return &Basin{
		Positions: []*Pos{},
	}
}

func (b *Basin) Exists(p *Pos) bool {
	for _, pos := range b.Positions {

		if pos.x == p.x && pos.y == p.y {
			return true
		}
	}
	return false
}

func (b *Basin) AddPosition(p *Pos) {

	for _, pos := range b.Positions {
		// check if already added
		if pos.x == p.x && pos.y == p.y {
			return
		}
	}

	b.Positions = append(b.Positions, p)
}

func main() {
	fmt.Println("2021 Advent of Code Day 9")
	fmt.Println("--- Part 1 ---")

	input := readInput("input.txt")

	lines := strings.Split(input, "\n")

	width := len(lines[0])
	height := len(lines)

	grid := &Grid{
		Width:     width,
		Height:    height,
		Positions: make([]*Pos, width*height),
	}

	y := 0
	for _, line := range lines {
		x := 0
		for _, ch := range line {
			grid.Set(x, y, &Pos{
				x:     x,
				y:     y,
				value: valueFrom(fmt.Sprintf("%c", ch)),
			})
			x++
		}
		y++
	}

	var lowpoints []*Pos

	for _, pos := range grid.Positions {
		if pos.IsLowPoint(grid) {
			lowpoints = append(lowpoints, pos)
		}
	}
	riskSum := 0
	for _, lp := range lowpoints {
		riskSum += 1 + lp.value
	}
	fmt.Println("RISK SUM", riskSum)

	fmt.Println("Part 2 -- Find Basins")

	basins := make([]*Basin, len(lowpoints))

	for i := 0; i < len(basins); i++ {
		basins[i] = NewBasin()
	}

	for i, lp := range lowpoints {
		b := basins[i]
		floodBasin(grid, b, lp, lp.value)
	}

	basinSizes := []int{}

	for _, b := range basins {
		basinSizes = append(basinSizes, len(b.Positions))
	}
	sort.Ints(basinSizes)
	l := len(basinSizes)

	fmt.Println("Result:", basinSizes[l-1]*basinSizes[l-2]*basinSizes[l-3])

	fmt.Println("Fin.")
}

func floodBasin(grid *Grid, basin *Basin, lp *Pos, lastVal int) {

	if basin.Exists(lp) {
		return
	}
	if lp.value >= 9 {
		return
	}
	if lp.value >= lastVal {
		basin.AddPosition(lp)

		floodBasin(grid, basin, grid.Get(lp.x-1, lp.y), lp.value)
		floodBasin(grid, basin, grid.Get(lp.x+1, lp.y), lp.value)
		floodBasin(grid, basin, grid.Get(lp.x, lp.y+1), lp.value)
		floodBasin(grid, basin, grid.Get(lp.x, lp.y-1), lp.value)
	}
}

type Pos struct {
	x, y  int
	value int
}

func (pos *Pos) IsLowPoint(grid *Grid) bool {

	if pos.value < grid.Get(pos.x-1, pos.y).value && pos.value < grid.Get(pos.x+1, pos.y).value && pos.value < grid.Get(pos.x, pos.y+1).value && pos.value < grid.Get(pos.x, pos.y-1).value {
		return true
	}
	{
		return false
	}
	return false
}

type Grid struct {
	Width     int
	Height    int
	Positions []*Pos
}

func (g *Grid) Get(x, y int) *Pos {

	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return &Pos{
			x:     -1,
			y:     -1,
			value: math.MaxInt32,
		}
	}

	return g.Positions[y*g.Width+x]
}

func (g *Grid) Set(x, y int, value *Pos) {
	g.Positions[y*g.Width+x] = value
}

func valueFrom(num string) int {
	v, _ := strconv.Atoi(num)
	return v
}

func Min(minimum int, pos int) int {
	if minimum < pos {
		return minimum
	}
	return pos
}
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func readInput(file string) string {
	if data, err := ioutil.ReadFile(file); err == nil {
		return string(data)

	}
	return ""
}
