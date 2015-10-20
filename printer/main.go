package printer

import (
	"reflect"
	"strings"
	"errors"
)

// Code taken from: https://gist.github.com/hvoecking/10772475
func convertJSONToString(original reflect.Value) (string, error) {
	switch original.Kind() {

	case reflect.Ptr:
	case reflect.Interface:
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
		return convertJSONToString(originalValue)

	case reflect.Struct:
		var allElements []string
		for i := 0; i < original.NumField(); i += 1 {
		  str, err := convertJSONToString(original.Field(i))
			if err != nil {
				return "", err
			}
			allElements = append(allElements, str)
		}
		return strings.Join(allElements[:],","), nil

	case reflect.Slice:
		var allElements []string
		for i := 0; i < original.Len(); i += 1 {
		  str, err := convertJSONToString(original.Index(i))
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
			keyStr, keyErr := convertJSONToString(key)
			valueStr, valueErr := convertJSONToString(originalValue)
			if valueErr != nil {
				return "", valueErr
			}
			if keyErr != nil {
				return "", keyErr
			}
			allElements = append(allElements, keyStr + ":" + valueStr)
		}
		return "{" + strings.Join(allElements[:],",") + "}", nil

	case reflect.String:
		stringValue := original.Interface().(string)
		return "\"" + stringValue + "\"", nil

	// And everything else will simply be taken from the original
	default:
		return "--default--", nil
	}
	return "", nil
}

func PrintJSON(parsedJSON interface{}) (string, error){
	// Wrap the original in a reflect.Value
	original := reflect.ValueOf(parsedJSON)
	str, err := convertJSONToString(original)
	if err != nil {
		 return "", err
	}
	return str, nil
}
