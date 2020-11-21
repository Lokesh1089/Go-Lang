package main
import (
	"fmt"
	"encoding/json"
)
func main() {

	type BookPrice struct{
		 Enggmaths int
		 Thermodynamics int
		 StrengthofMaterial int

	}

	data := `{"Enggmaths": 895, "Thermodynamics": 750, "StrengthofMaterial": 1200 }`

	p2 := BookPrice{}
	 json.Unmarshal([]byte(data),&p2)

	fmt.Println(p2.Enggmaths )
}