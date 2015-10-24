package to_json_test

import (
	"testing"
	"github.com/thejsj/go-json/to_json"
	"github.com/stretchr/testify/assert"
)

func TestToJSONNumbers(t *testing.T) {
	value, _ := to_json.ToJSON(5)
	assert.Equal(t, value, "5", "ToJSON should parse integers correctly")

	value1, _ := to_json.ToJSON(5.1)
	assert.Equal(t, value1, "5.1", "ToJSON should parse floats correctly")

	value2, _ := to_json.ToJSON(9999999)
	assert.Equal(t, value2, "9999999", "ToJSON should parse bigger integers correctly")

	value3, _ := to_json.ToJSON(99.99999)
	assert.Equal(t, value3, "99.99999", "ToJSON should parse bigger integers correctly")

	var int64Value int64
	int64Value = 88888
	value4, _ := to_json.ToJSON(int64Value)
	assert.Equal(t, value4, "88888", "ToJSON should parse bigger integers correctly")

	var int8Value int8
	int8Value = 82
	value5, _ := to_json.ToJSON(int8Value)
	assert.Equal(t, value5, "82", "ToJSON should parse bigger integers correctly")
}

func TestToJSONString(t *testing.T) {
	value, _ := to_json.ToJSON("hello")
	assert.Equal(t, value, "\"hello\"", "ToJSON should strings correctly")

	value1, _ := to_json.ToJSON("\"wow\"")
	assert.Equal(t, value1, "\"\"wow\"\"", "ToJSON should parse strings with strings correctly")
}

func TestToJSONArray(t *testing.T) {
	var arr [3]int
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arrString, _ := to_json.ToJSON(arr)
	assert.Equal(t, "[0,1,2]", arrString, "ToJSON should parse arrays of integers correctly")

	var arr1 [3]string
	arr1[0] = "hello"
	arr1[1] = "world"
	arr1[2] = "wow wow wow"
	arrString1, _ := to_json.ToJSON(arr1)
	assert.Equal(t, "[\"hello\",\"world\",\"wow wow wow\"]", arrString1, "ToJSON should parse arrays of strings correctly")

	var arr2 [3]float64
	arr2[0] = 2.67
	arr2[1] = 8.
	arr2[2] = 99999.1
	arrString2, _ := to_json.ToJSON(arr2)
	assert.Equal(t, "[2.67,8,99999.1]", arrString2, "ToJSON should parse arrays of floats correctly")
}

func TestToJSONMap(t *testing.T) {
	map1 := make(map[string]string)
	map1["hello"] =  "world"
	mapString1, _ := to_json.ToJSON(map1)
	assert.Equal(t, "{\"hello\":\"world\"}", mapString1, "ToJSON should parse string maps correctly")

	map2 := make(map[string]float64)
	map2["hello"] = 8.
	map2["world"] = 99.99
	mapString2, _ := to_json.ToJSON(map2)
	assert.Equal(t, "{\"hello\":8,\"world\":99.99}", mapString2, "ToJSON should parse float maps correctly")

	map3 := make(map[string]int)
	map3["hello"] = 8
	map3["world"] = 12345678
	mapString3, _ := to_json.ToJSON(map3)
	assert.Equal(t, "{\"hello\":8,\"world\":12345678}", mapString3, "ToJSON should parse int maps correctly")

	map4 := make(map[int]int)
	map4[1] = 8
	map4[2] = 12345678
	mapString4, _ := to_json.ToJSON(map4)
	assert.Equal(t, "{\"1\":8,\"2\":12345678}", mapString4, "ToJSON should parse int keys as strings")

  type Node struct {
		Next  *Node
		Value interface{}
	}
	map5 := make(map[*Node]int)
	_, err := to_json.ToJSON(map5)
	assert.NotNil(t, err, "ToJSON should throw an error when keys are not numbers or strings")
}
