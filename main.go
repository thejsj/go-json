package main;

import (
    "io/ioutil"
    "github.com/thejsj/go-json/parser"
    "log"
    "os"
)

func main() {
    bytes, err := ioutil.ReadAll(os.Stdin)
    inputString := string(bytes)

    if err != nil {
        log.Fatal(err)
    }

    parsedJSONString := parser.ParseJSON(inputString)
    log.Println(parsedJSONString)
}
