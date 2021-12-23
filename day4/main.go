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
	won       bool
	columns   int
	rows      int
	positions []*pos
}

func parseBoard(input string) *board {
	input = strings.Trim(input, " ")
	rows := len(strings.Split(input, "\n"))
	clean := strings.Split(input, "\n")[0]
	clean = strings.Replace(clean, "  ", " ", -1)
	cols := len(strings.Split(clean, " "))

	b := &board{
		won:       false,
		columns:   cols,
		rows:      rows,
		positions: make([]*pos, rows*cols),
	}

	parseInput := strings.Split(strings.Replace(strings.Replace(input, "\n", " ", -1), "  ", " ", -1), " ")
	for i, num := range parseInput {
		b.positions[i] = &pos{
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

	input := readInput("input.txt")
	// sanitize
	inputSplitted := strings.Split(strings.Replace(input, "  ", " ", 0), "\n\n")

	var boards []*board
	numbers := strings.Split(inputSplitted[0], ",")
	//Append all boards
	for _, b := range inputSplitted[1:] {
		boards = append(boards, parseBoard(b))
	}

	fmt.Printf("Random numbers:\n%s\n", numbers)
	fmt.Printf("First board:\n%s\n", inputSplitted[1])
	fmt.Printf("Found board count: %d\n", len(boards))

	playBingoSubSystem(boards, numbers)

	fmt.Println("--- Part 2 ---")

	playBingoSubSystemV2(boards, numbers)
	fmt.Println("Fin.")

}

func playBingoSubSystem(boards []*board, numbers []string) {

	for _, nums := range numbers {
		num := getValue(nums)

		for i, board := range boards {
			markNumberOnBoard(board, num)
			if checkWinConditions(*board) {
				fmt.Printf("BINGO! Board %d won!\n", i+1)
				fmt.Printf("Final score %d \n", calculateWinnerScore(*board)*num)
				return
			}
		}
	}
}

func playBingoSubSystemV2(boards []*board, numbers []string) {

	for _, nums := range numbers {
		num := getValue(nums)

		for _, board := range boards {
			markNumberOnBoard(board, num)
			if !board.won && checkWinConditions(*board) {
				board.won = true

				if allBoardsWon(boards) {
					fmt.Printf("Final score %d \n", calculateWinnerScore(*board)*num)
				}
			}
		}
	}

}

func allBoardsWon(boards []*board) bool {
	for _, b := range boards {
		if !b.won {
			return false
		}
	}
	return true
}

func calculateWinnerScore(b board) int {
	sum := 0
	for _, p := range b.positions {
		if !p.marked {
			sum += p.value
		}
	}
	return sum
}

func (b board) valueAt(col, row int) pos {
	return *b.positions[(row*b.columns)+col]
}

func checkWinConditions(b board) bool {

	for col := 0; col < b.columns; col++ {
		count := 0
		for row := 0; row < b.rows; row++ {
			if b.valueAt(col, row).marked {
				count++
			}
		}
		if count == b.rows {
			return true
		}
	}

	for row := 0; row < b.rows; row++ {
		count := 0
		for col := 0; col < b.columns; col++ {
			if b.valueAt(col, row).marked {
				count++
			}
		}
		if count == b.columns {
			return true
		}
	}

	return false
}

func markNumberOnBoard(b *board, num int) {
	for _, pos := range b.positions {
		if pos.value == num {
			pos.marked = true
		}
	}
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
