package main

import (
	"github.com/thejsj/go-json/parser"
	"io/ioutil"
	"log"
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
	log.Println(parsedJSON)

}
