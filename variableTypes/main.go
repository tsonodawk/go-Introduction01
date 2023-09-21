package main

import (
	"fmt"
	"strconv"
)

func main() {
	// INT
	var i int = 100

	var i2 int64 = 200

	fmt.Println(i + 50)

	// fmt.Println(i2 + i)

	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", int32(i2))
	// fmt.Printf("%T\n", strconv.Itoa(int(i2)))

	// FLOAT
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	// BOOL
	var t, f bool = true, false
	fmt.Println(t, f)

	// STRING

	// BYTE

	// ARRAY
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// Change Type
	chnageType()
}

func chnageType() {
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fl64 = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3 = %T\n", i3)
	// fmt.Println(i3)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

}
