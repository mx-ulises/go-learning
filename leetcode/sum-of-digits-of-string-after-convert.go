func GetFirstNumber(s string) string {
	number := []string{}
	for _, c := range s {
		num := int(c-'a') + 1
		number = append(number, strconv.Itoa(num))
	}
	return strings.Join(number, "")
}

func GetNewNumber(number string) string {
	new_number := 0
	for _, d := range number {
		new_number += int(d - '0')
	}
	return strconv.Itoa(new_number)
}

func getLucky(s string, k int) int {
	number := GetFirstNumber(s)
	for 0 < k {
		number = GetNewNumber(number)
		k--
	}
	output, _ := strconv.Atoi(number)
	return output
}
