package concurrencyGo

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing counter to 3, sequentially", func(t *testing.T) {
		counter := Counter{}
		want := 3
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(&counter, want, t)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 10000
		counter := Counter{}
		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for range wantedCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(&counter, wantedCount, t)
	})

}

func assertCounter(got *Counter, want int, t testing.TB) {
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
