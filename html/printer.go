package html

import (
	"bytes"
	"github.com/gotterdemarung/go-reports"
	"strings"
)

type Configuration struct {
	ReportClasses string
	ChunksClasses string
}

func ToHtml(c Configuration, r reports.Report) []byte {
	buf := bytes.NewBuffer([]byte{})

	openTag(buf, "div", c.ReportClasses)

	// Report header
	openTag(buf, "h1", "")
	buf.WriteString(r.Title)
	closeTag(buf, "h1")
	buf.WriteRune('\n')

	// Chunks
	for _, ch := range r.Chunks {
		printChunk(c, buf, ch)
		buf.WriteRune('\n')
	}

	closeTag(buf, "div")

	return buf.Bytes()
}

func printChunk(c Configuration, buf *bytes.Buffer, ch reports.Chunk) {
	openTag(buf, "div", c.ChunksClasses)

	// Chunk header
	openTag(buf, "h1", "")
	buf.WriteString(ch.Title)
	closeTag(buf, "h1")
	if ch.Description != "" {
		openTag(buf, "i", "")
		buf.WriteString(ch.Description)
		closeTag(buf, "i")
	}
	buf.WriteRune('\n')

	// Table
	buf.WriteString("<table border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n")
	buf.WriteString("<thead><tr>")
	for _, h := range ch.Headers {
		buf.WriteString("<th>")
		buf.WriteString(h.Title)
		buf.WriteString("</th>")
	}
	buf.WriteString("</tr></thead>\n")
	buf.WriteString("<tbody>")
	for _, row := range ch.Rowset {
		buf.WriteString("<tr>")
		for _, cell := range row.Data {
			printCell(c, buf, cell)
		}
		buf.WriteString("</tr>")
	}
	buf.WriteString("</tbody>")
	buf.WriteString("</table>\n")

	closeTag(buf, "div")
}

func printCell(c Configuration, buf *bytes.Buffer, cell reports.Cell) {
	classes := ""
	switch cell.Marker {
	case reports.MARKER_HIGHLIGHT:
		classes += " highlight"
	case reports.MARKER_NEGATIVE:
		classes += " negative"
	case reports.MARKER_POSITIVE:
		classes += " positive"
	}
	switch cell.GetAlign() {
	case reports.ALIGN_CENTER:
		classes += " center"
	case reports.ALIGN_RIGHT:
		classes += " right"
	default:
		classes += " left"
	}

	if cell.Value.IsNumber() {
		classes += " number"
	}

	openTag(buf, "td", classes)
	buf.WriteString(cell.String())
	closeTag(buf, "td")
	buf.WriteString("</td>")
}

func openTag(b *bytes.Buffer, name, classes string) {
	b.WriteRune('<')
	b.WriteString(name)
	if classes != "" {
		classes = strings.Trim(classes, " ")
		b.WriteString(" class=")
		b.WriteRune('"')
		b.WriteString(classes)
		b.WriteRune('"')
	}
	b.WriteRune('>')
}

func closeTag(b *bytes.Buffer, name string) {
	b.WriteString("</")
	b.WriteString(name)
	b.WriteRune('>')
}
