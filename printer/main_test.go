package printer_test

import (
	"testing"
	"github.com/thejsj/go-json/printer"
	"github.com/stretchr/testify/assert"
)

func TestPrintJSON(t *testing.T) {
	value, _ := printer.PrintJSON(5)
	assert.Equal(t, value, "5", "PrintJSON should parse integers correctly")

	value1, _ := printer.PrintJSON(5.1)
	assert.Equal(t, value1, "5.1", "PrintJSON should parse floats correctly")

	value2, _ := printer.PrintJSON(9999999)
	assert.Equal(t, value2, "9999999", "PrintJSON should parse bigger integers correctly")
}
