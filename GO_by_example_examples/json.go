/*
GO has built-in support for JSON encoding and decoding.
*/ 

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Using two structs to demonstrate encoding and decoding
// custom types.

type response1 struct {
	Page int
	Fruits []string
}

// Only exported fields will be encoded/decoded in JSON
// Fields must start with capital letters to be exported

type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Examples of slices and maps encoded to JSON arrays
	// and objects
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Json package can automatically encode custom data types
	res1D := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Use tags on struct field dec's to customize encoded
	// JSON key names
	res2D := &response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Decoding JSON data into Go values
	// A generic data structure
	byt := []byte(`{"num":6.13, "strs": ["a", "b"]}`)

	// Provide a var where the JSON package can put
	// the decoded data. map[string]interface{} will hold
	// a map of strings to arbitrary data types
	var dat map[string]interface{}

	// Decoding and checking for errors
	if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
	fmt.Println(dat)

	// To use val's in the decoded map, they must be converted
	// to the appropriate type. Converting num to the expected 
	// float64
	num := dat["num"].(float64)
	fmt.Println(num)

	// access nested data
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Decode JSON into custom data types.
	// Advantages are adding additional type-safety and
	// eliminating the need for type assertions
	str := `{:page": 1. "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// You can also stream JSON encodings directly to
	// os.Writers like os.Stdout or HTTP response bodies
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}