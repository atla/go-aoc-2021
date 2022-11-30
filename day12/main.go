package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type CaveNode struct {
	ID          string
	Connections map[string]*CaveNode
}

func NewCaveNode(id string) *CaveNode {
	return &CaveNode{
		ID:          id,
		Connections: make(map[string]*CaveNode),
	}
}

type DistinctPath struct {
	Nodes []CaveNode
}

type CaveSystem struct {
	Start    CaveNode
	End      CaveNode
	AllNodes []*CaveNode
}

func (cv *CaveSystem) AddSegment(a, b string) {

	nodeA := cv.FindOrCreate(a)
	nodeB := cv.FindOrCreate(b)

	// add connection in both directions
	nodeA.Connections[nodeB.ID] = nodeB
	nodeB.Connections[nodeA.ID] = nodeA
}

func (cv *CaveSystem) FindOrCreate(id string) *CaveNode {

	var node *CaveNode = nil

	for _, n := range cv.AllNodes {
		if n.ID == id {
			node = n
			break
		}
	}
	if node == nil {

		node = &CaveNode{
			ID:          id,
			Connections: make(map[string]*CaveNode),
		}
		cv.AddNode(node)
	}
	return node
}

func (cv *CaveSystem) AddNode(node *CaveNode) {
	cv.AllNodes = append(cv.AllNodes, node)
}

func main() {
	fmt.Println("2021 Advent of Code Day 12")
	fmt.Println("--- Part 1 ---")

	input := readInput("test.txt")

	lines := strings.Split(input, "\n")

	caveSystem := CaveSystem{
		AllNodes: []*CaveNode{},
	}

	for _, line := range lines {
		split := strings.Split(line, "-")
		caveSystem.AddSegment(split[0], split[1])
	}

	fmt.Println("Fin.")
}

func readInput(file string) string {
	if data, err := ioutil.ReadFile(file); err == nil {
		return string(data)

	}
	return ""
}
