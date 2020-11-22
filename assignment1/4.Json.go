package main
import (
	"fmt"
	"encoding/json")

type StudentDetails struct{
	Name string `json:"name"`
	Age int
	Major string `json:"major"`
	IsPass bool

}	

func main() {

  stuOne:= StudentDetails{
	  Name:"Santhosh kumar", 
	  Age: 17,
	  Major:"Computer Science",
	  IsPass: true ,
  }

//encoding
  js, err := json.Marshal(stuOne)
  if err!= nil {
	  fmt.Println(err) 
	 }

	 s := string(js)

	 fmt.Println(s)

// decoding

	 jsonData := `{"name":"Santhosh kumar","Age":17,"major":"Computer Science","IsPass":true}`

	 var sd StudentDetails
	 json.Unmarshal([]byte(jsonData), &sd)

	 fmt.Println(sd)


}