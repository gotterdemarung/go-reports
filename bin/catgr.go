package main

import (
	"fmt"
	"github.com/gotterdemarung/go-reports"
	"github.com/gotterdemarung/go-reports/cli"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: catgr <reportfile> [detail level]")
		fmt.Println("")
		os.Exit(1)
	}

	filename := os.Args[1]
	bts, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to read file", filename)
		fmt.Println("")
		os.Exit(1)
	}

	r, err := reports.UnpackReport(bts)
	if err != nil {
		fmt.Println("Unable to unpack report from", filename)
		fmt.Println("")
		os.Exit(1)
	}

	detail := 0
	if len(os.Args) > 2 {
		detail, _ = strconv.Atoi(os.Args[2])
	}

	cli.ReportPrinter(nil, int8(detail), true)(*r)
}
