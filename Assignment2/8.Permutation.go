package main
import "fmt"

func main(){

	  a := "ABC"
	  runeA := []rune(a)
	  permutation(runeA, 0, len(runeA)-1)
}

func permutation(runeB []rune, start , end int){
	if start == end {
		fmt.Println(string(runeB))
	} else {
		for i:=start;i<=end;i++{
			runeB[start],runeB[i]= runeB[i],runeB[start]
			permutation(runeB, start+1,end)
			runeB[start],runeB[i] = runeB[i],runeB[start]
			}
	}


}