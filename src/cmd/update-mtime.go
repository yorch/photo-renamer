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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/spf13/cobra"
)

var mtimeCounter int = 0

func updateMTimeDirectory(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			pathname := filepath.Join(directory, file.Name())
			updateMTimeForFile(pathname, file, nil)
		}
	}
}

func updateMTimeForFile(pathname string, f os.FileInfo, err error) error {
	openedFile, _ := os.Open(pathname)
	exifData, err := exif.Decode(openedFile)

	if err != nil {
		fmt.Printf("x Unable to load EXIF data for file: %s\n", pathname)
		return nil
	}

	filename := f.Name()
	exifDatetime, _ := exifData.DateTime()

	if err := os.Chtimes(pathname, exifDatetime, exifDatetime); err != nil {
		fmt.Printf("%s\n", err)
		return nil
	}

	fmt.Printf("✓ Updated file %s mtime to %v\n", filename, exifDatetime)
	mtimeCounter++

	return nil
}

func updateMTime(cmd *cobra.Command, args []string) {
	directory := args[0]
	recursiveMode, _ := cmd.Flags().GetBool("recursive")

	startTime := time.Now()
	
	if recursiveMode {
		filepath.Walk(directory, updateMTimeForFile)
	} else {
		updateMTimeDirectory(directory)
	}
	
	endTime := time.Now()

	fmt.Printf("\n\nʕ◔ϖ◔ʔ I successfully updated %d photos in %s\n", mtimeCounter, endTime.Sub(startTime))
}

var updateMTimeCmd = &cobra.Command{
	Use:   "updatemtime",
	Short: "Updates mtime of image files based on their EXIF info",
	Long:  "",
	Run:   updateMTime,
}

func init() {
	rootCmd.AddCommand(updateMTimeCmd)
	updateMTimeCmd.Flags().BoolP("recursive", "r", false, "recursively process subdirectories")
}
