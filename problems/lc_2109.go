package main

func addSpaces(s string, spaces []int) string {

	var bytes = make([]byte, 0, len(s)+len(spaces))

	for i, j := 0, 0; i < len(s); i++ {

		if j < len(spaces) && i == spaces[j] {
			bytes = append(bytes, ' ')
			j++
		}

		bytes = append(bytes, s[i])
	}

	return string(bytes)
}

/*
func addSpaces(s string, spaces []int) string {

    var bytes = []byte{}
    var m = map[int]bool{}

    for _, e := range spaces  {
        m[e] = true
    }

    for i := 0; i < len(s); i++ {

        if m[i] {
            bytes = append(bytes, ' ')
        }

        bytes = append(bytes, s[i])
    }

    return string(bytes)
}
*/
