// decoding

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//FruitBasket int value for Id
type FruitBasket struct {

	Name string
	Fruit [] string
	Id int
	Private bool
	Created time.Time
}

func main(){

	basket := FruitBasket{
		Name: "Standard",
		Fruit: []string {"apple", "orange", "banana"},
		Id: 197,
		Private: true ,
		Created: time.Now(),

	}

	var flist []byte
	flist, err := json.Marshal(basket)
	 if err != nil {
		 fmt.Println(err)
	 }

	 fmt.Println(string(flist))


	 data1 := `{"Name":"Standard","Fruit":["apple","orange","banana"],"Id":197,"Private":true,"Created":"2020-11-18T22:18:27.4290215+05:30"}`

		 var b FruitBasket
		 json.Unmarshal([]byte(data1), &b )
		 
		 fmt.Println( b.Fruit )
		 fmt.Println(b)
		 fmt.Println( b.Id)


	}