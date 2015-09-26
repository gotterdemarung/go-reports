package reports

import (
	"strconv"
)

type Value struct {
	Stringv				*string			`json:"s,omitempty"`
	Numberv				*float64		`json:"n,omitempty"`
	Precision			*int8			`json:"p,omitempty"`
}

var emptyString = ""
var zeroPrecision int8 = 0
var minus1Precision int8 = -1

// Creates new value object
func NewValue(o interface{}) Value {
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

// Returns true if value contains string
func (v Value) IsNumber() bool {
	return v.Numberv != nil && v.Precision != nil
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