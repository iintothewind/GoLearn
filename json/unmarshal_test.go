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

type Person struct {
	Name    string   `json:"Name"`
	Age     int      `json:"Age"`
	Parents []string `json:"Parents"`
}

func (p *Person) UnmarshalJSON(text []byte) error {
	var fields map[string]interface{}
	if err := json.Unmarshal(text, &fields); err != nil {
		return err
	}
	for key, value := range fields {
		switch key {
		case "Name":
			p.Name = value.(string)
		case "Age":
			p.Age = int(value.(float64))
		case "Parents":
			for _, parent := range value.([]interface{}) {
				p.Parents = append(p.Parents, parent.(string))
			}
		}
	}
	return nil

}

func TestUnmarshal04(t *testing.T) {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var p Person
	err := json.Unmarshal(b, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", p.Name)

}

func TestMarshal05(t *testing.T) {
	input := `{
  "default": {
    "values": [
      {
        "fieldType": "text",
        "label": ""
      }
    ]
  },
  "displayName": "Form Fields",
  "name": "formFields",
  "options": [
    {
      "displayName": "Values",
      "name": "values",
      "values": [
        {
          "default": "",
          "description": "Label appears above the input field",
          "displayName": "Field Label",
          "name": "fieldLabel",
          "placeholder": "e.g. What is your name?",
          "required": true,
          "type": "string"
        },
        {
          "default": "text",
          "description": "The type of field to add to the form",
          "displayName": "Field Type",
          "name": "fieldType",
          "options": [
            {
              "name": "Date",
              "value": "date"
            },
            {
              "name": "Dropdown List",
              "value": "dropdown"
            },
            {
              "name": "Number",
              "value": "number"
            },
            {
              "name": "Password",
              "value": "password"
            },
            {
              "name": "Text",
              "value": "text"
            },
            {
              "name": "Textarea",
              "value": "textarea"
            }
          ],
          "required": true,
          "type": "options"
        },
        {
          "default": {
            "values": [
              {
                "option": ""
              }
            ]
          },
          "description": "List of options that can be selected from the dropdown",
          "displayName": "Field Options",
          "displayOptions": {
            "show": {
              "fieldType": [
                "dropdown"
              ]
            }
          },
          "name": "fieldOptions",
          "options": [
            {
              "displayName": "Values",
              "name": "values",
              "values": [
                {
                  "default": "",
                  "displayName": "Option",
                  "name": "option",
                  "type": "string"
                }
              ]
            }
          ],
          "placeholder": "Add Field Option",
          "required": true,
          "type": "fixedCollection",
          "typeOptions": {
            "multipleValues": true,
            "sortable": true
          }
        },
        {
          "default": false,
          "description": "Whether to allow the user to select multiple options from the dropdown list",
          "displayName": "Multiple Choice",
          "displayOptions": {
            "show": {
              "fieldType": [
                "dropdown"
              ]
            }
          },
          "name": "multiselect",
          "type": "boolean"
        },
        {
          "default": false,
          "description": "Whether to require the user to enter a value for this field before submitting the form",
          "displayName": "Required Field",
          "name": "requiredField",
          "type": "boolean"
        }
      ]
    }
  ],
  "placeholder": "Add Form Field",
  "type": "fixedCollection",
  "typeOptions": {
    "multipleValues": true,
    "sortable": true
  }
}`

	var fields map[string]interface{}
	fields = make(map[string]interface{})

	if err := json.Unmarshal([]byte(input), &fields); err != nil {
		fmt.Errorf("error unmarshalling input: %v", err)
	}

	fmt.Println(fields)

}
