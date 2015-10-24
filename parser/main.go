package parser

import (
	"errors"
	"strconv"
	"strings"
)

// TODO Change order of functions
// TODO Add init function and vars at the top

func firstAndLastChars(first string, last string) func(str string) bool {
	return func(str string) bool {
		return string(str[0]) == first && str[len(str)-1:] == last
	}
}

func splitByChar(baseChar string) func(str string) []string {
	return func(str string) []string {
		var result []string
		doubleStringOpen := false
		singleStringOpen := false
		arrayOpen := false
		objectOpen := false
		arrayBracketCount := 0
		objectBracketCount := 0
		currStr := ""
		for i := 0; i < len(str); i++ {
			ch := string(str[i])
			if ch == "\"" {
				doubleStringOpen = !doubleStringOpen
			}
			if ch == "'" {
				singleStringOpen = !singleStringOpen
			}
			if ch == "[" {
				arrayBracketCount += 1
				arrayOpen = true
			}
			if ch == "]" {
				arrayBracketCount -= 1
				if arrayBracketCount == 0 {
					arrayOpen = false
				}
			}
			if ch == "{" {
				objectBracketCount += 1
				objectOpen = true
			}
			if ch == "}" {
				objectBracketCount -= 1
				if objectBracketCount == 0 {
					objectOpen = false
				}
			}
			if ch == baseChar &&
				!doubleStringOpen &&
				!singleStringOpen &&
				!arrayOpen &&
				!objectOpen {
				if currStr != "" {
					appendedStr := strings.TrimSpace(currStr)
					result = append(result, appendedStr)
				}
				currStr = ""
			} else {
				currStr += ch
			}
		}
		if currStr != "" {
			appendedStr := strings.TrimSpace(currStr)
			result = append(result, appendedStr)
		}
		return result
	}
}

func ParseJSON(str string) (interface{}, error) {
	isArray := firstAndLastChars("[", "]")
	isObj := firstAndLastChars("{", "}")
	hasDoubleQuotes := firstAndLastChars("\"", "\"")
	hasSingleQuotes := firstAndLastChars("''", "''")
	isString := func(str string) bool {
		// Do I need to only trim space or is there something else
		// that need to get trimmed?
		str = strings.TrimSpace(str)
		return (hasSingleQuotes(str) || hasDoubleQuotes(str)) && str[len(str)-2:] != "\\"
	}
	isNumber := func(str string) bool {
		floatValue, floatErr := strconv.ParseFloat(str, 64)
		intValue, intErr := strconv.ParseInt(str, 10, 64)
		if floatErr != nil && intErr != nil {
			return false
		}
		floatString := strconv.FormatFloat(floatValue, 'f', 6, 64)
		intString := strconv.FormatInt(intValue, 16)
		return floatString == str || intString == str
	}
	removeFirstAndLastChar := func(str string) string {
		// Do I need to only trim space or is there something else
		// that need to get trimmed?
		str = strings.TrimSpace(str)
		return str[1 : len(str)-1]
	}
	separateStringByCommans := splitByChar(",")
	separateStringByColons := splitByChar(":")

	str = strings.TrimSpace(str)
	if isArray(str) {
		var arrParts []interface{}
		strParts := separateStringByCommans(removeFirstAndLastChar(str))
		for i := 0; i < len(strParts); i++ {
			newPart, err := ParseJSON(strParts[i])
			if err != nil {
				return nil, err
			}
			arrParts = append(arrParts, newPart)
		}
		return strParts, nil
	}
	if isObj(str) {
		var obj map[string]interface{}
		obj = make(map[string]interface{})
		objParts := separateStringByCommans(removeFirstAndLastChar(str))
		for i := 0; i < len(objParts); i++ {
			subObject := objParts[i]
			keyValuePair := separateStringByColons(subObject)
			if len(keyValuePair) == 2 {
				key, keyErr := ParseJSON(keyValuePair[0])
				if keyErr != nil {
					return nil, keyErr
				}
				value, valueErr := ParseJSON(keyValuePair[1])
				if valueErr != nil {
					return nil, valueErr
				}
				keyString, isString := key.(string)
				if isString {
					obj[keyString] = value
				}
			}
		}
		return obj, nil
	}
	if isString(str) {
		return removeFirstAndLastChar(str), nil
	}
	if isNumber(str) {
		floatValue, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		return floatValue, nil
	}
	if str == "nil" {
		return nil, nil
	}
	if str == "false" {
		return false, nil
	}
	if str == "true" {
		return true, nil
	}
	return nil, errors.New("Type not recognized for " + string(str))
}
