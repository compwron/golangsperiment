package json

import (
	"encoding/json" 
	"fmt"
	"github.com/bitly/go-simplejson" // pretty import arbitrary format json
	"os"
	"log"
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

func ParseArbitraryJsonUsingLibrary() (*simplejson.Json, error) {
	fmt.Println("\nUsing library to parse json without knowing what it will be:")
	// from https://github.com/bitly/go-simplejson/blob/master/simplejson_test.go
	arbitraryJson, err := simplejson.NewJson([]byte(`{
		"test": {
			"string_array": ["asdf", "ghjk", "zxcv"],
			"string_array_null": ["abc", null, "efg"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"subkeyone": 1},
			{"subkeytwo": 2, "subkeythree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true,
			"sub_obj": {"a": 1}
		}
	}`))
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arbitraryJson)
	return arbitraryJson, err
}

// func WriteJsonToFile(err){
	// json.NewEncoder(w io.Writer) *Encoder
	// return nil
// }
