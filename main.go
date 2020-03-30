package main

import "fmt"

func main() {
	// colors := map[string]string{ // One of the ways to declare a map. First string between the square brackets indicates that the keys are of type string. And the next occurrence of string says all the values are string as well
	// 	"red":   "#ff0000",
	// 	"green": "#008000",
	// }

	// var colors map[string]string // Another way of declaring a map. Since we didn't assign an actual value, Go will initialize it with its zero value. I.e. an empty map - has no key value pairs inside of it

	colors := make(map[string]string) // Yet another way of declaring a map, using Go's built-in function. This line here and the previous way of declaring a map are pretty much equivalent for all intents and purposes

	colors["white"] = "#ffffff" // To add a key value pair

	delete(colors, "white") // To delete a key value pair, use the built in function 'delete'

	fmt.Println(colors)
}

// Notes:
// 1. In Go, a map is a collection of key-value pairs
//    Map is Go is like - Hash in Ruby - Object in JavaScript - Dict in Python
//    - In maps, both the keys and the values are statically typed
//      So whenever we add some number of keys to a map in Go, they must all be of the same exact type.
//      Likewise, all the different values that we add as well must also be of the exact same type.
//      Just the keys and the values themselves don't have to be of the same type just all the different values have to.
//      I.e. key1, key2, key3 must be all of the same type
//           and value1, value2, value3 must be all of the same type
