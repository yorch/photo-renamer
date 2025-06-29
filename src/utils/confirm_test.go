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

import (
	"strings"
	"testing"
)

func TestParseConfirmationResponse(t *testing.T) {
	tests := []struct {
		name         string
		response     string
		expected     bool
		expectErr    bool
		errContains  string
	}{
		// Valid "yes" responses
		{
			name:      "lowercase y",
			response:  "y",
			expected:  true,
			expectErr: false,
		},
		{
			name:      "uppercase Y",
			response:  "Y",
			expected:  true,
			expectErr: false,
		},
		{
			name:      "lowercase yes",
			response:  "yes",
			expected:  true,
			expectErr: false,
		},
		{
			name:      "capitalized Yes",
			response:  "Yes",
			expected:  true,
			expectErr: false,
		},
		{
			name:      "uppercase YES",
			response:  "YES",
			expected:  true,
			expectErr: false,
		},
		// Valid "no" responses
		{
			name:      "lowercase n",
			response:  "n",
			expected:  false,
			expectErr: false,
		},
		{
			name:      "uppercase N",
			response:  "N",
			expected:  false,
			expectErr: false,
		},
		{
			name:      "lowercase no",
			response:  "no",
			expected:  false,
			expectErr: false,
		},
		{
			name:      "capitalized No",
			response:  "No",
			expected:  false,
			expectErr: false,
		},
		{
			name:      "uppercase NO",
			response:  "NO",
			expected:  false,
			expectErr: false,
		},
		// Invalid responses
		{
			name:        "invalid response - maybe",
			response:    "maybe",
			expected:    false,
			expectErr:   true,
			errContains: "invalid response",
		},
		{
			name:        "invalid response - empty string",
			response:    "",
			expected:    false,
			expectErr:   true,
			errContains: "invalid response",
		},
		{
			name:        "invalid response - yep",
			response:    "yep",
			expected:    false,
			expectErr:   true,
			errContains: "invalid response",
		},
		{
			name:        "invalid response - nope",
			response:    "nope",
			expected:    false,
			expectErr:   true,
			errContains: "invalid response",
		},
		{
			name:        "invalid response - whitespace",
			response:    " ",
			expected:    false,
			expectErr:   true,
			errContains: "invalid response",
		},
		{
			name:        "invalid response - number",
			response:    "1",
			expected:    false,
			expectErr:   true,
			errContains: "invalid response",
		},
		// Case variations that should not work
		{
			name:        "mixed case - yEs (not in valid list)",
			response:    "yEs",
			expected:    false,
			expectErr:   true,
			errContains: "invalid response",
		},
		{
			name:        "mixed case - nO (not in valid list)",
			response:    "nO",
			expected:    false,
			expectErr:   true,
			errContains: "invalid response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseConfirmationResponse(tt.response)
			
			if tt.expectErr {
				if err == nil {
					t.Errorf("ParseConfirmationResponse(%q) expected error but got none", tt.response)
				} else if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("ParseConfirmationResponse(%q) error = %v, want error containing %q", tt.response, err, tt.errContains)
				}
			} else {
				if err != nil {
					t.Errorf("ParseConfirmationResponse(%q) unexpected error = %v", tt.response, err)
				} else if result != tt.expected {
					t.Errorf("ParseConfirmationResponse(%q) = %t, want %t", tt.response, result, tt.expected)
				}
			}
		})
	}
}

// Test to verify the exact list of valid responses
func TestValidResponseLists(t *testing.T) {
	// Test all expected "yes" responses
	yesResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	for _, response := range yesResponses {
		t.Run("yes_response_"+response, func(t *testing.T) {
			result, err := ParseConfirmationResponse(response)
			if err != nil {
				t.Errorf("ParseConfirmationResponse(%q) unexpected error = %v", response, err)
			}
			if !result {
				t.Errorf("ParseConfirmationResponse(%q) = false, want true", response)
			}
		})
	}

	// Test all expected "no" responses
	noResponses := []string{"n", "N", "no", "No", "NO"}
	for _, response := range noResponses {
		t.Run("no_response_"+response, func(t *testing.T) {
			result, err := ParseConfirmationResponse(response)
			if err != nil {
				t.Errorf("ParseConfirmationResponse(%q) unexpected error = %v", response, err)
			}
			if result {
				t.Errorf("ParseConfirmationResponse(%q) = true, want false", response)
			}
		})
	}
}