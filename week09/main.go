package main

import (
	"fmt"
	"math/rand"
)

func main() {
	result := fmt.Sprintf("%0.2f 나누기 %0.2f은 %0.2f입니다.\n", 1.0, 3.0, 1.0/3.0)
	fmt.Print(result)

	i := 1
	for i <= 10 {
		fmt.Printf("%4d\n", rand.Intn(1000)+1)
		i++
	}
}
