package reports

import (
	"fmt"
	"time"
	"bytes"
	"encoding/gob"
	"encoding/binary"
	"compress/gzip"
)

type packedReport struct {
	CreatedAt 			time.Time								// Create at time
	Type 				string									// Report type
	Title 				string									// Report title
	Filters 			Filters									// Filters for report
	Chunks 				Chunks									// Data for report
}

func PackReport(r Report) []byte {
	pr := packedReport{
		CreatedAt:  r.CreatedAt,
		Type: 		r.Type,
		Title: 		r.Title,
		Filters: 	r.Filters,
		Chunks: 	r.Chunks,
	}

	buf := new(bytes.Buffer)
	w, _ := gzip.NewWriterLevel(buf, gzip.BestCompression)
	enc := gob.NewEncoder(w)
	_ = enc.Encode(pr)
	w.Flush()
	return buf.Bytes()
}

func UnpackReport(x []byte) (*Report, error) {
	var ff packedReport
	r, err := gzip.NewReader(bytes.NewReader(x))
	if err != nil {
		return nil, err
	}
	dec := gob.NewDecoder(r)
	err = dec.Decode(&ff)
	if err != nil {
		return nil, err
	}

	rep := Report{
		Type: 			ff.Type,
		Title:			ff.Title,
		Filters: 		ff.Filters,
		Chunks: 		ff.Chunks,
		CreatedAt: 		ff.CreatedAt,
		UpdatedAt: 		ff.CreatedAt,
	}

	return &rep, nil
}

const (
	pack_marker_number byte = 0
	pack_marker_string byte = 1
)

func (v *Value) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	if v.IsNumber() {
		buf.WriteByte(pack_marker_number)
		buf.WriteByte(byte(*v.Precision))
		binary.Write(buf, binary.LittleEndian, *v.Numberv)
	} else {
		buf.WriteByte(pack_marker_string)
		sb := []byte(*v.Stringv)
		buf.Write(sb)
	}

	return buf.Bytes(), nil
}

func (v *Value) UnmarshalBinary(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Invalid data len")
	}

	if data[0] == pack_marker_number {
		// Number
		precision := int8(data[1])
		var f64 float64
		binary.Read(bytes.NewReader(data[2:]), binary.LittleEndian, &f64)

		v.Stringv = nil
		v.Precision = &precision
		v.Numberv = &f64
	} else {
		// String
		str := string(data[1:])

		v.Stringv = &str
		v.Numberv = nil
		v.Precision = nil
	}

	return nil
}
