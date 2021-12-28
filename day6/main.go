package main

import (
	"fmt"
	"io/ioutil"
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

	fishState := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, i := range strings.Split(input, ",") {
		fishState[valueFrom(i)]++
	}
	fmt.Println("Fish: ", simulate2(fishState, 256))
	fmt.Println("Fin.")
}

func simulate2(state []int, days int) int {

	for day := days; day > 0; day-- {
		nextState := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

		for i := 8; i > 0; i-- {
			nextState[i-1] = state[i]
		}
		nextState[8] = state[0]
		nextState[6] += state[0]
		state = nextState
	}
	sum := 0
	for _, k := range state {
		sum += k
	}
	return sum
}

/*
type FishContainer struct {
	part1 []uint8
	part2 []uint8
	part3 []uint8
	part4 []uint8
}

func (fc *FishContainer) Get(index uint64) uint8 {

	switch index / math.MaxInt32 {
	case 0:
		return fc.part1[index%math.MaxInt32]
	case 1:
		return fc.part2[index%math.MaxInt32]
	case 2:
		return fc.part3[index%math.MaxInt32]
	case 3:
		return fc.part4[index%math.MaxInt32]
	}

	return 0
}

func (fc *FishContainer) Set(index uint64, value uint8) {
	switch index / math.MaxInt32 {
	case 0:
		fc.part1[index%math.MaxInt32] = value
	case 1:
		fc.part2[index%math.MaxInt32] = value
	case 2:
		fc.part3[index%math.MaxInt32] = value
	case 3:
		fc.part4[index%math.MaxInt32] = value
	}

}

func MakeContainer() *FishContainer {
	return &FishContainer{
		part1: make([]uint8, math.MaxInt32),
		part2: make([]uint8, math.MaxInt32),
		part3: make([]uint8, math.MaxInt32),
		part4: make([]uint8, math.MaxInt32),
	}
}

func simulate(input []uint8, days int) uint64 {

	fish := MakeContainer()
	var currentFish uint64
	currentFish = uint64(len(input))

	for i, fi := range input {
		fish.Set(uint64(i), fi)
	}

	for day := days; day > 0; day-- {
		//	fmt.Printf("Processing day %d\n", day)
		// simulate day x
		spawnCount := 0
		var i uint64
		for i = 0; i < currentFish; i++ {
			if fish.Get(i) == 0 {
				fish.Set(currentFish+uint64(spawnCount), 8)
				spawnCount++
				fish.Set(i, 6)
			} else {
				fish.Set(i, fish.Get(i)-1)
			}
		}
		currentFish += uint64(spawnCount)
	}
	return currentFish
}
*/
func readInput(file string) string {
	if data, err := ioutil.ReadFile(file); err == nil {
		return string(data)

	}
	return ""
}
