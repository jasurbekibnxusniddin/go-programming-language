// You can edit this code!
// Click here and start typing.
package main

func maximumLength(s string) int {

	var (
		m   = make(map[string]int)
		str = ""
	)

	for _, r := range s {

		m[string(r)]++

		if len(s) == 0 || (len(s) > 0 && s[0] == byte(r)) {
			str += string(r)
		} else {
			m[str]++
			str = ""
		}
	}

	m[str]++

	max := 0
	lmax := 0

	for l, v := range m {
		if v > max {
			max = v
			lmax = len(l)
		}
	}

	if max >= 3 {
		return lmax
	}

	return -1
}

func main() {
	maximumLength("aaa")
}
