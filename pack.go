package reports

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"compress/gzip"
	"encoding/json"
	"encoding/binary"
)


// Packs report to binary format
func (r *Report) PackBin() ([]byte, error) {
	return PackReport(*r)
}

// Packs whole report
func PackReport(r Report) ([]byte, error) {
	pf, err := PackFilters(r.Filters);
	if err != nil {
		return nil, err
	}
	pc, err := PackChunks(r.Chunks);
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer([]byte{})

	// String header
	strh := []byte(fmt.Sprintf(
		"%d\n%s\n%d\n%s\n",
		r.UpdatedAt.Unix(), r.Type, r.TypeVersion, r.Title,
	))

	// Version
	_, err = buf.WriteString("a")
	if err != nil {
		return nil, err
	}

	// Bin header
	err = binary.Write(buf, binary.LittleEndian, int32(len(strh)))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, int32(len(pf)))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, int32(len(pc)))
	if err != nil {
		return nil, err
	}

	// String header
	_, err = buf.Write(strh)
	if err != nil {
		return nil, err
	}

	// Packed filters
	_, err = buf.Write(pf)
	if err != nil {
		return nil, err
	}
	// Packed data
	_, err = buf.Write(pc)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Packs filters to byte slice
func PackFilters(f Filters) ([]byte, error) {
	bts, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		return nil, err
	}
	w.Write(bts)
	w.Close()

	return b.Bytes(), nil
}

// Packs chunks to byte slice
func PackChunks(c []Chunk) ([]byte, error) {
	bts, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		return nil, err
	}
	w.Write(bts)
	w.Close()

	return b.Bytes(), nil
}

// Unpacks filters from bytes
func UnpackFilters(bts []byte) (*Filters, error) {
	r, err := gzip.NewReader(bytes.NewReader(bts))
	if r != nil {
		return nil, err
	}

	var f Filters = map[string]interface{}{}
	all, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(all, &f)
	if err != nil {
		return nil, err
	}

	return &f, nil
}

// Unpacks chunks from bytes
func UnpackChunks(bts []byte) ([]Chunk, error) {
	r, err := gzip.NewReader(bytes.NewReader(bts))
	if r != nil {
		return nil, err
	}

	var c []Chunk
	all, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(all, &c)

	if err != nil {
		return nil, err
	}

	return c, nil
}