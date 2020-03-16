/*
Encoding and decoding datatypes to and from JSON
*/

package main

import (
	"encoding/json"
	"fmt"
	//"os"
)


/*
Encodes data types into JSON format
	Bool: true
	Int: 1
	Float: 2.34
	String: "gopher"
	Array: []string{"apple", "peach", "pear"}
	Maps: map[string]int{"apple": 5, "lettuce": 7}
*/
func jsonEncoder(inputDatatype interface{}) {
	
	output, _ := json.Marshal(inputDatatype)
	
	// Preview
	fmt.Println(string(output))
	return
}

func main() {

	jsonEncoder([]string{"apple", "peach", "pear"})
}
