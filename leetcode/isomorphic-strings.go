func CheckAndAdd(s byte, t byte, charMap *map[byte]byte) bool {
	c, ok := (*charMap)[s]
	if ok {
		if c != t {
			return false
		}
	}
	(*charMap)[s] = t
	return true
}

func isIsomorphic(s string, t string) bool {
	sToT := map[byte]byte{}
	tToS := map[byte]byte{}
	n := len(s)
	for i := 0; i < n; i++ {
		if (CheckAndAdd(s[i], t[i], &sToT) && CheckAndAdd(t[i], s[i], &tToS)) == false {
			return false
		}
	}
	return true
}
