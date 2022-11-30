package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2021 Advent of Code Day 11")
	fmt.Println("--- Part 1 ---")

	input := readInput("input.txt")

	lines := strings.Split(input, "\n")

	width := len(lines[0])
	height := len(lines)

	grid := &Grid{
		TotalFlashes: 0,
		Width:        width,
		Height:       height,
		Positions:    make([]*Octopus, width*height),
	}

	parseGrid(lines, grid)

	allFleshes := false

	for step := 0; !allFleshes; step++ {

		for _, o := range grid.Positions {
			increaseEnergy(o, grid)
		}

		if allFlashes(grid) {
			fmt.Println("All flash after step:  ", step+1)
			allFleshes = true
		}

		resetFlashes(grid)
	}
	fmt.Printf("Total Flashes: %d\n", grid.TotalFlashes)

	fmt.Println("Fin.")
}

func allFlashes(grid *Grid) bool {

	for _, o := range grid.Positions {
		if o.hasFlashed == false {
			return false
		}
	}

	return true
}

func increaseEnergy(o *Octopus, grid *Grid) {
	o.energy++
	if o.energy > 9 {
		o.flash(grid)
	}
}

func resetFlashes(grid *Grid) {
	for _, oct := range grid.Positions {
		if oct.hasFlashed {
			oct.hasFlashed = false
			oct.energy = 0
		}

	}
}

func parseGrid(lines []string, grid *Grid) {
	y := 0
	for _, line := range lines {
		x := 0
		for _, ch := range line {
			grid.Set(x, y, &Octopus{
				x:      x,
				y:      y,
				energy: valueFrom(fmt.Sprintf("%c", ch)),
			})
			x++
		}
		y++
	}
}

type Octopus struct {
	x, y       int
	energy     int
	hasFlashed bool
}

func (o *Octopus) flash(grid *Grid) {

	if o.hasFlashed {
		return
	}
	o.hasFlashed = true
	grid.TotalFlashes++
	// increase all adjacent octopuses
	for y := o.y - 1; y < o.y+2; y++ {
		for x := o.x - 1; x < o.x+2; x++ {
			// skip mid
			if !(x == o.x && y == o.y) {
				if oct := grid.Get(x, y); oct != nil {
					increaseEnergy(oct, grid)
				}
			}
		}
	}

}

type Grid struct {
	TotalFlashes int
	Width        int
	Height       int
	Positions    []*Octopus
}

func (g *Grid) Get(x, y int) *Octopus {

	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return nil
	}

	return g.Positions[y*g.Width+x]
}

func (g *Grid) Set(x, y int, value *Octopus) {
	g.Positions[y*g.Width+x] = value
}

func valueFrom(num string) int {
	v, _ := strconv.Atoi(num)
	return v
}

func readInput(file string) string {
	if data, err := ioutil.ReadFile(file); err == nil {
		return string(data)

	}
	return ""
}
