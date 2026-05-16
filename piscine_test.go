package piscine

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// captureOutput captures the output from a function that prints to stdout
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// TestQuadA tests the QuadA function with various inputs
func TestQuadA(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected string
	}{
		{
			name:     "5x3 rectangle",
			x:        5,
			y:        3,
			expected: "o---o\n|   |\no---o\n",
		},
		{
			name:     "5x1 line",
			x:        5,
			y:        1,
			expected: "o---o\n",
		},
		{
			name:     "1x1 single point",
			x:        1,
			y:        1,
			expected: "o\n",
		},
		{
			name:     "1x5 vertical line",
			x:        1,
			y:        5,
			expected: "o\n|\n|\n|\no\n",
		},
		{
			name:     "0x0 empty",
			x:        0,
			y:        0,
			expected: "",
		},
		{
			name:     "negative x",
			x:        -1,
			y:        6,
			expected: "",
		},
		{
			name:     "negative y",
			x:        6,
			y:        -1,
			expected: "",
		},
		{
			name:     "20x1 wide line",
			x:        20,
			y:        1,
			expected: "o------------------o\n",
		},
		{
			name:     "10x8 rectangle",
			x:        10,
			y:        8,
			expected: "o--------o\n|        |\n|        |\n|        |\n|        |\n|        |\n|        |\no--------o\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				QuadA(tt.x, tt.y)
			})
			t.Logf("\nQuadA - Input: x=%d, y=%d\n--- Expected ---\n%s\n--- Got ---\n%s\n----------------", tt.x, tt.y, tt.expected, output)
			if output != tt.expected {
				t.Errorf("QuadA(%d, %d) = %q, want %q", tt.x, tt.y, output, tt.expected)
			}
		})
	}
}

// TestQuadB tests the QuadB function with various inputs
func TestQuadB(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected string
	}{
		{
			name:     "5x3 rectangle",
			x:        5,
			y:        3,
			expected: "/***\\\n*   *\n\\***/\n",
		},
		{
			name:     "5x1 line",
			x:        5,
			y:        1,
			expected: "/***\\\n",
		},
		{
			name:     "1x1 single point",
			x:        1,
			y:        1,
			expected: "/\n",
		},
		{
			name:     "1x5 vertical line",
			x:        1,
			y:        5,
			expected: "/\n*\n*\n*\n\\\n",
		},
		{
			name:     "0x0 empty",
			x:        0,
			y:        0,
			expected: "",
		},
		{
			name:     "negative x",
			x:        -1,
			y:        6,
			expected: "",
		},
		{
			name:     "negative y",
			x:        6,
			y:        -1,
			expected: "",
		},
		{
			name:     "18x6 rectangle",
			x:        18,
			y:        6,
			expected: "/****************\\\n*                *\n*                *\n*                *\n*                *\n\\****************/\n",
		},
		{
			name:     "9x3 rectangle",
			x:        9,
			y:        3,
			expected: "/*******\\\n*       *\n\\*******/\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				QuadB(tt.x, tt.y)
			})
			t.Logf("\nQuadB - Input: x=%d, y=%d\n--- Expected ---\n%s\n--- Got ---\n%s\n----------------", tt.x, tt.y, tt.expected, output)
			if output != tt.expected {
				t.Errorf("QuadB(%d, %d) = %q, want %q", tt.x, tt.y, output, tt.expected)
			}
		})
	}
}

// TestQuadC tests the QuadC function with various inputs
func TestQuadC(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected string
	}{
		{
			name:     "5x3 rectangle",
			x:        5,
			y:        3,
			expected: "ABBBA\nB   B\nCBBBC\n",
		},
		{
			name:     "5x1 line",
			x:        5,
			y:        1,
			expected: "ABBBA\n",
		},
		{
			name:     "1x1 single point",
			x:        1,
			y:        1,
			expected: "A\n",
		},
		{
			name:     "1x5 vertical line",
			x:        1,
			y:        5,
			expected: "A\nB\nB\nB\nC\n",
		},
		{
			name:     "0x0 empty",
			x:        0,
			y:        0,
			expected: "",
		},
		{
			name:     "negative x",
			x:        -1,
			y:        6,
			expected: "",
		},
		{
			name:     "negative y",
			x:        6,
			y:        -1,
			expected: "",
		},
		{
			name:     "13x7 rectangle",
			x:        13,
			y:        7,
			expected: "ABBBBBBBBBBBA\nB           B\nB           B\nB           B\nB           B\nB           B\nCBBBBBBBBBBBC\n",
		},
		{
			name:     "10x15 rectangle",
			x:        10,
			y:        15,
			expected: "ABBBBBBBBA\nB        B\nB        B\nB        B\nB        B\nB        B\nB        B\nB        B\nB        B\nB        B\nB        B\nB        B\nB        B\nB        B\nCBBBBBBBBC\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				QuadC(tt.x, tt.y)
			})
			t.Logf("\nQuadC - Input: x=%d, y=%d\n--- Expected ---\n%s\n--- Got ---\n%s\n----------------", tt.x, tt.y, tt.expected, output)
			if output != tt.expected {
				t.Errorf("QuadC(%d, %d) = %q, want %q", tt.x, tt.y, output, tt.expected)
			}
		})
	}
}

// TestQuadD tests the QuadD function with various inputs
func TestQuadD(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected string
	}{
		{
			name:     "5x3 rectangle",
			x:        5,
			y:        3,
			expected: "ABBBC\nB   B\nABBBC\n",
		},
		{
			name:     "5x1 line",
			x:        5,
			y:        1,
			expected: "ABBBC\n",
		},
		{
			name:     "1x1 single point",
			x:        1,
			y:        1,
			expected: "A\n",
		},
		{
			name:     "1x5 vertical line",
			x:        1,
			y:        5,
			expected: "A\nB\nB\nB\nA\n",
		},
		{
			name:     "0x0 empty",
			x:        0,
			y:        0,
			expected: "",
		},
		{
			name:     "negative x",
			x:        -1,
			y:        6,
			expected: "",
		},
		{
			name:     "negative y",
			x:        6,
			y:        -1,
			expected: "",
		},
		{
			name:     "3x16 rectangle",
			x:        3,
			y:        16,
			expected: "ABC\nB B\nB B\nB B\nB B\nB B\nB B\nB B\nB B\nB B\nB B\nB B\nB B\nB B\nB B\nABC\n",
		},
		{
			name:     "7x16 rectangle",
			x:        7,
			y:        16,
			expected: "ABBBBBC\nB     B\nB     B\nB     B\nB     B\nB     B\nB     B\nB     B\nB     B\nB     B\nB     B\nB     B\nB     B\nB     B\nB     B\nABBBBBC\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				QuadD(tt.x, tt.y)
			})
			t.Logf("\nQuadD - Input: x=%d, y=%d\n--- Expected ---\n%s\n--- Got ---\n%s\n----------------", tt.x, tt.y, tt.expected, output)
			if output != tt.expected {
				t.Errorf("QuadD(%d, %d) = %q, want %q", tt.x, tt.y, output, tt.expected)
			}
		})
	}
}

// TestQuadE tests the QuadE function with various inputs
func TestQuadE(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected string
	}{
		{
			name:     "5x3 rectangle",
			x:        5,
			y:        3,
			expected: "ABBBC\nB   B\nCBBBA\n",
		},
		{
			name:     "5x1 line",
			x:        5,
			y:        1,
			expected: "ABBBC\n",
		},
		{
			name:     "1x1 single point",
			x:        1,
			y:        1,
			expected: "A\n",
		},
		{
			name:     "1x5 vertical line",
			x:        1,
			y:        5,
			expected: "A\nB\nB\nB\nC\n",
		},
		{
			name:     "0x0 empty",
			x:        0,
			y:        0,
			expected: "",
		},
		{
			name:     "negative x",
			x:        -1,
			y:        6,
			expected: "",
		},
		{
			name:     "negative y",
			x:        6,
			y:        -1,
			expected: "",
		},
		{
			name:     "21x24 rectangle",
			x:        21,
			y:        24,
			expected: "ABBBBBBBBBBBBBBBBBBBC\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nB                   B\nCBBBBBBBBBBBBBBBBBBBA\n",
		},
		{
			name:     "18x8 rectangle",
			x:        18,
			y:        8,
			expected: "ABBBBBBBBBBBBBBBBC\nB                B\nB                B\nB                B\nB                B\nB                B\nB                B\nCBBBBBBBBBBBBBBBBA\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				QuadE(tt.x, tt.y)
			})
			t.Logf("\nQuadE - Input: x=%d, y=%d\n--- Expected ---\n%s\n--- Got ---\n%s\n----------------", tt.x, tt.y, tt.expected, output)
			if output != tt.expected {
				t.Errorf("QuadE(%d, %d) = %q, want %q", tt.x, tt.y, output, tt.expected)
			}
		})
	}
}
