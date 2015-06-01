package ctx

import (
	"net/http"
	"sync"
)

// NOTE: no mutex.
// Not sure if an RWMutex is needed here.
var contexts = make(map[*http.Request]context)

type context struct {
	mu sync.RWMutex
	d  map[string]interface{}
}

func (it *context) Get(key string) (val interface{}, ok bool) {
	it.mu.RLock()
	defer it.mu.RUnlock()
	val, ok = it.d[key]
	return
}

func (it *context) Set(key string, val interface{}) {
	it.mu.Lock()
	defer it.mu.Unlock()
	it.d[key] = val
}

func For(r *http.Request) context {
	c, _ := contexts[r]
	return c
}

func create(r *http.Request) {
	contexts[r] = context{d: make(map[string]interface{})}
}

func remove(r *http.Request) {
	delete(contexts, r)
}

var Handler = func(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		create(r)
		defer remove(r)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
