package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type pos struct {
	x   int
	y   int
	aim int
}

func main() {
	fmt.Println("2021 Advent of Code Day 2")
	fmt.Println("--- Part 1 ---")

	input := readInput("input.txt")
	start := pos{0, 0, 0}

	if result, err := followCommands(input, start); err == nil {
		fmt.Printf("Position %d \n", result)
		fmt.Printf("Result %d \n", result.x*result.y)
	}

	fmt.Println("--- Part 2 ---")

	start = pos{0, 0, 0}

	if result, err := followCommandsV2(input, start); err == nil {
		fmt.Printf("Position %d \n", result)
		fmt.Printf("Result %d \n", result.x*result.y)
	}
}

func followCommands(input []string, position pos) (pos, error) {

	if len(input) == 0 {
		return position, nil
	}

	split := strings.Split(input[0], " ")
	distance, _ := strconv.Atoi(split[1])

	switch split[0] {
	case "forward":
		position.x += distance
	case "down":
		position.y += distance
	case "up":
		position.y -= distance
	}

	return followCommands(input[1:], position)

}

func followCommandsV2(input []string, position pos) (pos, error) {

	if len(input) == 0 {
		return position, nil
	}

	split := strings.Split(input[0], " ")
	distance, _ := strconv.Atoi(split[1])

	switch split[0] {
	case "forward":
		position.x += distance
		position.y += position.aim * distance
	case "down":
		position.aim += distance
	case "up":
		position.aim -= distance
	}

	return followCommandsV2(input[1:], position)

}

func readInput(file string) []string {

	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n")
	}
	return nil
}
