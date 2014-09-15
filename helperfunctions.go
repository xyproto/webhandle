package webhandle

import (
	"errors"
	"fmt"
	"html"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

const (
	USERNAME_ALLOWED_LETTERS = "abcdefghijklmnopqrstuvwxyzæøåABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ_0123456789"
)

// Returns a form parameter, or an empty string
func GetFormParam(req *http.Request, key string) string {
	return req.PostFormValue(key)
}

// Returns an url parameter, or an empty string
func GetParam(req *http.Request, key string) string {
	values, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return ""
	}
	val := values.Get(key)
	return val
}

// Write the given string to the response writer. Convenience function.
func Ret(w http.ResponseWriter, s string) {
	fmt.Fprint(w, s)
}

// Get a value from an url.
// /hi/there/asdf with pos 2 returns asdf
func GetVal(url *url.URL, pos int) string {
	p := html.EscapeString(url.Path)
	fields := strings.Split(p, "/")
	if len(fields) <= pos {
		return ""
	}
	return fields[pos]
}

// Get the last value from an url.
// /hi/there/qwerty returns qwerty
func GetLast(url *url.URL) string {
	p := html.EscapeString(url.Path)
	fields := strings.Split(p, "/")
	if len(fields) == 0 {
		return ""
	}
	return fields[len(fields)-1]
}

// Converts "true" or "false" to a bool
func TruthValue(val string) bool {
	return "true" == val
}

// Split a string into two strings at the colon
// If there's no colon, return the string and an empty string
func HostPortSplit(s string) (string, string) {
	if strings.Contains(s, ":") {
		sl := strings.SplitN(s, ":", 2)
		return sl[0], sl[1]
	}
	return s, ""
}

func TableCell(b bool) string {
	if b {
		return "<td class=\"yes\">yes</td>"
	}
	return "<td class=\"no\">no</td>"
}

func RandomString(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = byte(rand.Int63() & 0xff)
	}
	return string(b)
}

func RandomHumanFriendlyString(length int) string {
	const (
		vowels     = "aeiouy" // email+browsers didn't like "æøå" too much
		consonants = "bcdfghjklmnpqrstvwxz"
	)
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			b[i] = vowels[rand.Intn(len(vowels))]
		} else {
			b[i] = consonants[rand.Intn(len(consonants))]
		}
	}
	return string(b)
}

func RandomCookieFriendlyString(length int) string {
	const ALLOWED = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = ALLOWED[rand.Intn(len(ALLOWED))]
	}
	return string(b)
}

func CleanUserInput(val string) string {
	return strings.Replace(val, "<", "&lt;", -1)
}

func Check(username, password string) error {
NEXT:
	for _, letter := range username {
		for _, allowedLetter := range USERNAME_ALLOWED_LETTERS {
			if letter == allowedLetter {
				continue NEXT
			}
		}
		return errors.New("Only a-å, A-Å, 0-9 and _ are allowed in usernames.")
	}
	if username == password {
		return errors.New("Username and password must be different, try another password.")
	}
	return nil
}