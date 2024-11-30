/*
2097. Valid Arrangement of Pairs

You are given a 0-indexed 2D integer array pairs where pairs[i] = [starti, endi]. An arrangement of pairs is valid if for every index i where 1 <= i < pairs.length, we have endi-1 == starti.

Return any valid arrangement of pairs.

Note: The inputs will be generated such that there exists a valid arrangement of pairs.

Example 1:

Input: pairs = [[5,1],[4,5],[11,9],[9,4]]
Output: [[11,9],[9,4],[4,5],[5,1]]
Explanation:
This is a valid arrangement since endi-1 always equals starti.
end0 = 9 == 9 = start1
end1 = 4 == 4 = start2
end2 = 5 == 5 = start3

https://leetcode.com/problems/valid-arrangement-of-pairs/
*/
package main

import "fmt"

func validArrangement(pairs [][]int) [][]int {

	adjacencyList := make(map[int][]int)
	inOutDegree := make(map[int]int)

	for _, pair := range pairs {
		adjacencyList[pair[0]] = append(adjacencyList[pair[0]], pair[1])
		inOutDegree[pair[0]]++
		inOutDegree[pair[1]]--
	}

	fmt.Println(adjacencyList)
	fmt.Println(inOutDegree)

	startNode := pairs[0][0]
	for node, degree := range inOutDegree {
		if degree == 1 {
			startNode = node
			break
		}
	}

	path := []int{}
	nodeStack := []int{startNode}

	for len(nodeStack) > 0 {
		lastIdx := len(nodeStack) - 1
		curr := nodeStack[lastIdx]
		neighbors := adjacencyList[curr]

		if len(neighbors) == 0 {
			path = append(path, curr)
			nodeStack = nodeStack[:lastIdx]
		} else {
			nextNode := neighbors[len(neighbors)-1]
			nodeStack = append(nodeStack, nextNode)
			adjacencyList[curr] = neighbors[:len(neighbors)-1]
		}
	}

	fmt.Println(path)
	arrangement := make([][]int, 0, len(path)-1)
	for i := len(path) - 1; i > 0; i-- {
		arrangement = append(arrangement, []int{path[i], path[i-1]})
	}

	return arrangement
}

// func main() {
// 	var pairs = [][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}
// 	p := validArrangement(pairs)

// 	fmt.Println(p)
// }
