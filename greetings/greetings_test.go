package greetings

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestStrings01(t *testing.T) {
	for _, s := range strings.Fields("a b c") {
		fmt.Println(s)
	}
	a, b := 0, 1
	println(a, b)
	type CustomType float64

	f := 1.0
	c := CustomType(f)

	println(c)

	type Person struct {
		Name string
		Age  int
	}

	type Employee struct {
		Name string
		Role string
	}

	p := Person{"John", 30}

	println(p.Name, p.Age)

}
