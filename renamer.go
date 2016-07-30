package main

import (
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
		fmt.Printf("x Unable to load EXIF data for file: %s", pathname)
		return nil
	}

	exifDatetime, _ := exifData.DateTime()

	newFilename := fmt.Sprintf("%s%s", exifDatetime.Format("20060102030405"), path.Ext(pathname))
	newPathname := strings.ToLower(strings.Replace(pathname, f.Name(), newFilename, 1))

	os.Rename(pathname, newPathname)
	fmt.Printf("✓ Renamed to: %s\n", newPathname)
	counter += 1

	return nil
}

func main() {
	flag.Parse()
	directory := flag.Arg(0)

	startTime := time.Now()
	filepath.Walk(directory, visit)
	endTime := time.Now()

	fmt.Printf("\n\nʕ◔ϖ◔ʔ I successfully renamed %d photos in %s", counter, endTime.Sub(startTime))
}
