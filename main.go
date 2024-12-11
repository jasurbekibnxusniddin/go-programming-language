// You can edit this code!
// Click here and start typing.
package main

// func maximumLength(s string) int {
// 	var (
// 		m   = make(map[string]int)
// 		str = ""
// 	)

// 	for i := 0; i < len(s); i++ { // Use byte iteration
// 		b := s[i] // Current byte
// 		if str == "" || (len(str) > 0 && str[0] == b) {
// 			str += string(b)
// 		} else {
// 			m[str]++
// 			str = string(b)
// 		}
// 	}

// 	if str != "" {
// 		m[str]++
// 	}

// 	fmt.Println(m)
// 	max := 0
// 	lmax := 0
// 	maxL := 0

// 	for l, v := range m {
// 		if v >= 3 && len(l) > lmax {
// 			max = v
// 			lmax = len(l)
// 		}

// 		if maxL < len(l) {
// 			maxL = len(l)
// 		}
// 	}

// 	if max >= 3 {
// 		return lmax
// 	} else if maxL >= 3 {
// 		return maxL - 2
// 	}

// 	return -1
// }

func maximumLength(s string) int {
	
	n := len(s)
	result := -1

	isSpecial := func(substr string) bool {
		
		for i := 1; i < len(substr); i++ {
		
			if substr[0] != substr[1] {
		
				return false
		
			}
		
		}
		
		return true
	}

	for length := 1; length <= n; length++ {
		
		counts := make(map[string]int)
		
		for i := 0; i <= n-length; i++ {
			
			sub := s[i : i+length]
			
			if isSpecial(sub) {
				
				counts[sub]++
				
				if counts[sub] >= 3 {
					
					result = length
				}
			}
		}
	}

	return result
}

func main() {
	l := maximumLength("abcaba")

	println(l)
}
