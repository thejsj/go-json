package main

import (
	"fmt"
	"github.com/thejsj/go-json/json"
	"io/ioutil"
	"log"
	"os"
)


var isArray func(str string) bool
var ToJSON func(parsedJSON interface{}) (string, error)
var ParseJSON func(str string) (interface{}, error)

func init () {
	ToJSON = json.ToJSON
	ParseJSON = json.ParseJSON
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	inputString := string(bytes)

	if err != nil {
		log.Fatal(err)
	}

	parsedJSON, err := json.ParseJSON(inputString)
	if err != nil {
		log.Fatal(err)
	}
	str, err := json.ToJSON(parsedJSON)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
