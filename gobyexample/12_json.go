package gobyexample

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type jsonResponse1 struct {
	Page   int
	Fruits []string
}

type jsonResponse2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type jsonResponse3 struct {
	Status int
	Body   string
}

func ShowJSON() {
	// note: marshalling means converting Go values to JSON

	// converting basic data types to JSON ([]byte)
	bolZ, _ := json.Marshal(true)
	fmt.Println(string(bolZ))

	intZ, _ := json.Marshal(91)
	fmt.Println(string(intZ))

	fltZ, _ := json.Marshal(11.92)
	fmt.Println(string(fltZ))

	strZ, _ := json.Marshal("gopher")
	fmt.Println(string(strZ))

	slcA := []string{"red", "blue", "green", "purple", "yellow"}
	slcZ, _ := json.Marshal(slcA)
	fmt.Println(string(slcZ))

	mapA := map[string]int{"Car": 5, "Plane": 7, "Boat": 3}
	mapZ, _ := json.Marshal(mapA)
	fmt.Println(string(mapZ))

	// converting custom data types to JSO
	res1D := &jsonResponse1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1Z, _ := json.Marshal(res1D)
	fmt.Println(string(res1Z))

	// custom data types with JSON tags
	res2D := &jsonResponse2{
		Page:   1,
		Fruits: []string{"orange", "grape", "banana"},
	}
	res2Z, _ := json.Marshal(res2D)
	fmt.Println(string(res2Z))

	// unmarshalling means converting JSON data to Go values

	// JSON to Go data types
	byt := []byte(`{"num":61.13,"strs":["pq","qp"]}`)
	var dat map[string]any
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// accessing values from the map
	strs := dat["strs"].([]any)
	str1 := strs[0].(string)
	str2 := strs[1].(string)
	fmt.Println(str1, str2)

	// JSON to custom data types
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := jsonResponse2{}
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// using encoder and decoder
	encoder := json.NewEncoder(os.Stdout)
	d := map[string]int{"Car": 5, "Plane": 7, "Boat": 3}
	encoder.Encode(d)

	dec := json.NewDecoder(strings.NewReader(str))
	res1 := jsonResponse2{}
	_ = dec.Decode(&res1)
	fmt.Println(res1)
}
