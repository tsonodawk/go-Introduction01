package main

import "fmt"

// func PrintSliceInts(slice []int) {
// 	for _, v := range slice {
// 		fmt.Println(v)
// 	}
// }

// func PrintSliceStrings(slice []string) {
// 	for _, v := range slice {
// 		fmt.Println(v)
// 	}
// }

// func PrintSlice[T any](slice []T) {
// func PrintSlice[T int | string](slice []T) {
// 	for _, v := range slice {
// 		fmt.Println(v)
// 	}
// }

// func f[T fmt.Stringer](xs []T) []string {
// 	result := []string{}
// 	for _, x := range xs {
// 		result = append(result, x.String())
// 	}

// 	return result
// }

// type MyInt int

// func (i MyInt) String() string {
// 	return fmt.Sprintf("%d", i)
// }

// type Number interface {
// 	int | int32 | int64 | float32 | float64
// }

// func Max[T Number](x, y T) T {
// 	if x > y {
// 		return x
// 	}

// 	return y
// }

// type Vectore[T any] []T

// type IntVectore = Vectore[int]

// type T[A any, B []c, C *A] struct {
// 	a A
// 	b B
// 	c C
// }

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](xs ...T) Set[T] {
	s := make(Set[T])
	for _, x := range xs {
		s.Add(x)
	}

	return s

}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}

func main() {
	// PrintSliceInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// PrintSliceStrings([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})

	// PrintSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// PrintSlice[string]([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})
	// PrintSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// PrintSlice([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})

	// fmt.Println(f([]MyInt{1, 2, 3, 4, 5, 6, 7, 8, 9}))

	// fmt.Println(Max[int](1, 2))

	// var v Vectore[int] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(v)

	// var v2 Vectore[float64] = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	// fmt.Println(v2)

	// v3 := IntVectore{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(v3)

	// var t T[int, []int, *int]

	// var a any = 1
	// a = "a"
	// a = true

	s := NewSet(1, 2, 3)
	fmt.Println(s)

	s.Remove(1)
	fmt.Println(s)

}
