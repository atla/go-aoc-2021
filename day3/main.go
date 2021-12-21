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
	fmt.Println("2021 Advent of Code Day 3")
	fmt.Println("--- Part 1 ---")

	input := readInput("input.txt")

	if result, err := findCommonBits(input); err == nil {
		r1 := binaryStringToInt(result)
		r2 := binaryStringToInt(invertBinaryString(result))
		fmt.Printf("Result %d \n", r1*r2)
	}

	fmt.Println("--- Part 2 ---")
	var oxyGen int64
	var scrubberco2 int64

	l := len(input[0])
	in := input
	for i := 0; i < l; i++ {
		zeroes, ones := findCommonBitInColumn(in, i)

		if len(ones) >= len(zeroes) {
			in = ones
		} else {
			in = zeroes
		}
		if len(in) == 1 {
			oxyGen = binaryStringToInt(in[0])
			fmt.Printf("--- Found oxygenGeneratorValue %d\n", oxyGen)
			break
		}
	}

	in2 := input
	for i := 0; i < l; i++ {
		zeroes, ones := findCommonBitInColumn(in2, i)
		if len(ones) >= len(zeroes) {
			in2 = zeroes
		} else {
			in2 = ones
		}
		if len(in2) == 1 {

			scrubberco2 = binaryStringToInt(in2[0])
			fmt.Printf("--- Found CO2 scrubber %d\n", scrubberco2)

			break
		}
	}
	fmt.Printf("--- Life support rating %d\n", oxyGen*scrubberco2)

	fmt.Println("Fin.")

}

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

func readInput(file string) []string {

	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n")
	}
	return nil
}
