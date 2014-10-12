package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// ET, not EST
	var d string = "2014/01/06 7:20:20 ET"
	loc, _ := time.LoadLocation("US/Eastern")

	// DateString Layout for parsing
	var layout string = "2006/01/02 15:04:05"

	test, err := time.ParseInLocation(layout, d[:strings.LastIndex(d, " ")], loc)
	if err != nil {
		panic(err)
	}

	fmt.Println(test)
	fmt.Println(test.UTC())

}
