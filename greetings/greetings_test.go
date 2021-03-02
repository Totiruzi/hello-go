package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings. Hello with a name checking for a possible return value
func TestHelloName(t *testing.T) {
    name := "Peter"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Peter")

    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Peter") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string, check for an error,
//  if any it will indicate that our error handling method works
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")

    if msg !="" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
