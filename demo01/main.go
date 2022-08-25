package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {

	fmt.Println(runtime.Version())
	x := 3.01

	a := int(math.Ceil(x))
	fmt.Println(a)

}
