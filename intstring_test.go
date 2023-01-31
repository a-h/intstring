package intstring

import (
	"encoding/json"
	"testing"
)

type Data struct {
	Value IntString `json:"value"`
}

var expectedJSON = `{"value":123}`

func TestMarshal(t *testing.T) {
	d := Data{
		Value: 123,
	}
	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("error marshaling value: %v", err)
	}
	if string(b) != expectedJSON {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedJSON, string(b))
	}
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Data
	}{
		{
			name:  "it can unmarshal a string to an integer",
			input: `{ "value": "123" }`,
		},
		{
			name:  "it can unmarshal an integer to an integer",
			input: `{ "value": 123 }`,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			var d Data
			err := json.Unmarshal([]byte(test.input), &d)
			if err != nil {
				t.Fatalf("error unmarshaling value: %v", err)
			}
			if d.Value != 123 {
				t.Errorf("expected value 123, got %v", d.Value)
			}
		})
	}
}
