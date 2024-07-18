package main

import "testing"

func TestCountCharOccurrences(t *testing.T) {
	result := CountCharOccurrences("hello", 'l')
	expected := 2
	if result != expected {
		t.Errorf("CountCharOccurrences(\"hello\", 'l') = %d; want %d", result, expected)
	}

	result = CountCharOccurrences("hello", 'x')
	expected = 0
	if result != expected {
		t.Errorf("CountCharOccurrences(\"hello\", 'x') = %d; want %d", result, expected)
	}
}
