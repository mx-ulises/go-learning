type Codec struct {
	Memory map[string]string
}

func Constructor() Codec {
	return Codec{Memory: map[string]string{}}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	hash := fnv.New32a()
	hash.Write([]byte(longUrl))
	hashString := fmt.Sprintf("%06x", hash.Sum32())
	shortUrl := "http://tinyurl.com/" + hashString[len(hashString)-6:]
	this.Memory[shortUrl] = longUrl
	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.Memory[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
