package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2021 Advent of Code Day 1")
	fmt.Println("--- Part 1 ---")

	input := readInput("input.txt")

	if result, err := findDepthMeasurementIncreases(input); err == nil {
		fmt.Printf("Larger than previous %d \n", result)

	}

	if result, err := findSlidingWindowIncreases(input); err == nil {
		fmt.Printf("SlidingWindow: Larger than previous %d \n", result)

	}
}

func readInput(file string) []int {
	var numbers []int

	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		splitted := strings.Split(input, "\n")

		for _, number := range splitted {
			n, _ := strconv.Atoi(number)
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func findSlidingWindowIncreases(input []int) (int, error) {

	if len(input) <= 1 {
		return -1, errors.New("input len <= 0")
	}

	increases := 0
	i := 1

	for i < len(input)-2 {
		left := input[i-1] + input[i] + input[i+1]
		right := input[i] + input[i+1] + input[i+2]

		if right > left {
			increases++
		}
		i++
	}
	return increases, nil

}

func findDepthMeasurementIncreases(input []int) (int, error) {

	if len(input) <= 1 {
		return -1, errors.New("input len <= 0")
	}

	increases := 0
	i := 1

	for i < len(input) {
		if input[i] >= input[i-1] {
			increases++
		}
		i++
	}
	return increases, nil
}
