package main

import (
	"flag"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"os"
	"path/filepath"
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

	filename := f.Name()
	exifDatetime, _ := exifData.DateTime()

	if err := os.Chtimes(pathname, exifDatetime, exifDatetime); err != nil {
		fmt.Printf("%s\n", err)
		return nil
	}

	fmt.Printf("✓ Updated file %s mtime to %v\n", filename, exifDatetime)
	counter += 1

	return nil
}

func main() {
	flag.Parse()
	directory := flag.Arg(0)

	startTime := time.Now()
	filepath.Walk(directory, visit)
	endTime := time.Now()

	fmt.Printf("\n\nʕ◔ϖ◔ʔ I successfully updated %d photos in %s\n", counter, endTime.Sub(startTime))
}
