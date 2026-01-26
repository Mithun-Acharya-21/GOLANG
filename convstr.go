package main
import(
	"fmt"
	"strconv"
)
func convertString(){
	var i int=21
	fmt.Printf("%v,%T",i,i)

	var j string
	j=strconv.Itoa(i)
	fmt.Printf("%v,%T",j,j)
}