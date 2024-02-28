package main

import (
	"errors"
	"fmt"
)

// :-
/**
Concepts:-
1. error interface:- There is built in type in golang called error. It if of the interface type
   type error interface {
		Error() string
   }

2. How to create errors ?
   There are two ways to create errors ->
	a) using errors package's New() method ->   err :=  errors.New("Error happened")
    b) using fmt Package printf() method   ->   err :=  errors.Errorf("Error happened")

3. How to create custom errors?
   Create any struct with properties that you want.
   Then implement the 'error' interface
   For eg:-
   type NetworkError struct {
   		code int
        msg  string
   }

	func (e NetworkError) Error() string {
       return fmt.Sprintf("network error(%d): %s", e.code, e.msg)
    }


4. How to create named errors?


5. How to wrap errors?
*/

type NetworkError struct {
	Code int
	Msg  string
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("network error(%d): %s", e.Code, e.Msg)
}

func customErrorUtil(name string) (bool, error) {
	userSlice := []string{"billa", "chuha", "bail", "kutta", "saanp", "bichchu"}
	isPresent := false
	for _, any_name := range userSlice {
		if name == any_name {
			isPresent = true
		}
	}
	if !isPresent {
		return false, NetworkError{
			Code: 404,
			Msg:  "user ois not present in the database",
		}
	} else {
		return true, nil
	}
}

func CustomErrorDemo() {
	isPresent, err := customErrorUtil("mahabilla")
	if err != nil {
		var neterr NetworkError
		if errors.As(err, &neterr) {
			fmt.Sprintf("error happened with code %d and msg %s", neterr.Code, neterr.Msg)
		}
		return
	}
	fmt.Println("User is present in the database ?? -> ", isPresent)
}

func Divide(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return dividend / divisor, nil
}

func BasicErrorDemo() {
	num, err := Divide(8, 0)
	if err != nil {
		_ = fmt.Errorf("Error happened -> %s ", err)
		return
	}
	fmt.Println("ans -> %d", num)
}

func main() {
	fmt.Println("Introduction to the concept of errors in golang")
	CustomErrorDemo()
}
