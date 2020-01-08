package main

import (
	common "github.com/govishwavijay/learning/libs"
)

type slice []int

func main() {
	array, _, _ := common.ReadArray()

	start := &common.Node{X: 0, Y: 0}

	common.PrintArray(array)

	var ll []int

	explore(&array, start, &ll)

}

func explore(array *[][]int, node *common.Node, found *[]int) {
	if array[node.X][node.Y] == 9 {
		append(found, common.NodeLength(node))
	}
}

func isLoop(node *common.Node) bool {
	m := node.Previous
	for m != nil {
		if m.X == node.X && m.Y == node.Y {
			return true
		}
	}
	return false
}
