package benchmark

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"../data"
	"../handlers"
)

func BenchmarkSearchHandler(b *testing.B) {
	mockStore := &data.MockStore{}
	mockStore.On("Search", "Fat Freddy's Cat").Return([]data.Kitten{
		data.Kitten{
			Name: "Fat Freddy's Cat",
		},
	})
	search := handlers.Search{DataStore: mockStore}
	for i := 0; i < b.N; i++ {
		r := httptest.NewRequest("POST", "/search",
			bytes.NewReader([]byte(`{"query":"Fat Freddy's Cat"}`)))
		rr := httptest.NewRecorder()
		search.ServeHTTP(rr, r)
	}
}
