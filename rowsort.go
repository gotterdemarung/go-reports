package reports

import (
	"fmt"
	"sort"
)

// Utility structure, used for sorting
type rowSorter struct {
	rowset Rowset
	column int
	asc    bool
}

// Sorts by column
func (p Rowset) SortBy(columnsIndex int, ascending bool) error {
	if columnsIndex < 0 {
		return fmt.Errorf("Invalid column index %d", columnsIndex)
	}
	if columnsIndex >= len(p) {
		return fmt.Errorf("Column %d not found in rowset. Total columns is ", columnsIndex, len(p))
	}

	rs := rowSorter{
		rowset: p,
		column: columnsIndex,
		asc:    ascending,
	}

	sort.Sort(rs)
	return nil
}

// Sort.Interface
func (p rowSorter) Len() int {
	return len(p.rowset)
}

// Sort.Interface
func (p rowSorter) Swap(i, j int) {
	p.rowset[i], p.rowset[j] = p.rowset[j], p.rowset[i]
}

// Sort.Interface
func (p rowSorter) Less(i, j int) bool {
	return p.asc == p.rowset[i].Data[p.column].Value.Less(
		p.rowset[j].Data[p.column].Value,
	)
}
