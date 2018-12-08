package day8

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type node struct {
	Nodes []node
	Meta  []int // Metadata
}

func parseLicense(filename string) []int {
	rawLogs, _ := ioutil.ReadFile(filename)
	data := strings.Split(string(rawLogs), " ")
	stream := []int{}
	for i := 0; i < len(data); i++ {
		nt, _ := strconv.Atoi(strings.Replace(data[i], "\n", "", 1))
		stream = append(stream, nt)
	}
	return stream
}

func buildTree(stream []int) (node, []int) {
	n := node{}
	children := stream[0]
	data := stream[1]
	stream = stream[2:] // Consume the children value from the stream

	for i := 0; i < children; i++ {
		var nd node
		nd, stream = buildTree(stream)
		n.Nodes = append(n.Nodes, nd)
	}

	for j := 0; j < data; j++ {
		n.Meta = append(n.Meta, stream[j])
	}

	// Consume the metadata from the stream
	stream = stream[data:]
	return n, stream
}

func sumMetadata(n node) int {
	sum := 0
	for i := 0; i < len(n.Nodes); i++ {
		sum += sumMetadata(n.Nodes[i])
	}

	for i := 0; i < len(n.Meta); i++ {
		sum += n.Meta[i]
	}
	return sum
}

func nodeValue(n node) int {
	sum := 0
	totalNodes := len(n.Nodes)
	totalMeta := len(n.Meta)

	if totalNodes > 0 {

		for i := 0; i < totalMeta; i++ {
			index := n.Meta[i] - 1
			if totalNodes > index {
				sum += nodeValue(n.Nodes[index])
			}
		}

	} else {

		for i := 0; i < totalMeta; i++ {
			sum += n.Meta[i]
		}

	}

	return sum
}

// Solve day 8
func Solve() {
	stream := parseLicense("./day8/license")

	// Assuming only one root node in the stream
	n, _ := buildTree(stream)
	fmt.Printf("Part 1: %d\n", sumMetadata(n))
	fmt.Printf("Part 2: %d", nodeValue(n))
}
