package reports

import (
	"time"
)

const (
	ALIGN_LEFT 			= 0
	ALIGN_RIGHT 		= 1
	ALIGN_CENTER 		= 2

	MARKER_NORMAL		= 0
	MARKER_HIGHLIGHT	= 1
	MARKER_POSITIVE		= 2
	MARKER_NEGATIVE		= 3
)

// Report structure
type Report struct {
	Id					*int									// Id
	CreatedAt 			time.Time								// Create at time
	UpdatedAt			time.Time								// Last update time
	Type 				string									// Report type
	TypeVersion			int										// Version of report type
	Title 				string									// Report title
	Filters 			Filters									// Filters for report
	Chunks 				[]Chunk									// Data for report
}

// Filters for report
type Filters map[string]interface{}

// Report chunk
type Chunk struct {
	Title 				string			`json:"name"`			// Chunk title
	Description 		string          `json:"desc"`			// Chunk description
	Headers				Headers			`json:"head"`			// Headers information
	Rowset 				Rowset			`json:"rows"`			// Main report data
	Priority			int8			`json:"priority"`		// Display priority
}

// Report chunk headers
type Headers []Header

// Report header
type Header struct {
	Title 				string			`json:"n"`				// Header name/title
	Sortable 			bool			`json:"o"`				// Allows sorting
	Searchable  		bool			`json:"f"`				// Allows filtering
	Priority			int8			`json:"p"`				// Display priority
}

// Report rowset
type Rowset []Row

// Report row
type Row struct {
	Data 				[]Cell			`json:"d"`				// Main row data
	Marker				*int			`json:"m"`				// Marker
	Comment 			*string 		`json:"c"`				// Comment text
}

// Report value
type Cell struct {
	Value 				Value			`json:"v"`				// Cell value
	Description			string			`json:"d"`				// Cell description
	Marker 				int				`json:"m"`				// Cell marker
}

// Creates new report
func NewReport(name, reporttype string, version int) *Report {
	return &Report{
		CreatedAt:		time.Now().UTC(),
		UpdatedAt:		time.Now().UTC(),
		Title:			name,
		TypeVersion: 	version,
		Filters: 		map[string]interface{}{},
		Chunks: 		[]Chunk{},
	}
}

// Appends new chunk to report
func (r *Report) Add(c Chunk) {
	r.Chunks = append(r.Chunks, c)
}

// Creates new chunk
func NewChunk(name, description string) *Chunk {
	return &Chunk{
		Title: 			name,
		Description: 	description,
		Headers: 		[]Header{},
		Rowset: 		[]Row{},
	}
}

// Adds simple sortable and searchable headers
func (c *Chunk) AddHeaders(names ...string) {
	for _, n := range names {
		h := Header{
			Title: 		n,
			Sortable: 	true,
			Searchable: true,
		}

		c.Headers = append(c.Headers, h)
	}
}

// Adds simple data row
func (c *Chunk) AddRow(values ...interface{}) {

	cells := make([]Cell, len(values))
	for i, v := range values {
		if cc, ok := v.(Cell); ok {
			cells[i] = cc;
		} else {
			cells[i] = Cell{
				Value: NewValue(v),
			}
		}
	}

	c.Rowset = append(c.Rowset, Row{
		Data: cells,
	})
}

func NewCell(val interface{}, Description string, Marker int) Cell {
	return Cell{
		Value: NewValue(val),
		Description: Description,
		Marker: Marker,
	}
}

func NewFloatCell(val float64, precision int8, Description string, Marker int) Cell {
	return Cell{
		Value: Value{Numberv: &val, Precision: &precision},
		Description: Description,
		Marker: Marker,
	}
}

// Cast cell value to string
func (c Cell) String() string {
	return c.Value.String()
}

func (c Cell) GetAlign() int {
	if c.Value.Numberv == nil {
		return ALIGN_LEFT;
	} else {
		return ALIGN_RIGHT;
	}
}
