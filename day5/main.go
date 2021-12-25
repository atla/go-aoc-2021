package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type pos struct {
	x, y  int
	value int
}

type line struct {
	x1, x2 int
	y1, y2 int
}

type grid struct {
	columns   int
	rows      int
	positions []*pos
}

func valueFrom(num string) int {
	v, _ := strconv.Atoi(num)
	return v
}

func main() {
	fmt.Println("2021 Advent of Code Day 5")
	fmt.Println("--- Part 1 ---")

	lines := []line{}
	width, height := 0, 0

	input := readInput("input.txt")
	input = strings.Replace(input, " ", "", -1)
	linesInput := strings.Split(input, "\n")
	for _, li := range linesInput {
		pairs := strings.Split(li, "->")
		start := strings.Split(pairs[0], ",")
		end := strings.Split(pairs[1], ",")

		line := line{
			x1: valueFrom(start[0]),
			y1: valueFrom(start[1]),
			x2: valueFrom(end[0]),
			y2: valueFrom(end[1]),
		}
		lines = append(lines, line)

		width = Max(line.x2, Max(line.x1, width))
		height = Max(line.y2, Max(line.y1, height))
	}

	field := InitField(width+1, height+1)

	for _, l := range lines {
		field = drawLine(field, l, width+1, false)
	}

	//Print(field, width+1, height+1)

	count := 0

	for _, i := range field {
		if i > 1 {
			count = count + 1
		}
	}

	fmt.Printf("Count is %d\n", count)

	fmt.Println("--- Part 2 ---")

	field2 := InitField(width+1, height+1)

	for _, l := range lines {
		field2 = drawLine(field2, l, width+1, true)
	}

	count2 := 0

	for _, i := range field2 {
		if i > 1 {
			count2 = count2 + 1
		}
	}
	fmt.Printf("Count 2 is %d\n", count2)

	//Print(field2, width+1, height+1)

	fmt.Println("Fin.")

}

func drawLine(i []int, l line, width int, v2 bool) []int {

	// check for diagonal
	if v2 && isDiagonal(l) {

		var xd, yd int
		if l.x2 > l.x1 {
			xd = 1
		} else {
			xd = -1
		}
		if l.y2 > l.y1 {
			yd = 1
		} else {
			yd = -1
		}

		run := true
		x := l.x1
		y := l.y1
		for run {
			i[y*width+x] = i[y*width+x] + 1

			run = !(x == l.x2)
			x += xd
			y += yd
		}

	} else if l.x1 == l.x2 {
		if l.y1 < l.y2 {
			for y := l.y1; y <= l.y2; y++ {
				i[y*width+l.x1] = i[y*width+l.x1] + 1
			}
		} else {
			for y := l.y2; y <= l.y1; y++ {
				i[y*width+l.x1] = i[y*width+l.x1] + 1
			}
		}

	} else if l.y1 == l.y2 {
		if l.x1 < l.x2 {
			for x := l.x1; x <= l.x2; x++ {
				i[l.y1*width+x] = i[l.y1*width+x] + 1
			}
		} else {
			for x := l.x2; x <= l.x1; x++ {
				i[l.y1*width+x] = i[l.y1*width+x] + 1
			}
		}

	}
	return i
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func isDiagonal(l line) bool {
	return Abs(l.x1-l.x2) == Abs(l.y1-l.y2)
}

func InitField(width int, height int) []int {
	field := make([]int, width*height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			field[y*width+x] = 0
		}
	}
	return field
}

func Print(field []int, width int, height int) {
	fmt.Println("Field: ")
	for y := 0; y < height; y++ {
		for x := 0; x < height; x++ {
			if field[y*width+x] == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%d ", field[y*width+x])
			}

		}
		fmt.Println("")
	}

}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (b grid) valueAt(col, row int) pos {
	return *b.positions[(row*b.columns)+col]
}

func readInput(file string) string {
	if data, err := ioutil.ReadFile(file); err == nil {
		return string(data)

	}
	return ""
}
