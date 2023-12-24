func InitializePhoneMap() *map[byte][]byte {
	phoneMap := make(map[byte][]byte)
	phoneMap['2'] = []byte{'a', 'b', 'c'}
	phoneMap['3'] = []byte{'d', 'e', 'f'}
	phoneMap['4'] = []byte{'g', 'h', 'i'}
	phoneMap['5'] = []byte{'j', 'k', 'l'}
	phoneMap['6'] = []byte{'m', 'n', 'o'}
	phoneMap['7'] = []byte{'p', 'q', 'r', 's'}
	phoneMap['8'] = []byte{'t', 'u', 'v'}
	phoneMap['9'] = []byte{'w', 'x', 'y', 'z'}
	return &phoneMap
}

func GetCombinations(phoneMap *map[byte][]byte, combinations *[]string, preffix *[]byte, digits *string) {
	if len(*preffix) == len(*digits) {
		(*combinations) = append(*combinations, string(*preffix))
	} else {
		d := (*digits)[len(*preffix)]
		for i := 0; i < len((*phoneMap)[d]); i++ {
			c := (*phoneMap)[d][i]
			(*preffix) = append((*preffix), c)
			GetCombinations(phoneMap, combinations, preffix, digits)
			(*preffix) = (*preffix)[:len(*preffix)-1]
		}
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	phoneMap := InitializePhoneMap()
	combinations := make([]string, 0)
	preffix := make([]byte, 0)
	GetCombinations(phoneMap, &combinations, &preffix, &digits)
	return combinations
}
