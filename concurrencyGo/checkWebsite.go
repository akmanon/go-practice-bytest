package concurrencyGo

type WebsiteChecker func(string) bool

type result struct {
	url    string
	status bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)
	for _, url := range urls {
		go func() {
			resultsChannel <- result{url, wc(url)}
		}()
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.url] = r.status
	}
	return results
}
