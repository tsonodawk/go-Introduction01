package main

import "fmt"

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	// slice()

	// sliceAppendMakeLenCap()

	// sliceCopy()

	// sliceFor()

	// sliceVariableLengthArgument()

	// sliceMap()

	// sliceMapFor()

	// sliceChannel()

	// sliceChannelSelect()

	slicePointer()
}

func Double(i int) {
	i = i * 2
}
func Doublev2(i *int) {
	*i = *i * 2
}
func Doublev3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func slicePointer() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(&n)

	Double(n)
	fmt.Println(n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	// *p = 300
	// fmt.Println(n)

	// n = 200
	// fmt.Println(*p)

	Doublev2(&n)
	fmt.Println(n)

	Doublev2(p)
	fmt.Println(*p)

	var sl []int = []int{1, 2, 3}
	Doublev3(sl)
	fmt.Println(sl)

}

func sliceChannelSelect() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch2 <- "A"
	ch1 <- 1
	ch2 <- "B"
	ch1 <- 2

	// v1 := <-ch1
	// v2 := <-ch2

	// fmt.Println(v1)
	// fmt.Println(v2)

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!!!")
	default:
		fmt.Println("default")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0

L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("received", i3)
		default:
			// fmt.Println("default")
			if n > 30 {
				break L
			}
		}
	}
}

func sliceChannel() {
	ch1 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	for i := range ch1 {
		fmt.Println(i)
	}

	// ch := make(chan int, 2)
	// ch <- 100
	// ch <- 200

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// ch <- 300
	// fmt.Println(<-ch)

}

func sliceMapFor() {

	m := map[string]int{"Apple": 100, "Banaa": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

}

func sliceMap() {
	var m = map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Println(m)

	m2 := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3])

	s, ok := m4[3]

	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s, ok)

	m4[2] = "US"
	m4[3] = "UK"
	fmt.Println(m4)

	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))

}

func sliceVariableLengthArgument() {

	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(Sum())

	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))

}

func sliceFor() {

	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// for _, v := range sl {
	// 	fmt.Println(v)
	// }

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

}

func sliceCopy() {
	// sl := []int{100, 200}
	// sl2 := sl

	// sl2[0] = 1000
	// fmt.Println(sl)

	// var i int = 10
	// i2 := i
	// i2 = 100
	// fmt.Println(i, i2)

	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(sl)

	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)

	n := copy(sl2, sl)

	fmt.Println(n, sl2)

	sl3 := make([]int, len(sl))

	n2 := copy(sl3, sl)
	fmt.Println(n2, sl3)

}

func sliceAppendMakeLenCap() {
	sl := []int{100, 200}
	fmt.Println(sl)

	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600, 700)
	fmt.Println(sl)

	sl2 := make([]int, 5)
	fmt.Println(sl2)

	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))

	sl3 := make([]int, 5, 10)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

}

func slice() {
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:2])
	fmt.Println(sl5[2:])
	fmt.Println(sl5[:])
	fmt.Println(sl5[len(sl5)-1])
	fmt.Println(sl5[1 : len(sl5)-1])

}
