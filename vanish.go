package vanish

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

func RemoveFields(jsonData []byte, fields []string) ([]byte, error) {
	var i interface{}

	// Try to check if string is JSON
	err := json.Unmarshal(jsonData, &i)
	if err != nil {
		return nil, errors.New("the input data is not JSON format. Caused by: " + err.Error())
	}
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil, errors.New("the input data is not JSON format")
	}

	return json.Marshal(rebuild(m, fields))
}

func rebuild(m map[string]interface{}, fields []string) map[string]interface{} {
	res := make(map[string]interface{})
	currents, childs := getCurrentRemoveFields(fields)
	for _, field := range currents {
		delete(m, field)
	}

	for key, val := range m {
		valOf := reflect.ValueOf(val)

		switch valOf.Kind() {
		case reflect.Map:
			{
				if m, ok := val.(map[string]interface{}); ok {
					res[key] = rebuild(m, childs)
				}
			}

		case reflect.Slice, reflect.Array:
			{
				var slice []interface{}

				for i := 0; i < valOf.Len(); i++ {
					item := valOf.Index(i).Interface()

					if m, ok := item.(map[string]interface{}); ok {
						item = rebuild(m, childs)
					}

					slice = append(slice, item)
				}

				res[key] = slice
			}

		default:
			res[key] = val
		}
	}

	return res
}

// Get removable fields
func getCurrentRemoveFields(fields []string) ([]string, []string) {
	var currents, childs []string
	for _, f := range fields {
		r := strings.SplitN(f, ".", 2)
		switch len(r) {
		case 1:
			currents = append(currents, r[0])
		case 2:
			childs = append(childs, r[1])
		}
	}

	return currents, childs
}
