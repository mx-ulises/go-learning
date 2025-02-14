func stringHash(s string, k int) string {
    output := []byte{}
    i := 0
    for i < len(s) {
        current := 0;
        for j := 0; j < k; j++ {
            current += int(s[i] - 'a')
            i++
        }
        output = append(output, 'a' + byte(current % 26))
    }
    return string(output)
}
