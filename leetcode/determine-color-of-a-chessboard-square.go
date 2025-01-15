func squareIsWhite(coordinates string) bool {
	col := coordinates[0] - 'a'
	row := coordinates[1] - '1'
	return ((col ^ row) & 1) == 1
}
