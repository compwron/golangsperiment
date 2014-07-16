package json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Foo struct {
	Name string
	Body string
	Time int64
}

func WriteJsonToFile() {
	fooBar := Foo{"Bar", "Hello", 1294706395881547000}
	marshalledFooBar, err := json.Marshal(fooBar)
	err := ioutil.WriteFile("a_foo.json", marshalledFooBar, 0644)
	if err != nil{
		log.Fatal(err)
	}
}
