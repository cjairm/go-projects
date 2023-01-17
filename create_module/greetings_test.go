package greetings

import(
  "testing"
  "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value
func TestHelloName(t *testing.T) {
  name := "Carlos"
  expected := regexp.MustCompile(`\b`+name+`\b`)
  message, err := Hello("Carlos")
  if !expected.MatchString(message) || err != nil {
    t.Fatalf(`Hello("Carlos") = %q, %v, expected match for %#q, nil`, message, err, expected)
  }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error
func TestHelloEmpty(t *testing.T) {
  message, err := Hello("")
  if message != "" || err == nil {
    t.Fatalf(`Hello("") = %q, %v, expected "", error`, message, err)
  }
}
