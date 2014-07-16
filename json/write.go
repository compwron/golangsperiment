package json

import (
	"encoding/json"
	"log"
	 "io/ioutil"
	 "fmt"
)

type Foo struct {
	Name string
	Body string
	Time int64
}

func WriteJsonToFile() {
	filename := "a_foo.json"
	fooBar := Foo{"Bar", "Hello", 1294706395881547000}
	fmt.Println("\nWriting json ", fooBar, "to file: ", filename)
	
	marshalledFooBar, err := json.Marshal(fooBar)
	err = ioutil.WriteFile(filename, marshalledFooBar, 0644)
	
	fmt.Println("\nRead json from file", filename)
    dat, err := ioutil.ReadFile(filename)
    fmt.Println(string(dat))

    if err != nil{
		log.Fatal(err)
	}
}
