package webhandle

import (
	"bytes"
	"fmt"
	"net/http"
)

var globalStringCache map[string]string

type FakeResponseWriter struct {
	buf bytes.Buffer
}

func (f FakeResponseWriter) Header() http.Header {
	return http.Header{}
}

func (f FakeResponseWriter) Write(b []byte) (int, error) {
	return f.buf.Write(b)
}

func (f FakeResponseWriter) WriteHeader(i int) {
}

func (f FakeResponseWriter) String() string {
	return f.buf.String()
}

// Wrap a http.HandlerFunc so that the output is cached (with an id)
// Do not cache functions with side-effects! (that sets the mimetype for instance)
// The safest thing for now is to only cache images.
func NewCacheWrapper(id string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if _, ok := globalStringCache[id]; !ok {
			fake := &FakeResponseWriter{}
			// Make the handler write to the buffer
			handler(fake, req)
			globalStringCache[id] = fake.String()
		}
		fmt.Fprint(w, globalStringCache[id])
	}
}
