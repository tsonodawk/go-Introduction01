package main

import "fmt"

// type error interface {
// 	Error() string
// }

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println(err)
	// fmt.Println(err.Error())
	fmt.Println(err.(*MyError).ErrCode)
	// fmt.Println(err.(*MyError).Message)
	// fmt.Println(err.(*MyError))

}
