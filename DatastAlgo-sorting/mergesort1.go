package main
import (
    "fmt"
)
func merge(a, b []int) []int {
  size, i, j := len(a)+len(b), 0, 0
  result := make([]int, size)
  for k := 0; k < size; k++ {
    switch true {
      case i == len(a):
           result[k] = b[j]
          j++
      case j == len(b):
            result[k] = a[i]
        i++
      case a[i] > b[j]:
         result[k] = b[j]
             j++
      case a[i] < b[j]:
              result[k] = a[i]
          i++
      case a[i] == b[j]:
          result[k] = b[j]
          j++
                  }
      }
   return result
  }
func mergeSort(numbers []int) []int {
    if len(numbers) < 2 {
       return numbers
     }
	middle := int(len(numbers) / 2)
	r:=numbers[middle:]
	l :=numbers[:middle]
  return merge(mergeSort(l), mergeSort(r))
}
func main() {
 a := []int{4,8,5,1,6,9}
 fmt.Println(mergeSort(a))
}