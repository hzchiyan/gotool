package structs

import "reflect"

func getFields(v reflect.Value, tagName string) []*Field {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	var fields []*Field
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag := field.Tag.Get(tagName); tag == "-" {
			continue
		}
		f := &Field{
			field: field,
			value: v.FieldByName(field.Name),
		}
		fields = append(fields, f)
	}
	return fields
}
