package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

//() = 1 [] = 2  {} =3 <> = 4

type stack []uint8

func (s stack) Push(v uint8) stack {
	return append(s, v)
}

func (s stack) IsEmpty() bool {
	return len(s) == 0
}

func (s stack) Pop() (stack, uint8) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}

func main() {
	fmt.Println("2021 Advent of Code Day 10 - Syntax Scoring")
	fmt.Println("--- Part 1 ---")

	input := readInput("input.txt")

	autoCompleteScores := []int{}

	lines := strings.Split(input, "\n")
	sum := 0
	for l, line := range lines {

		stack := make(stack, 0)

	loop:
		for i := 0; i < len(line); i++ {

			switch {
			case line[i] == '(':
				stack = stack.Push(1)
			case line[i] == ')':
				if stack.IsEmpty() {
					break loop
				}
				s2, p := stack.Pop()
				stack = s2

				if p != 1 {
					fmt.Printf("Line %d, Expected %s, but found ) instead.\n", l, convRune(p))
					sum += 3
					break loop
				}

			case line[i] == '[':
				stack = stack.Push(2)
			case line[i] == ']':
				if stack.IsEmpty() {
					break loop
				}
				s2, p := stack.Pop()
				stack = s2

				if p != 2 {
					fmt.Printf("Line %d, Expected %s, but found ] instead.\n", l, convRune(p))
					sum += 57
					break loop
				}
			case line[i] == '{':
				stack = stack.Push(3)
			case line[i] == '}':
				if stack.IsEmpty() {
					break
				}
				s2, p := stack.Pop()
				stack = s2

				if p != 3 {
					fmt.Printf("Line %d, Expected %s, but found } instead.\n", l, convRune(p))
					sum += 1197
					break loop
				}
			case line[i] == '<':
				stack = stack.Push(4)
			case line[i] == '>':
				if stack.IsEmpty() {
					break loop
				}
				s2, p := stack.Pop()
				stack = s2

				if p != 4 {
					fmt.Printf("Line %d, Expected %s, but found > instead.\n", l, convRune(p))
					sum += 25137
					break loop
				}
			}

			if i == len(line)-1 {
				fmt.Println("Last character in line")

				if len(stack) > 0 {

					totalScore := 0

					for !stack.IsEmpty() {

						totalScore *= 5
						s2, v := stack.Pop()
						stack = s2
						switch v {
						case 1:
							totalScore += 1
						case 2:
							totalScore += 2
						case 3:
							totalScore += 3
						case 4:
							totalScore += 4

						}
					}
					//
					autoCompleteScores = append(autoCompleteScores, totalScore)
				}
			}
		}
	}
	fmt.Printf("Total Syntax Error Score: %d\n", sum)

	sort.Ints(autoCompleteScores)
	pos := len(autoCompleteScores) / 2
	fmt.Println("winner  is ", autoCompleteScores[pos])

	fmt.Println("Fin.")
}

func convRune(p uint8) string {
	switch p {
	case 1:
		return ")"
	case 2:
		return "]"
	case 3:
		return "}"
	case 4:
		return ">"
	}
	return ""
}

func valueFrom(num string) int {
	v, _ := strconv.Atoi(num)
	return v
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
