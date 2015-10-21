package cli

import (
	"github.com/gotterdemarung/go-reports"
	"github.com/mgutz/ansi"
	"io"
	"os"
	"strings"
)

var ln = []byte("\n")

var nocolor = func(in string) string {
	return in
}

type palette struct {
	Title           func(string) string
	TitleDeco       func(string) string
	Description     func(string) string
	TableHeader     func(string) string
	MarkerHighlight func(string) string
	MarkerPositive  func(string) string
	MarkerNegative  func(string) string
}

var colored = palette{
	Title:           ansi.ColorFunc("28"),
	TitleDeco:       ansi.ColorFunc("22"),
	Description:     ansi.ColorFunc("238"),
	TableHeader:     ansi.ColorFunc("238:233"),
	MarkerHighlight: ansi.ColorFunc("white+h"),
	MarkerPositive:  ansi.ColorFunc("119"),
	MarkerNegative:  ansi.ColorFunc("202"),
}

var notcolored = palette{
	Title:           nocolor,
	TitleDeco:       nocolor,
	Description:     nocolor,
	TableHeader:     nocolor,
	MarkerHighlight: nocolor,
	MarkerPositive:  nocolor,
	MarkerNegative:  nocolor,
}

// Return report printer function
func ReportPrinter(w io.Writer, minPriority int8, colors bool) func(r reports.Report) error {
	var pal *palette
	if colors {
		pal = &colored
	} else {
		pal = &notcolored
	}

	if w == nil {
		w = os.Stdout
	}

	return func(r reports.Report) error {
		printReport(r, w, minPriority, pal)
		return nil
	}
}

// Prints report to provider writer
func printReport(r reports.Report, w io.Writer, minPriority int8, pal *palette) {

	w.Write(ln)
	w.Write([]byte(pal.Title(" " + r.Title + "\n")))
	w.Write([]byte(pal.TitleDeco(" " + strings.Repeat("=", len(r.Title)) + "\n")))

	for _, c := range r.Chunks {
		if c.Priority >= minPriority {
			printChunk(c, w, minPriority, pal)
		}
	}

	w.Write(ln)
}

func printChunk(c reports.Chunk, w io.Writer, minPriority int8, pal *palette) {
	w.Write(ln)
	w.Write([]byte(pal.Title(" " + c.Title + "\n")))
	w.Write([]byte(pal.TitleDeco(" " + strings.Repeat("-", len(c.Title)) + "\n")))
	if c.Description != "" {
		w.Write([]byte(pal.Description(" " + c.Description + "\n")))
	}
	w.Write(ln)

	// Calculating lengths
	lengths := make([]int, len(c.Headers))
	for i, h := range c.Headers {
		lengths[i] = len(h.Title) + 2
	}

	for _, r := range c.Rowset {
		for i, c := range r.Data {
			l := len(c.String()) + 2
			if l > lengths[i] {
				lengths[i] = l
			}
		}
	}

	// Printing
	for i, h := range c.Headers {
		printHeader(h, lengths[i], w, pal)
		w.Write([]byte(" "))
	}
	w.Write(ln)
	for _, r := range c.Rowset {
		for i, c := range r.Data {
			printCell(c, lengths[i], w, pal)
			w.Write([]byte(" "))
		}
		w.Write(ln)
	}
	w.Write(ln)
}

func printHeader(c reports.Header, width int, w io.Writer, pal *palette) {
	toPrint := c.Title
	if len(toPrint) < width {
		before := (width - len(toPrint)) / 2
		after := width - len(toPrint) - before
		toPrint = strings.Repeat(" ", before) + toPrint + strings.Repeat(" ", after)
	}

	w.Write([]byte(pal.TableHeader(toPrint)))
}

func printCell(c reports.Cell, width int, w io.Writer, pal *palette) {
	toPrint := c.String()
	if len(toPrint) < width {
		if c.GetAlign() == reports.ALIGN_RIGHT {
			toPrint = strings.Repeat(" ", width-len(toPrint)-1) + toPrint + " "
		} else {
			toPrint = " " + toPrint + strings.Repeat(" ", width-len(toPrint)-1)
		}
	}

	switch c.Marker {
	case reports.MARKER_HIGHLIGHT:
		toPrint = pal.MarkerHighlight(toPrint)
	case reports.MARKER_NEGATIVE:
		toPrint = pal.MarkerNegative(toPrint)
	case reports.MARKER_POSITIVE:
		toPrint = pal.MarkerPositive(toPrint)
	}

	w.Write([]byte(toPrint))
}
