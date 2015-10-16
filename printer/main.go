package printer

import (
	"reflect"
	"fmt"
)

func PrintJSON(parsedJSON interface{}) string {
	value := reflect.ValueOf(parsedJSON)
	if (value.Kind() == reflect.Map) {
		for i := 0; i < value.Len(); i++ {
			fmt.Println("i")
			fmt.Println(i)
		}
	}
	return "";
}
