func GetKeyMap(key *string) *map[byte]byte {
	ALPHABET := "abcdefghijklmnopqrstuvwxyz"
	keyMap := map[byte]byte{' ': ' '}
	j, n := 0, len(*key)
	for i := 0; i < n; i++ {
		c := (*key)[i]
		_, ok := keyMap[c]
		if !ok {
			keyMap[c] = ALPHABET[j]
			j++
		}
	}
	return &keyMap
}

func decodeMessage(key string, message string) string {
	keyMap := GetKeyMap(&key)
	m := len(message)
	decodedMessage := make([]byte, m)
	for i := 0; i < m; i++ {
		c := message[i]
		decodedMessage[i] = (*keyMap)[c]
	}
	return string(decodedMessage)
}
