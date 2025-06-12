package concurrencyGo

import (
	"reflect"
	"testing"
	"time"
)

func mockCheckWebsite(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}
	got := CheckWebsite(mockCheckWebsite, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, gott %v ", got, want)
	}

}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(time.Millisecond * 20)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsite(slowStubWebsiteChecker, urls)
	}
}
