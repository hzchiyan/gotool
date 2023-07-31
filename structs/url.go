package structs

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

// Encoder is an interface implemented by any type that wishes to encode
// itself into URL values in a non-standard way.
type Encoder interface {
	EncodeValues(key string, v *url.Values) error
}

var timeType = reflect.TypeOf(time.Time{})
var encoderType = reflect.TypeOf(new(Encoder)).Elem()

// reflectValue populates the values parameter from the struct fields in val.
// Embedded structs are followed recursively (using the rules defined in the
// Values function documentation) breadth-first.
func reflectValue(values url.Values, val reflect.Value, scope string) error {
	var embedded []reflect.Value
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous { // unexported
			continue
		}
		sv := val.Field(i)
		tag := sf.Tag.Get(DefaultUrlTagName)
		if tag == "-" {
			continue
		}
		name, opts := parseTag(tag)
		if name == "" {
			if sf.Anonymous {
				v := reflect.Indirect(sv)
				if v.IsValid() && v.Kind() == reflect.Struct {
					// save embedded struct for later processing
					embedded = append(embedded, v)
					continue
				}
			}
			name = sf.Name
		}
		if scope != "" {
			name = scope + "[" + name + "]"
		}
		if opts.Has("omitempty") && isEmptyValue(sv) {
			continue
		}
		if sv.Type().Implements(encoderType) {
			// if sv is a nil pointer and the custom encoder is defined on a non-pointer
			// method receiver, set sv to the zero value of the underlying type
			if !reflect.Indirect(sv).IsValid() && sv.Type().Elem().Implements(encoderType) {
				sv = reflect.New(sv.Type().Elem())
			}
			m := sv.Interface().(Encoder)
			if err := m.EncodeValues(name, &values); err != nil {
				return err
			}
			continue
		}
		// recursively dereference pointers. break on nil pointers
		for sv.Kind() == reflect.Ptr {
			if sv.IsNil() {
				break
			}
			sv = sv.Elem()
		}
		if sv.Kind() == reflect.Slice || sv.Kind() == reflect.Array {
			if sv.Len() == 0 {
				// skip if slice or array is empty
				continue
			}
			var del string
			if opts.Has("comma") {
				del = ","
			} else if opts.Has("space") {
				del = " "
			} else if opts.Has("semicolon") {
				del = ";"
			} else if opts.Has("brackets") {
				name = name + "[]"
			} else {
				del = sf.Tag.Get("del")
			}
			if del != "" {
				s := new(bytes.Buffer)
				first := true
				for i := 0; i < sv.Len(); i++ {
					if first {
						first = false
					} else {
						s.WriteString(del)
					}
					s.WriteString(valueString(sv.Index(i), opts, sf))
				}
				values.Add(name, s.String())
			} else {
				for i := 0; i < sv.Len(); i++ {
					k := name
					if opts.Has("numbered") {
						k = fmt.Sprintf("%s%d", name, i)
					}
					values.Add(k, valueString(sv.Index(i), opts, sf))
				}
			}
			continue
		}
		if sv.Type() == timeType {
			values.Add(name, valueString(sv, opts, sf))
			continue
		}
		if sv.Kind() == reflect.Struct {
			if err := reflectValue(values, sv, name); err != nil {
				return err
			}
			continue
		}
		values.Add(name, valueString(sv, opts, sf))
	}
	for _, f := range embedded {
		if err := reflectValue(values, f, scope); err != nil {
			return err
		}
	}
	return nil
}

// valueString returns the string representation of a value.
func valueString(v reflect.Value, opts tagOptions, sf reflect.StructField) string {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}
	if v.Kind() == reflect.Bool && opts.Has("int") {
		if v.Bool() {
			return "1"
		}
		return "0"
	}
	if v.Type() == timeType {
		t := v.Interface().(time.Time)
		if opts.Has("unix") {
			return strconv.FormatInt(t.Unix(), 10)
		}
		if opts.Has("unixmilli") {
			return strconv.FormatInt((t.UnixNano() / 1e6), 10)
		}
		if opts.Has("unixnano") {
			return strconv.FormatInt(t.UnixNano(), 10)
		}
		if layout := sf.Tag.Get("layout"); layout != "" {
			return t.Format(layout)
		}
		return t.Format(time.RFC3339)
	}

	return fmt.Sprint(v.Interface())
}

// isEmptyValue checks if a value should be considered empty for the purposes
// of omitting fields with the "omitempty" option.
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	type zeroable interface {
		IsZero() bool
	}
	if z, ok := v.Interface().(zeroable); ok {
		return z.IsZero()
	}
	return false
}
