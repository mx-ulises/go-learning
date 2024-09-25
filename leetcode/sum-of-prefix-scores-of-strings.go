type Trie struct {
	Children map[byte]*Trie
	Score    int
}

func Add(trie *Trie, s *string, i int) {
	n := len(*s)
	if i == n {
		return
	}
	c := (*s)[i]
	_, ok := trie.Children[c]
	if !ok {
		trie.Children[c] = &Trie{Children: map[byte]*Trie{}}
	}
	trie.Children[c].Score++
	Add(trie.Children[c], s, i+1)
}

func GetScore(trie *Trie, s *string, i int) int {
	n := len(*s)
	if i == n {
		return 0
	}
	c := (*s)[i]
	_, ok := trie.Children[c]
	if !ok {
		return 0
	}
	return trie.Children[c].Score + GetScore(trie.Children[c], s, i+1)
}

func sumPrefixScores(words []string) []int {
	trie := &Trie{Children: map[byte]*Trie{}}
	for _, s := range words {
		Add(trie, &s, 0)
	}
	scores := make([]int, len(words))
	for i, s := range words {
		scores[i] = GetScore(trie, &s, 0)
	}
	return scores
}
