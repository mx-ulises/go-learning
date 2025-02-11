func isSuffix(output []byte, part string) bool {
    n, m := len(output), len(part)
    if n < m {
        return false
    }
    offset := n - m
    for i := 0; i < m; i++ {
        if output[offset+i] != part[i] {
            return false
        }
    }
    return true
}

func removeOccurrences(s string, part string) string {
    output := []byte{}
    for i := 0; i < len(s); i++ {
        output = append(output, s[i])
        for isSuffix(output, part) {
            output = output[0:len(output)-len(part)]
            fmt.Println(string(output))
        }
    }
    return string(output)
}
