package json

import (
	"encoding/json" // known json format
	"fmt"
	// "github.com/bitly/go-simplejson" // arbitrary format json
	"os"
	// "log"
)

type Configuration struct {
	Users  []string
	Groups []string
	Port   int
}

func ParseKnownFormatJsonFromFile() (*Configuration, error) {
	file, _ := os.Open("json/conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("\nParsing json into known structure:")
	fmt.Println(configuration)
	return &configuration, nil
}

func ParseArbitraryJson() (map[string]interface {}, error) {
	someJson := []byte(`{"foo":"bar","baz":6,"stuff":["a","b"], "isTrue":false}`)
	var jsonHolder interface{}
	err := json.Unmarshal(someJson, &jsonHolder)

	parsedMap := jsonHolder.(map[string]interface{})
	fmt.Println("\nParsing json without knowing what it will be:")
	for key, value := range parsedMap {
		fmt.Println(key, value)
	}
	return parsedMap, err
}

// func ParseArbitraryJson() (*simplejson.Json, error) {
// 	// from https://github.com/bitly/go-simplejson/blob/master/simplejson_test.go
// 	file, _ := os.Open("json/example1.json")
// 	arbitraryJson, err := simplejson.NewFromReader(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for key, value := range arbitraryJson.Map() {
// 		fmt.Println(key, value)
// 	}

// 	// fmt.Printf("%v\n", arbitraryJson)
// 	fmt.Println(arbitraryJson)
// 	return arbitraryJson, err

// }
