package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type pos struct {
	x, y   int
	marked bool
	value  int
}

type board struct {
	width     int
	height    int
	positions []pos
}

func parseBoard(input string) board {
	rows := len(strings.Split(input, "\n"))
	clean := strings.Split(input, "\n")[0]
	clean = strings.Replace(clean, "  ", " ", -1)
	cols := len(strings.Split(clean, " "))

	b := board{
		width:     cols,
		height:    rows,
		positions: make([]pos, rows*cols),
	}

	parseInput := strings.Split(strings.Replace(strings.Replace(input, "\n", " ", -1), "  ", " ", -1), " ")
	for i, num := range parseInput {
		b.positions[i] = pos{
			x:     i % cols,
			y:     i / cols,
			value: getValue(num),
		}
	}
	return b
}

func getValue(num string) int {
	v, _ := strconv.Atoi(num)
	return v
}

func main() {
	fmt.Println("2021 Advent of Code Day 4")
	fmt.Println("--- Part 1 ---")

	input := readInput("test.txt")
	// sanitize
	inputSplitted := strings.Split(strings.Replace(input, "  ", " ", 0), "\n\n")

	var boards []board

	//Append all boards
	for _, b := range inputSplitted[1:] {
		boards = append(boards, parseBoard(b))
	}

	fmt.Printf("Random numbers:\n%s\n", inputSplitted[0])
	fmt.Printf("First board:\n%s\n", inputSplitted[1])
	fmt.Printf("Found board count: %d\n", len(boards))

	//fmt.Println("--- Part 2 ---")

	fmt.Println("Fin.")

}

func readInput(file string) string {

	if data, err := ioutil.ReadFile(file); err == nil {
		return string(data)

	}
	return ""
}

/*

func invertBinaryString(input string) string {
	input = strings.Replace(input, "1", "2", -1)
	input = strings.Replace(input, "0", "1", -1)
	return strings.Replace(input, "2", "0", -1)
}
func findCommonBitInColumn(input []string, column int) ([]string, []string) {

	ones, zeroes := []string{}, []string{}

	for i := 0; i < len(input); i++ {
		if input[i][column] == '1' {
			ones = append(ones, input[i])
		} else if input[i][column] == '0' {
			zeroes = append(zeroes, input[i])

		}
	}
	return zeroes, ones
}

func findCommonBits(input []string) (string, error) {

	result := ""
	length := len(input[0])
	rows := len(input)

	for i := 0; i < length; i++ {

		countSum := 0

		for _, s := range input {
			if s[i] == '1' {
				countSum++
			}
		}

		if float32(countSum)*1.0 > float32(rows)*0.5 {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result, nil
}

func binaryStringToInt(binaryString string) int64 {
	if i, err := strconv.ParseInt(binaryString, 2, 64); err == nil {
		return i
	}
	return -1
}
*/
