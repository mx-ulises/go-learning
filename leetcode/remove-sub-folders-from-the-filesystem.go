type TrieNode struct {
	TerminalValue *string
	Children      map[string]*TrieNode
}

func AddAndCheck(node *TrieNode, folderStructure *[]string, i int, f *string) bool {
	if node.TerminalValue != nil {
		return false
	}
	if i == len(*folderStructure) {
		node.TerminalValue = f
		return true
	}
	subfolder := (*folderStructure)[i]
	_, ok := node.Children[subfolder]
	if ok == false {
		node.Children[subfolder] = &TrieNode{Children: map[string]*TrieNode{}}
	}
	return AddAndCheck(node.Children[subfolder], folderStructure, i+1, f)
}

func GetFolderStructure(f *string) *[]string {
	folderStructure := []string{}
	current := []rune{}
	for _, c := range *f {
		if c == '/' {
			folderStructure = append(folderStructure, string(current))
			current = []rune{}
		} else {
			current = append(current, c)
		}
	}
	folderStructure = append(folderStructure, string(current))
	return &folderStructure
}

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	folders := []string{}
	root := TrieNode{Children: map[string]*TrieNode{}}
	for _, f := range folder {
		folderStructure := GetFolderStructure(&f)
		if AddAndCheck(&root, folderStructure, 0, &f) == true {
			folders = append(folders, f)
		}
	}
	return folders
}
