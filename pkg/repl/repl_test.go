package repl

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"Hello World", []string{"hello", "world"}},
		{"   Multiple   spaces    between   words   ", []string{"multiple", "spaces", "between", "words"}},
		{"MiXeD CaSe", []string{"mixed", "case"}},
		{"  Trim leading and trailing spaces  ", []string{"trim", "leading", "and", "trailing", "spaces"}},
		{"ALLUPPERCASE", []string{"alluppercase"}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := cleanInput(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("cleanInput(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}
