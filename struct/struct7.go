
// tags

package main 

import("fmt"
	 "reflect"
)


type password struct{

	enterpassword int `"required min 8 digits"`
}

func main() {
c := reflect.TypeOf(password{ })
field, _ := c.FieldByName("enterpassword")
fmt.Println(field.Tag)
}