// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	//var i int = 55
	//var f float32 = 3.99

	i := 55
	f := 3.99
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(i))
	fmt.Println(f, math.Ceil(3.49))
	fmt.Println(strings.Title("kim inha"))
	fmt.Printf("i는 %d \n", i)
	fmt.Print("i는 ", i, "\n")
	fmt.Println("i는", i)

}
