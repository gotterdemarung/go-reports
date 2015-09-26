package reports_test

import (
	"testing"
	"encoding/json"
	"github.com/gotterdemarung/go-reports"
)

func TestValueIsNumber(t *testing.T) {
	if reports.NewValue("foo").IsNumber() {
		t.Error("Expected false for strings")
	}
	if reports.NewValue(nil).IsNumber() {
		t.Error("Expected false for nil")
	}
	if !reports.NewValue(1).IsNumber() {
		t.Error("Expected true for integer")
	}
	if !reports.NewValue(.1).IsNumber() {
		t.Error("Expected true for floats")
	}
}

func TestValueMarshall(t *testing.T) {
	v := reports.NewValue("foo")

	bts, err := json.Marshal(v)
	if err != nil {
		t.Error(err)
	}
	if `{"s":"foo"}` != string(bts) {
		t.Log(string(bts))
		t.Error("JSON marshalling returns wrong result")
	}
	var v2 reports.Value = reports.Value{}
	err = json.Unmarshal(bts, &v2)
	if err != nil {
		t.Error(err)
	}
	if *v.Stringv != *v2.Stringv {
		t.Log(v, v2)
		t.Error("JSON unmarshalling error")
	}



	v = reports.NewValue(88);
	bts, err = json.Marshal(v)
	if err != nil {
		t.Error(err)
	}
	if `{"n":88,"p":0}` != string(bts) {
		t.Log(string(bts))
		t.Error("JSON marshalling returns wrong result")
	}

	v = reports.NewValue(.1);
	bts, err = json.Marshal(v)
	if err != nil {
		t.Error(err)
	}
	if `{"n":0.1,"p":-1}` != string(bts) {
		t.Log(string(bts))
		t.Error("JSON marshalling returns wrong result")
	}
}
