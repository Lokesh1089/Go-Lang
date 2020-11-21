package main
import (
	"fmt"
	"encoding/json"
) 

type Box struct {

	Width int
	Height int
	Colour string
	Open bool
}

func main() {

	b := Box{
		Width: 15,
		Height: 30,
		Colour: "yellow",
		Open: false,
	}

	bj, err := json.Marshal(b)

	if err!= nil {

		fmt.Println(err)
	}

	str := string(bj)
	fmt.Println(str)

}