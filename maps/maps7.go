package main

import "fmt"

func main () {

  veggie := map[string]int {
	"tomato" : 35,
	"onion" : 200,
	"brinjal" : 45,
	"potato": 70,
	"carrot": 90,
  }

  fmt.Println(veggie)

  //ve, ok := veggie["beans"]

  //fmt.Println( ve, ok)

  ve, ok:= veggie["potato"]

  fmt.Println(ve, ok)

}