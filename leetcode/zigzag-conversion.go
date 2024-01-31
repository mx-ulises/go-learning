func GetPositionMap(numRows int) []int {
	positionMap := []int{}
	for i := 0; i < numRows; i++ {
		positionMap = append(positionMap, i)
	}
	for i := 2; i < numRows; i++ {
		positionMap = append(positionMap, numRows-i)
	}
	return positionMap
}

func convert(s string, numRows int) string {
	positionMap := GetPositionMap(numRows)
	keys := len(positionMap)
	stringBuilder := make([][]byte, numRows)
	for i := 0; i < len(s); i++ {
		key := i % keys
		row := positionMap[key]
		stringBuilder[row] = append(stringBuilder[row], s[i])
	}
	output := []byte{}
	for i := 0; i < len(stringBuilder); i++ {
		output = append(output, stringBuilder[i]...)
	}
	return string(output)
}
