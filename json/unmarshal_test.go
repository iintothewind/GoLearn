package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"
)

type Size int

const (
	Unrecognized Size = iota
	Small
	Large
)

func (s *Size) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	default:
		*s = Unrecognized
	case "small":
		*s = Small
	case "large":
		*s = Large
	}
	return nil
}

func (s Size) MarshalText() ([]byte, error) {
	var name string
	switch s {
	default:
		name = "unrecognized"
	case Small:
		name = "small"
	case Large:
		name = "large"
	}
	return []byte(name), nil
}

func TestUnMarshal01(t *testing.T) {
	blob := `["small","regular","large","unrecognized","small","normal","small","large"]`
	var inventory []Size
	if err := json.Unmarshal([]byte(blob), &inventory); err != nil {
		log.Fatal(err)
	}

	counts := make(map[Size]int)
	for _, size := range inventory {
		counts[size] += 1
	}

	fmt.Printf("Inventory Counts:\n* Small:        %d\n* Large:        %d\n* Unrecognized: %d\n",
		counts[Small], counts[Large], counts[Unrecognized])

}

type Message struct {
	Name string
	Body string
	Time int64
}

func TestMarshal02(t *testing.T) {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}

func TestUnmarshal03(t *testing.T) {
	text := `{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`
	var m Message
	err := json.Unmarshal([]byte(text), &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", m)
}

func TestUnmarshal04(t *testing.T) {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", f.(map[string]interface{})["Name"])

}
