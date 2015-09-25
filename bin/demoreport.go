package main

import (
	"os"
	"github.com/gotterdemarung/go-reports"
	"github.com/gotterdemarung/go-reports/cli"
)

func main() {
	// Building demo report
	demo := reports.NewReport("Demo report", "demo", 1)

	// Building chunk of data
	ch := reports.NewChunk("Chunk #1", "This is example chunk")
	ch.AddHeaders("Name", "Integers", "Floats")
	ch.AddRow("Foo", 15, 0.3)
	ch.AddRow("Log bar", 15863, 1.001)
	ch.AddRow("Third one", reports.Cell{Value: -82, Marker: reports.MARKER_HIGHLIGHT}, 1.001)
	ch.AddRow("4", reports.Cell{Value: -82, Marker: reports.MARKER_NEGATIVE}, 1.001)
	ch.AddRow("*****", reports.Cell{Value: -82, Marker: reports.MARKER_POSITIVE}, 1.001)


	demo.Add(*ch)

	if len(os.Args) > 1 {
		// Printing to file
	} else {
		// Colors
		cli.ReportPrinter(nil, 0, true)(*demo)

		// No colors
		cli.ReportPrinter(nil, 0, false)(*demo)
	}
}