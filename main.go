package main;

import (
    "io/ioutil"
    "strings"
    "strconv"
    "log"
    "os"
)

func FirstAndLastChars (first string, last string) func (str string) bool {
    return func (str string) bool {
        return string(str[0]) == first && str[len(str)-1:] == last
    }
}

func SplitByChar (baseChar string) func (str string) []string {
    return func (str string) []string {
        var result []string;
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
            if
                ch == baseChar &&
                !doubleStringOpen &&
                !singleStringOpen &&
                !arrayOpen &&
                !objectOpen {
                if currStr != ""  {
                    appendedStr := strings.TrimSpace(currStr)
                    result = append(result, appendedStr)
                }
                currStr = "";
            } else {
              currStr += ch;
            }
        }
        if currStr != "" {
            appendedStr := strings.TrimSpace(currStr)
            result = append(result, appendedStr)
        }
        return result
    }
}

func ParseJSON (str string) interface{} {
    IsArray := FirstAndLastChars("[", "]")
    IsObj := FirstAndLastChars("{", "}")
    HasDoubleQuotes := FirstAndLastChars("\"", "\"")
    HasSingleQuotes := FirstAndLastChars("''", "''")
    IsString := func (str string) bool {
        // Do I need to only trim space or is there something else
        // that need to get trimmed?
        str = strings.TrimSpace(str)
        return (HasSingleQuotes(str) || HasDoubleQuotes(str)) && str[len(str)-2:] != "\\";
    }
    IsNumber := func (str string) bool {
        floatValue, err := strconv.ParseFloat(str, 64)
        if err != nil {
            log.Println(err)
        }
        floatString := strconv.FormatFloat(floatValue, 'f', 6, 64)
        return floatString == str
    }
    RemoveFirstAndLastChar := func (str string) string {
        // Do I need to only trim space or is there something else
        // that need to get trimmed?
        str = strings.TrimSpace(str)
        return str[1:len(str) - 1]
    }
    SeparateStringByCommans := SplitByChar(",")
    SeparateStringByColons := SplitByChar(":")

    str = strings.TrimSpace(str)
    if IsArray (str) {
        var arrParts []interface{}
        strParts := SeparateStringByCommans(RemoveFirstAndLastChar(str))
        for i := 0; i < len(strParts); i++ {
            arrParts = append(arrParts, ParseJSON(strParts[i]))
        }
        return strParts
    }
    if IsObj(str) {
        // TODO: Handle objects
        var obj map[string]interface{}
        objParts := SeparateStringByCommans(RemoveFirstAndLastChar(str))
        for i := 0; i < len(objParts); i++ {
            subObject := objParts[i]
            keyValuePair := SeparateStringByColons(subObject)
            if len(keyValuePair) == 2 {
                key := ParseJSON(keyValuePair[0])
                value := ParseJSON(keyValuePair[1])
                keyString, isString := key.(string)
                if isString {
                    // NOTE: Throws error `panic: assignment to entry in nil map`
                    obj[keyString] = value
                }
            }
        }
        return obj
    }
    if IsString (str) {
        return RemoveFirstAndLastChar(str)
    }
    if IsNumber (str) {
        floatValue, err := strconv.ParseFloat(str, 64)
        if err != nil {
            // TODO: Proper error handling
             log.Println(err)
        }
        return floatValue
    }
    if str == "nil" {
        return nil
    }
    if str == "false" {
        return false
    }
    if str == "true" {
        return true
    }
    // TODO: Throw error
    return false
}

func main() {
    bytes, err := ioutil.ReadAll(os.Stdin)
    inputString := string(bytes)

    if err != nil {
        log.Fatal(err)
    }

    parsedJSONString := ParseJSON(inputString)
    log.Println(parsedJSONString)
}
