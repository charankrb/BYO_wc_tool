package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestWC(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected struct {
			lines int
			words int
			chars int
		}
	}{
		{
			name:  "Single line",
			input: "Hello world",
			expected: struct {
				lines int
				words int
				chars int
			}{lines: 1, words: 2, chars: 12}, // 12 includes the newline character
		},
		{
			name:  "Multiple lines",
			input: "Hello world\nHello Go\nHello test",
			expected: struct {
				lines int
				words int
				chars int
			}{lines: 3, words: 6, chars: 32}, // 32 includes newline characters
		},
		{
			name:  "Empty input",
			input: "",
			expected: struct {
				lines int
				words int
				chars int
			}{lines: 0, words: 0, chars: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			lines, words, chars := wc(scanner)

			if lines != tt.expected.lines {
				t.Errorf("Expected lines: %d, got: %d", tt.expected.lines, lines)
			}
			if words != tt.expected.words {
				t.Errorf("Expected words: %d, got: %d", tt.expected.words, words)
			}
			if chars != tt.expected.chars {
				t.Errorf("Expected chars: %d, got: %d", tt.expected.chars, chars)
			}
		})
	}
}
