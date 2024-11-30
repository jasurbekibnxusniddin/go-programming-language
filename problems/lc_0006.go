package main

func maxArea(height []int) int {

	var (
		v = 0
		l = len(height)
	)

	if l < 1 {
		return v
	}

	for i := 0; i < l; i++ {

		minH := height[i]
		println(l - 1 - i)
		h := height[l-1-i]

		if minH > h {
			minH = h
		}
		cl := h - minH

		if v < minH*cl {
			v = minH * cl
		}
	}

	return v
}

func main() {
	maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
}