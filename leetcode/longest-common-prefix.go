func GetChar(strs *[]string, i int) (byte, bool) {
	if len((*strs)[0]) == i {
		return 'a', false
	}
	c := (*strs)[0][i]
	for _, s := range *strs {
		if len(s) == i || s[i] != c {
			return c, false
		}
	}
	return c, true
}

func longestCommonPrefix(strs []string) string {
	output := []byte{}
	var c byte
	valid, i := true, 0
	for valid {
		output = append(output, c)
		c, valid = GetChar(&strs, i)
		i++
	}
	return string(output[1:])
}
