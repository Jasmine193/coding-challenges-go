package wc

import "testing"

func TestCountBytes(t *testing.T) {
	data := []byte("Hello, World!")
	expected := 13
	result := CountBytes(data)

	if result != expected {
		t.Errorf("CountBytes failed: expected %d, got %d", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	data := []byte("Hello\nworld\n\n\nThis is a test.")
	expected := 4
	result := CountLines(data)
	if result != expected {
		t.Errorf("CountLines failed: expected %d, got %d", expected, result)
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "multiple words",
			input: []byte("Hello, World!"),
			want:  2,
		},
		{
			name:  "multiple spaces",
			input: []byte("Hello     World"),
			want:  2,
		},
		{
			name:  "new lines",
			input: []byte("Hello\nWorld\nGo"),
			want:  3,
		},
		{
			name:  "tabs",
			input: []byte("Hello\tWorld\tGo"),
			want:  3,
		},
		{
			name:  "leading whitespace",
			input: []byte("   Hello World"),
			want:  2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountWords(test.input)
			if result != test.want {
				t.Errorf("CountWords failed for %s: expected %d, got %d", test.name, test.want, result)
			}
		})
	}
}

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "ASCII characters",
			input: []byte("Hello, World!"),
			want:  13,
		},
		{
			name:  "Unicode characters",
			input: []byte("こんにちは世界"), // "Hello, World" in Japanese
			want:  7,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountCharacters(test.input)
			if result != test.want {
				t.Errorf("CountCharacters failed for %s: expected %d, got %d", test.name, test.want, result)
			}
		})
	}
}
