package main

import (
	"github.com/thejsj/go-json/json"
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
