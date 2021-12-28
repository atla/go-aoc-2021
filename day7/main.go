package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func valueFrom(num string) int {
	v, _ := strconv.Atoi(num)
	return v
}

func main() {
	fmt.Println("2021 Advent of Code Day 6")
	fmt.Println("--- Part 1 ---")

	input := readInput("input.txt")

	var crabs []int
	for _, i := range strings.Split(input, ",") {
		crabs = append(crabs, valueFrom(i))
	}

	minimum := math.MaxInt32

	minPos := findMinPos(crabs)
	maxPos := findMaxPos(crabs)

	for pos := minPos; pos <= maxPos; pos++ {
		fuelForPos := calculateFuelForPosition(crabs, pos, func(target, crab int) int {
			return Abs(target - crab)
		})
		minimum = Min(minimum, fuelForPos)
	}

	fmt.Println("Minimum: ", minimum)

	minimum2 := math.MaxInt32

	for pos := minPos; pos <= maxPos; pos++ {
		fuelForPos := calculateFuelForPosition(crabs, pos, func(target, crab int) int {
			cost := 0
			for i := 0; i < Abs(target-crab); i++ {
				cost += 1 + i
			}
			return cost
		})
		minimum2 = Min(minimum2, fuelForPos)
	}
	fmt.Println("Minimum2: ", minimum2)
	fmt.Println("Fin.")
}

func findMaxPos(crabs []int) int {
	max := math.MinInt32
	for _, c := range crabs {
		max = Max(max, c)
	}
	return max
}

func findMinPos(crabs []int) int {
	min := math.MaxInt32
	for _, c := range crabs {
		min = Min(min, c)
	}
	return min
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
func calculateFuelForPosition(crabs []int, position int, costfunc func(int, int) int) int {
	cost := 0
	for _, c := range crabs {
		cost += costfunc(position, c)
	}
	return cost
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
