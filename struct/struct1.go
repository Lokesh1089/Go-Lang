package main

import "fmt"

type Doctor struct {
	
	number int
	actorName string
	companions []string
}

func main (){

	 aDoctor := Doctor{

			number: 3,
			actorName: "vijay kumar",
			companions: []string {
					"surya",	
					 "madhavan",
					  "ashwin",
					},
				 }

				
				fmt.Println( aDoctor)

				fmt.Println( aDoctor.actorName)
				fmt.Println(aDoctor.companions[2])

			}