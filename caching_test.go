package webhandle

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFakeHandler(t *testing.T) {
	const msg = "hi"

	fake := NewFakeResponseWriter()
	fmt.Fprint(fake, msg)
	if fake.String() != msg {
		t.Errorf("Fake response writer should return \"%s\", but returned: %s\n", msg, fake.String())
	}

	fake2 := &FakeResponseWriter{}
	fmt.Fprint(fake2, msg)
	if fake2.String() != msg {
		t.Errorf("Fake response writer should return \"%s\", but returned: %s\n", msg, fake2.String())
	}
}

var counter int

// Return a string and increase the counter
func generateTextAndCount() string {
	counter++
	return "generated"
}

// Testing NewCacheWrapper
func TestCaching(t *testing.T) {

	generateTextAndCount()
	generateTextAndCount()
	generateTextAndCount()

	if counter != 3 {
		t.Errorf("GenerateTextAndCount failed")
	}

	handler := NewCacheWrapper("abc123", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, generateTextAndCount())
	})
	fake := NewFakeResponseWriter()
	handler(fake, nil)
	handler(fake, nil)
	handler(fake, nil)

	if fake.String() == "" {
		t.Errorf("No text in cache!")
	}
	if counter != 4 {
		t.Errorf("Requests are not cached, but generated for each call!")
	}
}
