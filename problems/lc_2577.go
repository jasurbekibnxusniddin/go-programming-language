package main

import "container/heap"

type Cell struct {
	time, row, col int
}

type Cells []*Cell

func (c Cells) Len() int {
	return len(c)
}

func (c Cells) Less(i, j int) bool {
	return c[i].time < c[j].time
}

func (c Cells) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *Cells) Push(x any) {
	*c = append(*c, x.(*Cell))
}

func (c *Cells) Pop() any {
	old := *c
	n := len(old)
	cell := old[n-1]
	old[n-1] = nil
	*c = old[0 : n-1]
	return cell
}

func minimumTime(G [][]int) int {

	if G[0][1] > 1 && G[1][0] > 1 {
		return -1
	}

	row, col := len(G), len(G[0])
	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	V := make([][]bool, row)

	for i := range V {
		V[i] = make([]bool, col)
	}

	V[0][0] = true

	Q := &Cells{}
	heap.Push(Q, &Cell{0, 0, 0})

	for Q.Len() > 0 {
		cell := heap.Pop(Q).(*Cell)

		for _, d := range dirs {
			r, c := cell.row+d[0], cell.col+d[1]
			if 0 <= r && r < row && 0 <= c && c < col && !V[r][c] {
				t := cell.time + 1
				if t < G[r][c] {
					t = G[r][c] + (G[r][c]-t)%2
				}

				if r == row-1 && c == col-1 {
					return t
				}
				heap.Push(Q, &Cell{t, r, c})
				V[r][c] = true
			}
		}
	}

	return -1
}
