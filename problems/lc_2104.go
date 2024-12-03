package main

import "math"

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}

func subArrayRanges(nums []int) int64 {

	sum := int64(0)

	for si := 0; si < len(nums); si++ {
		smin := math.MaxInt
		smax := math.MinInt

		for ei := si; ei < len(nums); ei++ {
			smin = min(smin, nums[ei])
			smax = max(smax, nums[ei])

			sum += int64(smax - smin)
		}
	}

	return sum
}

// func subArrayRanges(nums []int) int64 {
//     n := len(nums)
//     var sum int64 = 0

//     maxStack := []int{}
//     for i := 0; i <= n; i++ {
//         for len(maxStack) > 0 && (i == n || nums[maxStack[len(maxStack)-1]] < nums[i]) {
//             mid := maxStack[len(maxStack)-1]
//             maxStack = maxStack[:len(maxStack)-1]
//             left := -1
//             if len(maxStack) > 0 {
//                 left = maxStack[len(maxStack)-1]
//             }
//             right := i
//             sum += int64(nums[mid]) * int64(mid-left) * int64(right-mid)
//         }
//         maxStack = append(maxStack, i)
//     }

//     minStack := []int{}
//     for i := 0; i <= n; i++ {
//         for len(minStack) > 0 && (i == n || nums[minStack[len(minStack)-1]] > nums[i]) {
//             mid := minStack[len(minStack)-1]
//             minStack = minStack[:len(minStack)-1]
//             left := -1
//             if len(minStack) > 0 {
//                 left = minStack[len(minStack)-1]
//             }
//             right := i
//             sum -= int64(nums[mid]) * int64(mid-left) * int64(right-mid)
//         }
//         minStack = append(minStack, i)
//     }

//     return sum
// }

// func subArrayRanges(nums []int) int64 {

// 	var sum int64 = 0

// 	for i := 0; i < len(nums); i++ {
// 		for j := i+2; j <= len(nums); j++ {
// 			sum += dif(nums[i:j])
// 		}
// 	}

// 	return sum
// }

// func dif(slc []int) int64 {

//     max := slc[0]
// 	min := slc[0]

// 	for _, e := range slc {
// 		if e > max {
// 			max = e
// 		}

// 		if min > e {
// 			min = e
// 		}
// 	}

// 	return int64(max - min)
// }
