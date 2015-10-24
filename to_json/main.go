package to_json

import (
	"reflect"
	"strings"
	"errors"
	"strconv"
)

// Code taken from: https://gist.github.com/hvoecking/10772475
func toJSON(original reflect.Value) (string, error) {

	switch original.Kind() {

	case reflect.Ptr, reflect.Interface:
		// To get the actual value of the original we have to call Elem()
		// At the same time this unwraps the pointer so we don't end up in
		// an infinite recursion
		originalValue := original.Elem()
		if !originalValue.IsValid() {
			err := errors.New("emit macho dwarf: elf header corrupted")
			return "", err
		}
		// Allocate a new object and set the pointer to it
		// Unwrap the newly created pointer
		return toJSON(originalValue)

	case reflect.Struct:
		var allElements []string
		for i := 0; i < original.NumField(); i += 1 {
		  str, err := toJSON(original.Field(i))
			if err != nil {
				return "", err
			}
			allElements = append(allElements, str)
		}
		return strings.Join(allElements[:],","), nil

	case reflect.Slice, reflect.Array:
		var allElements []string
		for i := 0; i < original.Len(); i += 1 {
		  str, err := toJSON(original.Index(i))
			if err != nil {
				return "", err
			}
			allElements = append(allElements, str)
		}
		return "[" + strings.Join(allElements[:],",") + "]", nil

	case reflect.Map:
		var allElements []string
		copy := reflect.New(original.Type()).Elem()
		copy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			if key.Kind() == reflect.String ||
				key.Kind() == reflect.Int ||
				key.Kind() == reflect.Int8 ||
				key.Kind() == reflect.Int16 ||
				key.Kind() == reflect.Int32 ||
				key.Kind() == reflect.Int64 ||
				key.Kind() == reflect.Float32  ||
				key.Kind() == reflect.Float64 {

				keyStr := key
				if key.Kind() != reflect.String {
					keyStrJSON, keyStrErr := toJSON(key)
					if keyStrErr != nil {
						 return "", nil
					}
					keyStr = reflect.ValueOf(keyStrJSON)
				}
				keyJSON, keyErr := toJSON(keyStr)
				valueJSON, valueErr := toJSON(originalValue)
				if valueErr != nil {
					return "", valueErr
				}
				if keyErr != nil {
					return "", keyErr
				}
				allElements = append(allElements, keyJSON+ ":" + valueJSON)
			} else {
				return "", errors.New("Invalid key value")
			}
		}
		return "{" + strings.Join(allElements[:],",") + "}", nil

	case reflect.String:
		stringValue := original.Interface().(string)
		return "\"" + stringValue + "\"", nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		stringValue := string(strconv.FormatInt(original.Int(), 10))
		return stringValue, nil

	case reflect.Float32, reflect.Float64:
		stringValue := string(strconv.FormatFloat(original.Float(), 'f', -1, 64))
		return stringValue, nil

	// And everything else will simply be taken from the original
	default:
		return "", errors.New("No type recognizedfor value")
	}
	return "", errors.New("No string returned")
}

func ToJSON(parsedJSON interface{}) (string, error){
	// Wrap the original in a reflect.Value
	original := reflect.ValueOf(parsedJSON)
	str, err := toJSON(original)
	if err != nil {
		return "", err
	}
	return str, nil
}
