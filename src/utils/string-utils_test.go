// Copyright © 2021 Jorge Barnaby <jorge.barnaby@gmail.com>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package utils

import "testing"

func TestPosString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		element  string
		expected int
	}{
		{
			name:     "element found at beginning",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "apple",
			expected: 0,
		},
		{
			name:     "element found in middle",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "banana",
			expected: 1,
		},
		{
			name:     "element found at end",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "cherry",
			expected: 2,
		},
		{
			name:     "element not found",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "orange",
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			element:  "apple",
			expected: -1,
		},
		{
			name:     "empty element in slice",
			slice:    []string{"", "banana", "cherry"},
			element:  "",
			expected: 0,
		},
		{
			name:     "case sensitive",
			slice:    []string{"Apple", "banana", "cherry"},
			element:  "apple",
			expected: -1,
		},
		{
			name:     "duplicate elements - returns first occurrence",
			slice:    []string{"apple", "banana", "apple"},
			element:  "apple",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PosString(tt.slice, tt.element)
			if result != tt.expected {
				t.Errorf("PosString(%v, %q) = %d; want %d", tt.slice, tt.element, result, tt.expected)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		element  string
		expected bool
	}{
		{
			name:     "element found",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "banana",
			expected: true,
		},
		{
			name:     "element not found",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "orange",
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			element:  "apple",
			expected: false,
		},
		{
			name:     "empty element found",
			slice:    []string{"", "banana", "cherry"},
			element:  "",
			expected: true,
		},
		{
			name:     "empty element not found",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "",
			expected: false,
		},
		{
			name:     "case sensitive - different case",
			slice:    []string{"Apple", "banana", "cherry"},
			element:  "apple",
			expected: false,
		},
		{
			name:     "case sensitive - exact match",
			slice:    []string{"Apple", "banana", "cherry"},
			element:  "Apple",
			expected: true,
		},
		{
			name:     "duplicate elements",
			slice:    []string{"apple", "banana", "apple"},
			element:  "apple",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsString(tt.slice, tt.element)
			if result != tt.expected {
				t.Errorf("ContainsString(%v, %q) = %t; want %t", tt.slice, tt.element, result, tt.expected)
			}
		})
	}
}