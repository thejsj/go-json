package main

import (
	"github.com/thejsj/go-json/parser"
	"github.com/thejsj/go-json/printer"
	"io/ioutil"
	"log"
	"fmt"
	"os"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	inputString := string(bytes)

	if err != nil {
		log.Fatal(err)
	}

	parsedJSON, err := parser.ParseJSON(inputString)
	if err != nil {
		log.Fatal(err)
	}
	str, err := printer.PrintJSON(parsedJSON)
	if err != nil {
		 log.Fatal(err)
	}
	fmt.Println(str)
}
