func getCountAndDomain(cpdomain string) (int, string) {
	for i := 0; i < len(cpdomain); i++ {
		if cpdomain[i] == ' ' {
			count, _ := strconv.Atoi(cpdomain[:i])
			domain := cpdomain[i+1:]
			return count, domain
		}
	}
	return -1, ""
}

func addSubdomains(subdomainCount map[string]int, count int, domain string) {
	subdomainCount[domain] += count
	for i := 0; i < len(domain); i++ {
		if domain[i] == '.' {
			subdomainCount[domain[i+1:]] += count
		}
	}
}

func buildOutput(subdomainCount map[string]int) []string {
	output := make([]string, len(subdomainCount))
	i := 0
	for domain, count := range subdomainCount {
		output[i] = strconv.Itoa(count) + " " + domain
		i++
	}
	return output
}

func subdomainVisits(cpdomains []string) []string {
	subdomainCount := map[string]int{}
	for _, cpdomain := range cpdomains {
		count, domain := getCountAndDomain(cpdomain)
		addSubdomains(subdomainCount, count, domain)
	}
	return buildOutput(subdomainCount)
}
