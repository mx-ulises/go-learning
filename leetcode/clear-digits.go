func clearDigits(s string) string {
    output := []rune{}
    for _, c := range s {
        if '0' <= c && c <= '9' {
            output = output[:len(output)-1]
        } else {
            output = append(output, c)
        }
    }
    return string(output)
}
