package cli

import (
	c "github.com/gotterdemarung/cfmt"
	"github.com/gotterdemarung/go-reports"
	"io"
	"os"
	"strings"
)

type formatter func(string) c.Format

func protoColor(fg int) formatter {
	return func(in string) c.Format {
		return c.Format{
			Value: in,
			Fg:    fg,
		}
	}
}

func protoColorB(fg, bg int) formatter {
	return func(in string) c.Format {
		return c.Format{
			Value: in,
			Fg:    fg,
			Bg:    bg,
		}
	}
}

var nocolor = func(in string) c.Format {
	return c.Format{
		Value: in,
	}
}

type palette struct {
	Title           formatter
	TitleDeco       formatter
	Description     formatter
	TableHeader     formatter
	MarkerHighlight formatter
	MarkerPositive  formatter
	MarkerNegative  formatter
	Grid            formatter
}

var colored = palette{
	Title:           protoColor(28),
	TitleDeco:       protoColor(22),
	Description:     protoColor(238),
	TableHeader:     protoColor(243),
	MarkerHighlight: protoColor(255),
	MarkerPositive:  protoColor(119),
	MarkerNegative:  protoColor(202),
	Grid:            protoColor(235),
}

var notcolored = palette{
	Title:           nocolor,
	TitleDeco:       nocolor,
	Description:     nocolor,
	TableHeader:     nocolor,
	MarkerHighlight: nocolor,
	MarkerPositive:  nocolor,
	MarkerNegative:  nocolor,
	Grid:            nocolor,
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

	c.Fprint(w, "\n")
	c.Fprint(w, c.FHeader(r.Title))
	c.Fprint(w, "\n")

	for _, ch := range r.Chunks {
		if ch.Priority >= minPriority {
			printChunk(ch, w, minPriority, pal)
		}
	}

	c.Fprint(w, "\n")
}

func printChunk(ch reports.Chunk, w io.Writer, minPriority int8, pal *palette) {

	c.Fprint(w, "\n")
	c.Fprint(w, c.FHeader(ch.Title))
	if ch.Description != "" {
		c.Fprint(w, " ", pal.Description(ch.Description), "\n")
	}
	c.Fprint(w, "\n")

	// Calculating lengths
	lengths := make([]int, len(ch.Headers))
	for i, h := range ch.Headers {
		lengths[i] = len(h.Title) + 2
	}

	for _, r := range ch.Rowset {
		for i, cc := range r.Data {
			l := len(cc.String()) + 2
			if l > lengths[i] {
				lengths[i] = l
			}
		}
	}

	// Printing
	for i, h := range ch.Headers {
		printHeader(h, lengths[i], w, pal)
		if i < len(ch.Headers)-1 {
			c.Fprint(w, " ")
		}
	}
	c.Fprint(w, "\n")
	for i, _ := range ch.Headers {
		c.Fprint(w, pal.Grid(strings.Repeat("─", lengths[i])))
		if i < len(ch.Headers)-1 {
			c.Fprint(w, pal.Grid("┴"))
		}
	}
	c.Fprint(w, "\n")
	for _, r := range ch.Rowset {
		for i, cc := range r.Data {
			printCell(cc, lengths[i], w, pal)
			c.Fprint(w, " ")
		}
		c.Fprint(w, "\n")
	}
	c.Fprint(w, "\n")
}

func printHeader(ch reports.Header, width int, w io.Writer, pal *palette) {
	toPrint := ch.Title
	if len(toPrint) < width {
		before := (width - len(toPrint)) / 2
		after := width - len(toPrint) - before
		toPrint = strings.Repeat(" ", before) + toPrint + strings.Repeat(" ", after)
	}

	c.Fprint(w, pal.TableHeader(toPrint))
}

func printCell(cc reports.Cell, width int, w io.Writer, pal *palette) {
	toPrint := cc.String()
	if len(toPrint) < width {
		if cc.GetAlign() == reports.ALIGN_RIGHT {
			toPrint = strings.Repeat(" ", width-len(toPrint)-1) + toPrint + " "
		} else {
			toPrint = " " + toPrint + strings.Repeat(" ", width-len(toPrint)-1)
		}
	}

	var f c.Format
	switch cc.Marker {
	case reports.MARKER_HIGHLIGHT:
		f = pal.MarkerHighlight(toPrint)
	case reports.MARKER_NEGATIVE:
		f = pal.MarkerNegative(toPrint)
	case reports.MARKER_POSITIVE:
		f = pal.MarkerPositive(toPrint)
	default:
		f = nocolor(toPrint)
	}

	c.Fprint(w, f)
}
