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
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/spf13/cobra"
	"github.com/yorch/photo-renamer/utils"
)

var counter int = 0

func visit(pathname string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return nil
	}

	if info == nil {
		fmt.Printf("There was an error trying to read the given file / dir\n")
		return nil
	}

	if !info.IsDir() {
		openedFile, _ := os.Open(pathname)
		exifData, err := exif.Decode(openedFile)

		if err != nil {
			fmt.Printf("x Unable to load EXIF data for file: %s\n", pathname)
			return nil
		}

		currentFilename := info.Name()
		exifDatetime, _ := exifData.DateTime()
		formattedDatetime := exifDatetime.Format("20060102150405")
		newPathname, err := findAvailablePathname(pathname, currentFilename, formattedDatetime)

		if err != nil {
			fmt.Printf("%s\n", err)
			return nil
		}

		os.Rename(pathname, newPathname)
		fmt.Printf("✓ Renamed to: %s\n", newPathname)
		counter++
	}

	return nil
}

func findAvailablePathname(pathname string, currentFilename string, formattedDatetime string) (string, error) {
	count := 0
	for count < 5 {
		newFilename := fmt.Sprintf("%s%s", formattedDatetime, path.Ext(pathname))
		if currentFilename == newFilename {
			return "", errors.New("! File already renamed: " + pathname)
		}
		if count > 0 {
			newFilename = fmt.Sprintf("%s-%v%s", formattedDatetime, count, path.Ext(pathname))
		}
		newPathname := strings.ToLower(strings.Replace(pathname, currentFilename, newFilename, 1))

		if _, err := os.Stat(newPathname); err != nil {
			return newPathname, nil
		}
		count++
	}
	return "", errors.New("x Could not find available filename for: " + currentFilename)
}

func renamer(cmd *cobra.Command, args []string) {
	fmt.Println()
	fmt.Printf("PHOTO RENAMER\n")
	fmt.Printf("=============\n")
	fmt.Println()

	directory := args[0]
	fmt.Printf("Will rename images in directory '%s', continue? (yes/no)\n", directory)

	if utils.AskForConfirmation() {
		startTime := time.Now()

		filepath.Walk(directory, visit)

		fmt.Println()
		fmt.Println()
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		if counter == 0 {
			fmt.Printf("¯\\_(ツ)_/¯ Could not renamed any photo in %s\n", duration)
		} else {
			fmt.Printf("ʕ◔ϖ◔ʔ I successfully renamed %d photos in %s\n", counter, duration)
		}
	} else {
		fmt.Printf("Cancelled")
	}
	fmt.Println()
}

func renamerArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires a directory argument")
	}
	return nil
}

// var renamerCmd = &cobra.Command{
// 	Use:   "renamer",
// 	Short: "Renames image files based on their EXIF info",
// 	Long:  "",
// 	Args:  renamerArgs,
// 	Run:   renamer,
// }

// func init() {
// 	rootCmd.AddCommand(renamerCmd)
// }
