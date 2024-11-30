func freqAlphabets(s string) string {
	intToRune := map[string]rune{
		"1": 'a', "2": 'b', "3": 'c', "4": 'd', "5": 'e',
		"6": 'f', "7": 'g', "8": 'h', "9": 'i', "10": 'j',
		"11": 'k', "12": 'l', "13": 'm', "14": 'n', "15": 'o',
		"16": 'p', "17": 'q', "18": 'r', "19": 's', "20": 't',
		"21": 'u', "22": 'v', "23": 'w', "24": 'x', "25": 'y',
		"26": 'z',
	}
	stack := []rune{}
	output := []rune{}
	for _, c := range s {
		if c == '#' {
			i := 0
			for i = 0; i < len(stack)-2; i++ {
				output = append(output, intToRune[string(stack[i])])
			}
			output = append(output, intToRune[string(stack[i:])])
			stack = stack[0:0]
		} else {
			stack = append(stack, c)
		}
	}
	for i := 0; i < len(stack); i++ {
		output = append(output, intToRune[string(stack[i])])
	}
	return string(output)
}
