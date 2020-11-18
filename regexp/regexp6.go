package main
import (
	"fmt"
	//"bytes"
	"regexp"
)

func main() {
	str,_ := regexp.MatchString("l([a-z]+)nd", "legend")
	fmt.Println(str)

	r:= regexp.MustCompile("m([a-z]+)e")

	fmt.Println(r.MatchString("make"))
	fmt.Println(r.FindStringIndex("make made"))
	fmt.Println(r.FindStringSubmatch("make  cake  made "))
	fmt.Println(r.FindStringSubmatchIndex("make cake made"))
	fmt.Println(r.FindAllString("make made meditate", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("make made meditate", -1))
	fmt.Println(r.FindAllString("make made meditate", 2))
	fmt.Println(r.Match([]byte("make")))

	r = regexp.MustCompile("k([a-z]+)ne")
	fmt.Println(r)

}