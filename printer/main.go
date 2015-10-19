package printer

import (
	"reflect"
	"strings"
	"errors"
)

// Code taken from: https://gist.github.com/hvoecking/10772475
func convertJSONToString(copy, original reflect.Value) (string, error) {
	switch original.Kind() {
	// The first cases handle nested structures and translate them recursively

	// If it is a pointer we need to unwrap and call once again
	case reflect.Ptr:
		// To get the actual value of the original we have to call Elem()
		// At the same time this unwraps the pointer so we don't end up in
		// an infinite recursion
		originalValue := original.Elem()
		// Check if the pointer is nil
		if !originalValue.IsValid() {
			err := errors.New("emit macho dwarf: elf header corrupted")
			return "", err
		}
		// Allocate a new object and set the pointer to it
		copy.Set(reflect.New(originalValue.Type()))
		// Unwrap the newly created pointer
		return convertJSONToString(copy.Elem(), originalValue)

	// If it is an interface (which is very similar to a pointer), do basically the
	// same as for the pointer. Though a pointer is not the same as an interface so
	// note that we have to call Elem() after creating a new object because otherwise
	// we would end up with an actual pointer
	case reflect.Interface:
		// Get rid of the wrapping interface
		originalValue := original.Elem()
		// Create a new object. Now new gives us a pointer, but we want the value it
		// points to, so we have to call Elem() to unwrap it
		copyValue := reflect.New(originalValue.Type()).Elem()
		convertJSONToString(copyValue, originalValue)
		copy.Set(copyValue)
		return convertJSONToString(copy.Elem(), originalValue)

	// If it is a struct we translate each field
	case reflect.Struct:
		var allElements []string
		for i := 0; i < original.NumField(); i += 1 {
		  str, err := convertJSONToString(copy.Field(i), original.Field(i))
			if err != nil {
				return "", err
			}
			allElements = append(allElements, str)
		}
		return strings.Join(allElements[:],","), nil

	// If it is a slice we create a new slice and translate each element
	case reflect.Slice:
		var allElements []string
		copy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i += 1 {
		  str, err := convertJSONToString(copy.Index(i), original.Index(i))
			if err != nil {
				return "", err
			}
			allElements = append(allElements, str)
		}
		return "[" + strings.Join(allElements[:],",") + "]", nil

	// If it is a map we create a new map and translate each value
	case reflect.Map:
		var allElements []string
		copy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			// New gives us a pointer, but again we want the value
			copyValue := reflect.New(originalValue.Type()).Elem()
			keyStr, keyErr := convertJSONToString(copyValue, originalValue)
			valueStr, valueErr := convertJSONToString(copyValue, originalValue)
			if valueErr != nil {
				return "", valueErr
			}
			if keyErr != nil {
				return "", keyErr
			}
			allElements = append(allElements, keyStr + ":" + valueStr)
		}
		return "{" + strings.Join(allElements[:],",") + "}", nil

	// Otherwise we cannot traverse anywhere so this finishes the the recursion

	// If it is a string translate it (yay finally we're doing what we came for)
	case reflect.String:
		translatedString := original.Interface().(string)
		return "\"" + translatedString + "\"", nil

	// And everything else will simply be taken from the original
	default:
		copy.Set(original)
		return "--default--", nil
	}
	return "", nil

}

func PrintJSON(parsedJSON interface{}) (string, error){
	// Wrap the original in a reflect.Value
	original := reflect.ValueOf(parsedJSON)

	copy := reflect.New(original.Type()).Elem()
	str, err := convertJSONToString(copy, original)
	if err != nil {
		 return "", err
	}
	return str, nil
}
