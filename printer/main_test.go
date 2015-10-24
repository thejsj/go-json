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
	assert.Equal(t, "[0,1,2]", arrString, "PrintJSON should parse arrays of integeres correctly")
}
