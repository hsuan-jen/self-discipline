package utils

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	OptIgnore    = "-"
	OptOmitempty = "omitempty"
	OptDive      = "dive"
	OptWildcard  = "wildcard"
)

const (
	flagIgnore = 1 << iota
	flagOmiEmpty
	flagDive
	flagWildcard
)

// Struct2Map convert a golang sturct to a map
// key can be specified by tag, LIKE `json:"tag"`.
// If there is no tag, struct filed name will be used instead
func Struct2Map(s interface{}) (res map[string]interface{}, err error) {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return nil, fmt.Errorf("%s is a nil pointer", v.Kind().String())
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// only accept struct param
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("s is not a struct but %s", v.Kind().String())
	}

	t := v.Type()
	res = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)

		// ignore unexported field
		if fieldType.PkgPath != "" {
			continue
		}
		// read tag
		tagVal, flag := readTag(fieldType)

		if flag&flagIgnore != 0 {
			continue
		}

		fieldValue := v.Field(i)
		if flag&flagOmiEmpty != 0 && fieldValue.IsZero() {
			continue
		}

		// ignore nil pointer in field
		if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
			continue
		}
		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}

		// get kind
		switch fieldValue.Kind() {
		case reflect.Slice, reflect.Array:
			res[tagVal] = fieldValue
		case reflect.Struct:
			// recursive
			deepRes, deepErr := Struct2Map(fieldValue.Interface())
			if deepErr != nil {
				return nil, deepErr
			}
			if flag&flagDive != 0 {
				for k, v := range deepRes {
					res[k] = v
				}
			} else {
				res[tagVal] = deepRes
			}
		case reflect.Map:
			res[tagVal] = fieldValue
		case reflect.Chan:
			res[tagVal] = fieldValue
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
			res[tagVal] = fieldValue.Int()
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
			res[tagVal] = fieldValue.Uint()
		case reflect.Float32, reflect.Float64:
			res[tagVal] = fieldValue.Float()
		case reflect.String:
			if flag&flagWildcard != 0 {
				res[tagVal] = "%" + fieldValue.String() + "%"
			} else {
				res[tagVal] = fieldValue.String()
			}
		case reflect.Bool:
			res[tagVal] = fieldValue.Bool()
		case reflect.Complex64, reflect.Complex128:
			res[tagVal] = fieldValue.Complex()
		case reflect.Interface:
			res[tagVal] = fieldValue.Interface()
		default:
		}
	}
	return
}

// readTag read tag with format `json:"name,omitempty"` or `json:"-"`
// For now, only supports above format
func readTag(f reflect.StructField) (string, int) {
	tag := "json"
	val, ok := f.Tag.Lookup(tag)
	fieldTag := ""
	flag := 0

	// no tag, skip this field
	if !ok {
		flag |= flagIgnore
		return "", flag
	}
	opts := strings.Split(val, ",")

	fieldTag = opts[0]
	for i := 0; i < len(opts); i++ {
		switch opts[i] {
		case OptIgnore:
			flag |= flagIgnore
		case OptOmitempty:
			flag |= flagOmiEmpty
		case OptDive:
			flag |= flagDive
		case OptWildcard:
			flag |= flagWildcard
		}
	}

	return fieldTag, flag
}
