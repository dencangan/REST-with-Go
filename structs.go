/* Working with structs 

Structs are typed collections of fields. They're useful for grouping data together to form records.
*/

package main

import "fmt"

type fruit struct {
	name string
	count int
}

func chosenFruit(fruitName string) *fruit {

	f := fruit{name: fruitName}
	f.count = 1
	return &f
}

func main() {
	fmt.Println(fruit{name: "Orange", count: 2})

	f := fruit{}
	f.count = 1
	f.name = "Apple"

	fmt.Println(f)
}