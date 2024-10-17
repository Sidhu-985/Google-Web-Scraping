package main

import (
	"errors"
	"fmt"
)

func main(){

	var m int = 3
	var n int = 4

	var result ,remainder,error = divide(m,n)

	if (error!=nil){
		fmt.Printf(error.Error())
	}

	fmt.Printf("The result of integer after dividing is %v and remainder = %v",result,remainder)
}

func divide(num1 int,num2 int) (int,int,error){
	var result int = num1/num2
	var err error

	if (num2==0){
		err = errors.New("divided by zero not possible")
	}

	var remainder int = num1%num2
	return result,remainder,err
}