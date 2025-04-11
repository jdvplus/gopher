package greetings

import (
	"regexp"
	"testing"
)

// test greetings.Hello with a valid name
func TestHelloName(t *testing.T) {
	name := "jun-boy"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("jun-boy")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("jun-boy") = %q, %v, want match for %q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}