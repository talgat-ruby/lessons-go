package greet

import (
	"testing"
)

func TestGreet(t *testing.T) {
	expected := "Hello Leila! Your table is 10."
	text := Greet("Leila", 10)

	if text != expected {
		t.Errorf("Greet(%s, %d), got: %s, expected: %s", "Leila", 10, text, expected)
	}
}

func TestTableGreet(t *testing.T) {
	table := []struct {
		name     string
		table    int
		expected string
	}{
		{"Leila", 5, "Hello Leila! Your table is 5."},
		{"Zoidberg", -1, "Hello Zoidberg! Your table is -1."},
		{"Prf. Farnsworth", 10, "Hello Prf. Farnsworth! Your table is 10."},
	}

	for _, row := range table {
		text := Greet(row.name, row.table)
		if text != row.expected {
			t.Errorf("Greet(%s, %d) was incorrect, got: %s, expected: %s.", row.name, row.table, text, row.expected)
		}
	}
}
