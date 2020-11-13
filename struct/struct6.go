
// tags

package main 

import (
	"fmt"
	"reflect"
)

type Customer struct{

	name string `"required max:100"`
	location string
}

func main() {

	t := reflect.TypeOf(Customer{ })
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)


}