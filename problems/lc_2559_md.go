package main

// func vowelStrings_(words []string, queries [][]int) []int {

// 	m := map[byte]bool{
// 		'a': true,
// 		'e': true,
// 		'i': true,
// 		'o': true,
// 		'u': true,
// 	}

// 	ans := make([]int, 0, len(queries))

// 	for _, q := range queries {

// 		c := 0

// 		for _, s := range words[q[0]:q[1]+1] {

// 			if m[s[0]] && m[s[len(s)-1]] {
// 				c++
// 			}
// 		}

// 		ans = append(ans, c)
// 	}

// 	return ans
// }

// way 2

// func vowelStrings(words []string, queries [][]int) []int {

//     vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

//     prefixSum := make([]int, len(words) + 1)

//     for i, word := range words {
//         if vowels[word[0]] && vowels[word[len(word)-1]] {
//             prefixSum[i + 1] = prefixSum[i] + 1
//         } else {
//             prefixSum[i + 1] = prefixSum[i]
//         }
//     }

//     result := []int{}
//     for _, query := range queries {
//         start, end := query[0], query[1]
//         result = append(result, prefixSum[end + 1] - prefixSum[start])
//     }

//     return result
// }

// way 3
func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func vowelStrings(words []string, queries [][]int) []int {

	prefix := make([]int, len(words)) // create prefix array to store the number of vowels in the words

	for idx, word := range words { // iterate over the words
		prefix[idx] = 0 // initialize the prefix array
		if idx > 0 { // if the index is greater than 0, then we need to add the previous prefix to the current prefix to get the total number of vowels in the words
			prefix[idx] = prefix[idx-1]
		}

		if isVowel(word[0]) && isVowel(word[len(word)-1]) { // if the first and last character of the word is a vowel, then we need to add 1 to the prefix array
			prefix[idx]++
		}
	}

	ans := make([]int, len(queries)) // create the answer array to store the number of vowels in the words

	for idx, query := range queries { // iterate over the queries
		l := query[0]
		r := query[1]
		if l == 0 { // if the left index is 0, then we need to add the prefix at the right index to the answer array
			ans[idx] = prefix[r]
			continue
		}
		ans[idx] = prefix[r] - prefix[l-1] // if the left index is not 0, then we need to subtract the prefix at the left index from the prefix at the right index to get the number of vowels in the words between the left and right index		
	}
	
	return ans
}

// func main() {

// 	words := []string{"aba", "bcb", "ece", "aa", "e"}
// 	queries := [][]int{{0, 2}, {1, 4}, {1, 1}}

// 	fmt.Println(vowelStrings(words, queries))
// }
