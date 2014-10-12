package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func VisitFile(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}

	if !!fi.IsDir() {
		return nil // not a file.  ignore.
	}

	matched, err := filepath.Match("*.go", fi.Name())
	if err != nil {
		fmt.Println(err) // malformed pattern
		return err       // this is fatal.
	}

	if matched {
		do_something(fp, fi)
	}
	return nil
}

func do_something(fp string, fi os.FileInfo) {
	abs_file, _ := filepath.Abs(fp)
	fmt.Println(abs_file, fi.Size(), fi.Mode(), fi.ModTime().UTC())
}

func main() {
	filepath.Walk("..", VisitFile)
}
