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

package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// Test for update-mtime command
func TestUpdateMTimeCommand(t *testing.T) {
	// Test that the command is properly configured
	if updateMTimeCmd.Use != "updatemtime" {
		t.Errorf("updateMTimeCmd.Use = %q, want %q", updateMTimeCmd.Use, "updatemtime")
	}
	
	if updateMTimeCmd.Short == "" {
		t.Error("updateMTimeCmd.Short should not be empty")
	}
	
	if updateMTimeCmd.Run == nil {
		t.Error("updateMTimeCmd.Run should not be nil")
	}
}

// Test the command registration
func TestUpdateMTimeCommandRegistration(t *testing.T) {
	// Check if the command has been added to rootCmd
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "updatemtime" {
			found = true
			break
		}
	}
	
	if !found {
		t.Error("updateMTimeCmd should be registered with rootCmd")
	}
}

func TestRenamerArgs(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		expectErr bool
		errMsg    string
	}{
		{
			name:      "valid single directory argument",
			args:      []string{"/path/to/directory"},
			expectErr: false,
		},
		{
			name:      "valid multiple arguments",
			args:      []string{"/path/to/directory", "extra-arg"},
			expectErr: false,
		},
		{
			name:      "no arguments",
			args:      []string{},
			expectErr: true,
			errMsg:    "requires a directory argument",
		},
		{
			name:      "empty argument list",
			args:      nil,
			expectErr: true,
			errMsg:    "requires a directory argument",
		},
	}

	cmd := &cobra.Command{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := renamerArgs(cmd, tt.args)
			if tt.expectErr {
				if err == nil {
					t.Errorf("renamerArgs() expected error but got none")
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("renamerArgs() error = %v, want error containing %q", err, tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("renamerArgs() unexpected error = %v", err)
				}
			}
		})
	}
}

func TestFindAvailablePathname(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := ioutil.TempDir("", "photo-renamer-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name               string
		pathname           string
		currentFilename    string
		formattedDatetime  string
		setupFiles         []string // files to create before test
		expectedResult     string
		expectErr          bool
		errMsg             string
	}{
		{
			name:              "simple rename - no conflicts",
			pathname:          filepath.Join(tempDir, "DSC_1234.jpg"),
			currentFilename:   "DSC_1234.jpg",
			formattedDatetime: "20210331103203",
			setupFiles:        []string{},
			expectedResult:    filepath.Join(tempDir, "20210331103203.jpg"),
			expectErr:         false,
		},
		{
			name:              "file already renamed - same name",
			pathname:          filepath.Join(tempDir, "20210331103203.jpg"),
			currentFilename:   "20210331103203.jpg",
			formattedDatetime: "20210331103203",
			setupFiles:        []string{},
			expectErr:         true,
			errMsg:            "File already renamed",
		},
		{
			name:              "conflict exists - use counter",
			pathname:          filepath.Join(tempDir, "DSC_5678.jpg"),
			currentFilename:   "DSC_5678.jpg",
			formattedDatetime: "20210331103203",
			setupFiles:        []string{"20210331103203.jpg"}, // create conflicting file
			expectedResult:    filepath.Join(tempDir, "20210331103203-1.jpg"),
			expectErr:         false,
		},
		{
			name:              "multiple conflicts - use higher counter",
			pathname:          filepath.Join(tempDir, "DSC_9999.jpg"),
			currentFilename:   "DSC_9999.jpg",
			formattedDatetime: "20210331103203",
			setupFiles:        []string{"20210331103203.jpg", "20210331103203-1.jpg", "20210331103203-2.jpg"},
			expectedResult:    filepath.Join(tempDir, "20210331103203-3.jpg"),
			expectErr:         false,
		},
		{
			name:              "too many conflicts - should fail",
			pathname:          filepath.Join(tempDir, "DSC_MANY.jpg"),
			currentFilename:   "DSC_MANY.jpg",
			formattedDatetime: "20210331103203",
			setupFiles:        []string{"20210331103203.jpg", "20210331103203-1.jpg", "20210331103203-2.jpg", "20210331103203-3.jpg", "20210331103203-4.jpg"},
			expectErr:         true,
			errMsg:            "Could not find available filename",
		},
		{
			name:              "different file extension",
			pathname:          filepath.Join(tempDir, "DSC_RAW.cr2"),
			currentFilename:   "DSC_RAW.cr2",
			formattedDatetime: "20210331103203",
			setupFiles:        []string{},
			expectedResult:    filepath.Join(tempDir, "20210331103203.cr2"),
			expectErr:         false,
		},
		{
			name:              "uppercase extension becomes lowercase",
			pathname:          filepath.Join(tempDir, "DSC_UP.JPG"),
			currentFilename:   "DSC_UP.JPG",
			formattedDatetime: "20210331103203",
			setupFiles:        []string{},
			expectedResult:    filepath.Join(tempDir, "20210331103203.jpg"),
			expectErr:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up any existing files from previous tests
			files, _ := filepath.Glob(filepath.Join(tempDir, "*"))
			for _, f := range files {
				os.Remove(f)
			}

			// Setup test files
			for _, filename := range tt.setupFiles {
				fullPath := filepath.Join(tempDir, filename)
				if err := ioutil.WriteFile(fullPath, []byte("dummy content"), 0644); err != nil {
					t.Fatalf("Failed to create test file %s: %v", fullPath, err)
				}
			}

			result, err := findAvailablePathname(tt.pathname, tt.currentFilename, tt.formattedDatetime)

			if tt.expectErr {
				if err == nil {
					t.Errorf("findAvailablePathname() expected error but got none")
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("findAvailablePathname() error = %v, want error containing %q", err, tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("findAvailablePathname() unexpected error = %v", err)
				} else if result != tt.expectedResult {
					t.Errorf("findAvailablePathname() = %q, want %q", result, tt.expectedResult)
				}
			}
		})
	}
}

// Test helper function to verify the logic of filename generation
func TestFindAvailablePathnameLogic(t *testing.T) {
	tests := []struct {
		name              string
		pathname          string
		currentFilename   string
		formattedDatetime string
		expectedBaseName  string
	}{
		{
			name:              "basic filename generation",
			pathname:          "/path/to/DSC_1234.jpg",
			currentFilename:   "DSC_1234.jpg",
			formattedDatetime: "20210331103203",
			expectedBaseName:  "20210331103203.jpg",
		},
		{
			name:              "preserve extension case conversion",
			pathname:          "/path/to/IMG_5678.JPG",
			currentFilename:   "IMG_5678.JPG",
			formattedDatetime: "20210401120000",
			expectedBaseName:  "20210401120000.jpg",
		},
		{
			name:              "RAW file extension",
			pathname:          "/path/to/DSC_9999.CR2",
			currentFilename:   "DSC_9999.CR2",
			formattedDatetime: "20210515160000",
			expectedBaseName:  "20210515160000.cr2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary directory
			tempDir, err := ioutil.TempDir("", "photo-renamer-logic-test")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tempDir)

			// Update pathname to use temp directory
			testPathname := filepath.Join(tempDir, tt.currentFilename)

			result, err := findAvailablePathname(testPathname, tt.currentFilename, tt.formattedDatetime)
			if err != nil {
				t.Errorf("findAvailablePathname() unexpected error = %v", err)
				return
			}

			// Extract the filename from the result
			resultFilename := filepath.Base(result)
			if resultFilename != tt.expectedBaseName {
				t.Errorf("findAvailablePathname() resulted in filename %q, want %q", resultFilename, tt.expectedBaseName)
			}

			// Verify the result is lowercase
			if strings.ToLower(result) != result {
				t.Errorf("findAvailablePathname() result %q should be lowercase", result)
			}
		})
	}
}