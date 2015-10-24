package printer_test

import (
	"testing"
	"github.com/thejsj/go-json/printer"
	"github.com/stretchr/testify/assert"
)

func TestPrintJSONNumbers(t *testing.T) {
	value, _ := printer.PrintJSON(5)
	assert.Equal(t, value, "5", "PrintJSON should parse integers correctly")

	value1, _ := printer.PrintJSON(5.1)
	assert.Equal(t, value1, "5.1", "PrintJSON should parse floats correctly")

	value2, _ := printer.PrintJSON(9999999)
	assert.Equal(t, value2, "9999999", "PrintJSON should parse bigger integers correctly")

	value3, _ := printer.PrintJSON(99.99999)
	assert.Equal(t, value3, "99.99999", "PrintJSON should parse bigger integers correctly")
}

func TestPrintJSONString(t *testing.T) {
	value, _ := printer.PrintJSON("hello")
	assert.Equal(t, value, "\"hello\"", "PrintJSON should strings correctly")

	value1, _ := printer.PrintJSON("\"wow\"")
	assert.Equal(t, value1, "\"\"wow\"\"", "PrintJSON should parse strings with strings correctly")
}

func TestPrintJSONArray(t *testing.T) {
	var arr [3]int
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arrString, _ := printer.PrintJSON(arr)
	assert.Equal(t, "[0,1,2]", arrString, "PrintJSON should parse arrays of integers correctly")

	var arr1 [3]string
	arr1[0] = "hello"
	arr1[1] = "world"
	arr1[2] = "wow wow wow"
	arrString1, _ := printer.PrintJSON(arr1)
	assert.Equal(t, "[\"hello\",\"world\",\"wow wow wow\"]", arrString1, "PrintJSON should parse arrays of strings correctly")

	var arr2 [3]float64
	arr2[0] = 2.67
	arr2[1] = 8.
	arr2[2] = 99999.1
	arrString2, _ := printer.PrintJSON(arr2)
	assert.Equal(t, "[2.67,8,99999.1]", arrString2, "PrintJSON should parse arrays of floats correctly")
}

func TestPrintJSONMap(t *testing.T) {
	map1 := make(map[string]string)
	map1["hello"] =  "world"
	mapString1, _ := printer.PrintJSON(map1)
	assert.Equal(t, "{\"hello\":\"world\"}", mapString1, "PrintJSON should parse string maps correctly")

	map2 := make(map[string]float64)
	map2["hello"] = 8.
	map2["world"] = 99.99
	mapString2, _ := printer.PrintJSON(map2)
	assert.Equal(t, "{\"hello\":8,\"world\":99.99}", mapString2, "PrintJSON should parse float maps correctly")

	map3 := make(map[string]int)
	map3["hello"] = 8
	map3["world"] = 12345678
	mapString3, _ := printer.PrintJSON(map3)
	assert.Equal(t, "{\"hello\":8,\"world\":12345678}", mapString3, "PrintJSON should parse int maps correctly")

	map4 := make(map[int]int)
	map4[1] = 8
	map4[2] = 12345678
	mapString4, _ := printer.PrintJSON(map4)
	assert.Equal(t, "{\"1\":8,\"2\":12345678}", mapString4, "PrintJSON should parse int keys as strings")

  type Node struct {
		Next  *Node
		Value interface{}
	}
	map5 := make(map[*Node]int)
	value, err := printer.PrintJSON(map5)
	assert.NotNil(t, err, "PrintJSON should throw an error when keys are not numbers or strings")
}
