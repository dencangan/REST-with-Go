/*
Encoding and decoding datatypes to and from JSON
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

// Social media 
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type User struct {
	Name string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Users struct {
	Users []User `json:"users"`
}

func main() {

	// Opening json
	jsonFile, err := os.Open("../REST-and-Go/test.json")

	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Successfully opened test.json")
	
	// Once done, close json file
	defer jsonFile.Close()
	

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}
