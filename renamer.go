package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var counter int = 0

func visit(pathname string, f os.FileInfo, err error) error {
	openedFile, _ := os.Open(pathname)
	exifData, err := exif.Decode(openedFile)

	if err != nil {
		fmt.Printf("x Unable to load EXIF data for file: %s\n", pathname)
		return nil
	}

	currentFilename := f.Name()
	exifDatetime, _ := exifData.DateTime()
	formattedDatetime := exifDatetime.Format("20060102030405")
	newPathname, err := findAvailablePathname(pathname, currentFilename, formattedDatetime)

	if err != nil {
		fmt.Printf("%s\n", err)
		return nil
	}

	os.Rename(pathname, newPathname)
	fmt.Printf("✓ Renamed to: %s\n", newPathname)
	counter += 1

	return nil
}

func findAvailablePathname(pathname string, currentFilename string, formattedDatetime string) (string, error) {
	count := 0
	for count < 5 {
		newFilename := fmt.Sprintf("%s%s", formattedDatetime, path.Ext(pathname))
		if currentFilename == newFilename {
			return "", errors.New("- File already renamed: " + currentFilename)
		}
		if count > 0 {
			newFilename = fmt.Sprintf("%s-%v%s", formattedDatetime, count, path.Ext(pathname))
		}
		newPathname := strings.ToLower(strings.Replace(pathname, currentFilename, newFilename, 1))

		if _, err := os.Stat(newPathname); err != nil {
			return newPathname, nil
		}
		count += 1
	}
	return "", errors.New("x Could not find available filename for: " + currentFilename)
}

func main() {
	flag.Parse()
	directory := flag.Arg(0)

	startTime := time.Now()
	filepath.Walk(directory, visit)
	endTime := time.Now()

	fmt.Printf("\n\nʕ◔ϖ◔ʔ I successfully renamed %d photos in %s\n", counter, endTime.Sub(startTime))
}
