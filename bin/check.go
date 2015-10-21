package main

import (
	"github.com/gotterdemarung/go-reports"
	"github.com/gotterdemarung/go-reports/cli"
	"io/ioutil"
)

func main() {

	r := reports.NewReport("ololo", "trololo")

	ch := reports.NewChunk("Chunk #1", "This is example chunk")
	ch.AddHeaders("Name", "Integers", "Floats")
	ch.AddRow("Foo", 15, 0.3)
	ch.AddRow("Foo 0", 0, "")
	ch.AddRow("Log bar", 15863, 1.001)
	ch.AddRow("Third one", reports.NewCell(-82, "", reports.MARKER_HIGHLIGHT), 1.001)
	ch.AddRow("4", reports.NewCell(-82, "", reports.MARKER_POSITIVE), 8.9)
	ch.AddRow("*****", reports.NewFloatCell(-82, 3, "", reports.MARKER_NEGATIVE), 91.2735)

	r.Add(*ch)

	cli.ReportPrinter(nil, 0, true)(*r)

	b0 := reports.PackReport(*r)
	rrrr, err := reports.UnpackReport(b0)
	if err != nil {
		panic(err)
	}

	cli.ReportPrinter(nil, 0, true)(*rrrr)

	// Writing
	r.SaveToFile("test.gr")

	// Reading
	bts, _ := ioutil.ReadFile("test.gr")
	r2, err := reports.UnpackReport(bts)
	if err != nil {
		panic(err)
	}

	cli.ReportPrinter(nil, 0, true)(*r2)
}
