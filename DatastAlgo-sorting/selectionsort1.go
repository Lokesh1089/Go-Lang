package main
import(
	"fmt"
)
func main(){
  arr :=[]int{3,4,5,2,1}
  selsort(arr)
}

func selsort(a []int){
   
	len := len(a)

	for i:=0;i<len-1;i++{
		for j:=i+1;j<len;j++{
			if a[i]>a[j]{
				a[i],a[j]=a[j],a[i]
			}
		}
	}
fmt.Println("After sorting -> ", a)

}