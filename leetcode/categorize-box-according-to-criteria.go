package leetcode

func isBulky(length, width, height int) bool {
	if 10_000 <= length || 10_000 <= width || 10_000 <= height {
		return true
	}
	return 1_000_000_000 <= (length * width * height)
}

func isHeavy(mass int) bool {
	return 100 <= mass
}

func categorizeBox(length int, width int, height int, mass int) string {
	bulky, heavy := isBulky(length, width, height), isHeavy(mass)
	if bulky && heavy {
		return "Both"
	} else if bulky {
		return "Bulky"
	} else if heavy {
		return "Heavy"
	}
	return "Neither"
}
