/*
Encoding and decoding datatypes to and from JSON
*/

package main

import (
	"encoding/json"
	"fmt"
	//"os"
)

type response1 struct {
	Page int
	Fruits []string
}

// Tags on struct field to customize endoded JSON key names
type response2 struct {
	Page int `json: "page"`
	Fruits []string `json: "fruits"`
}

func jsonEncoder(inputDatatype string) {
	
	output, _ := json.Marshal(inputDatatype)
	
	// Preview
	fmt.Println(string(output))
}

func main() {

	// Encoding basic data types to JSON strings
	bolB, _:= json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
    fmt.Println(string(intB))
    fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Array	
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	// Maps
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

	res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	}