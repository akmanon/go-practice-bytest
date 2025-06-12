package concurrencyGo

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("comparing two urls, returning the fastest one", func(t *testing.T) {
		slowServer := delayServer(10 * time.Millisecond)
		fastServer := delayServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL
		got, err := Racer(slowURL, fastURL, time.Millisecond*10)
		if err != nil {
			t.Fatalf("Error not expected, but got %v", err)
		}
		if got != want {
			t.Errorf("got %q, want %q ", got, want)
		}
	})
	t.Run("return error if server doesn't respond within the specified time", func(t *testing.T) {
		slowServer := delayServer(15 * time.Millisecond)

		defer slowServer.Close()

		_, err := Racer(slowServer.URL, slowServer.URL, time.Millisecond*10)
		if err == nil {
			t.Fatalf("Error expected, but got none")
		}
	})

}

func delayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
