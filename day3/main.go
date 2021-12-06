package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type pos struct {
	x   int
	y   int
	aim int
}

func main() {
	fmt.Println("2021 Advent of Code Day 3")
	fmt.Println("--- Part 1 ---")

	input := readInput("test.txt")
	start := pos{0, 0, 0}

	if result, err := findCommonBits(input, start); err == nil {
		fmt.Printf("Result %s \n", result)

	}
}

func findCommonBits(input []string, position pos) (string, error) {

	result := ""
	length := len(input[0])

	for i := 0; i < length; i++ {

		countSum := 0

		for _, s := range input {
			if s[i] == '1' {
				countSum++
			
		}

		if float32(countSum)*1.0 > float32(length)*0.5 {
			result += "1"
		} else {
			result += "0"
		}
	}

	return result, nil
}

func readInput(file string) []string {

	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n")
	}
	return nil
}
