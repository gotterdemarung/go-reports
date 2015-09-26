package reports

import (
	"strconv"
)

type Value struct {
	Stringv				*string			`json:"s"`
	Numberv				*float64		`json:"n"`
	Precision			*int8			`json:"p"`
}

var emptyString = ""
var zeroPrecision int8 = 0
var minus1Precision int8 = -1

// Creates new value object
func newValue(o interface{}) Value {
	if o == nil {
		return Value{
			Stringv: &emptyString,
		}
	} else if v, ok := o.(Value); ok {
		return v
	} else if v, ok := o.(string); ok {
		return Value{
			Stringv: &v,
		}
	} else if v, ok := o.(int); ok {
		fv := float64(v)
		return Value{
			Numberv: &fv,
			Precision: &zeroPrecision,
		}
	} else if v, ok := o.(float64); ok {
		return Value{
			Numberv: &v,
			Precision: &minus1Precision,
		}
	} else {
		return Value{
			Stringv: &emptyString,
		}
	}
}

// Returns string representation of value
func (v *Value) String() string {
	if v.Stringv != nil {
		return *v.Stringv
	} else if v.Numberv != nil && v.Precision != nil {
		return strconv.FormatFloat(*v.Numberv, 'f', int(*v.Precision), 32)
	} else {
		return ""
	}
}