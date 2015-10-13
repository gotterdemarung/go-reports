package reports

import (
	"io/ioutil"
)

// Packs and saves report to file
func (r *Report) SaveToFile(filename string) error {
	bts := PackReport(*r)
	return ioutil.WriteFile(filename, bts, 0644)
}