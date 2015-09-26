package reports_test

import (
	"testing"
	"github.com/gotterdemarung/go-reports"
)

func TestFiltersPacking(t *testing.T) {
	ff := map[string]interface{}{}
	ff["foo"] = "bar"
	ff["baz"] = true

	bts, err := reports.PackFilters(ff)
	if err != nil {
		t.Fatal(err)
	}

	ff2, err := reports.UnpackFilters(bts)
	if err != nil {
		t.Fatal(err)
	}
	if ff2 == nil {
		t.Log(bts)
		t.Fatal("Nil response from unpacker")
	}

	if len(ff) != len(*ff2) {
		t.Fatal("Length mismatch")
	}
}