type Trie struct {
	Digits map[int]*Trie
}

func Add(trie *Trie, digits *[]int) {
	n := len(*digits)
	if n == 0 {
		return
	}
	d := (*digits)[n-1]
	(*digits) = (*digits)[:n-1]
	_, ok := trie.Digits[d]
	if !ok {
		trie.Digits[d] = &Trie{Digits: map[int]*Trie{}}
	}
	Add(trie.Digits[d], digits)
}

func SuffixSize(trie *Trie, digits *[]int) int {
	n := len(*digits)
	if n == 0 {
		return 0
	}
	d := (*digits)[n-1]
	(*digits) = (*digits)[:n-1]
	_, ok := trie.Digits[d]
	if !ok {
		return 0
	}
	return 1 + SuffixSize(trie.Digits[d], digits)
}

func GetDigits(num int) *[]int {
	digits := []int{}
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	return &digits
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	trie := &Trie{Digits: map[int]*Trie{}}
	for _, num := range arr1 {
		digits := GetDigits(num)
		Add(trie, digits)
	}
	maximal := 0
	for _, num := range arr2 {
		digits := GetDigits(num)
		if len(*digits) <= maximal {
			continue
		}
		maximal = max(maximal, SuffixSize(trie, digits))
	}
	return maximal
}
