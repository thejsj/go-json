package main

import (
	"github.com/thejsj/go-json/parse_json"
	"github.com/thejsj/go-json/to_json"
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

	parsedJSON, err := parse_json.ParseJSON(inputString)
	if err != nil {
		log.Fatal(err)
	}
	str, err := to_json.ToJSON(parsedJSON)
	if err != nil {
		 log.Fatal(err)
	}
	fmt.Println(str)
}
