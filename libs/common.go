package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Node : node  with x, y co-ordinate
type Node struct {
	X        int
	Y        int
	Previous *Node
}

//ReadArray : creates two diamentional array as per user input
func ReadArray() ([][]int, int, int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Number of rows ? ")
	row, _ := reader.ReadString('\n')
	row = strings.Replace(row, "\n", "", -1)
	rows, _ := strconv.Atoi(row)

	fmt.Print("Number of columns ? ")
	col, _ := reader.ReadString('\n')
	col = strings.Replace(col, "\n", "", -1)
	cols, _ := strconv.Atoi(col)

	fmt.Println("Enter", rows*cols, "numbers for 2 diamentional array")
	num, _ := reader.ReadString('\n')
	num = strings.Trim(strings.Replace(num, "\n", "", -1), "\n ")
	nums := strings.Split(num, " ")

	array := make([][]int, rows)
	for i := range array {
		array[i] = make([]int, cols)
		for j := range array[i] {
			val, _ := strconv.Atoi(nums[(i*cols)+j])
			array[i][j] = val
			fmt.Print(array[i][j], " ")
		}
		fmt.Println()
	}

	return array, rows, cols
}

//PrintArray : Prints array
func PrintArray(array [][]int) {
	for i := range array {
		for j := range array[i] {
			fmt.Print(array[i][j], " ")
		}
		fmt.Println()
	}
}

//CreateExplorationArray : creates an array with zeros of given rows and columns
func CreateExplorationArray(rows int, cols int) [][]int {
	array := make([][]int, rows)
	for i := range array {
		array[i] = make([]int, cols)
		for j := range array[i] {
			array[i][j] = 0
		}
	}
	return array
}

//CloneNode : clones the Node
func CloneNode(n *Node) Node {
	m := &Node{X: n.X, Y: n.Y}
	r := m
	for n.Previous != nil {
		m.Previous = &Node{X: n.Previous.X, Y: n.Previous.Y}
		m = m.Previous
		n = n.Previous
	}
	return *r
}

//PrintNodes : Prints the nodes
func PrintNodes(n Node) {
	printNodes(n)
	fmt.Println()
}

func printNodes(n Node) {
	if n.Previous != nil {
		printNodes(*n.Previous)
		fmt.Print(" <- (", n.X, ",", n.Y, ")")
	} else {
		fmt.Print("(", n.X, ",", n.Y, ")")
	}
}

//NodeLength : counts the length
func NodeLength(node *Node) int {
	len := 1
	for node.Previous != nil {
		len++
		node = node.Previous
	}
	return len
}
