package concurrencyGo

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(time.Millisecond * 100)
	return s.response
}
func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "Hello, world"
	srv := Server(&SpyStore{data, false})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	srv.ServeHTTP(res, req)
	if res.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, res.Body.String(), data)
	}
}
