package main

import "fmt"

func main() {
	var num int = 10
	var changef float32 = float32(num)

	changei := int8(num)

	var str string = "goorm"
	changestr := []byte(str)
	str2 := string(changestr)

	fmt.Println(num)
	fmt.Println(changef, changei)

	fmt.Println(str)
	fmt.Println(changef, changei)

	fmt.Println(str)
	fmt.Println(changestr)
	fmt.Println(str2)

	//연산의 결괏값과 피연산자의 자료형이 일치하지 않아 컴파일 에러 발생
	//var num1, num2 int = 3, 4
	//var result float32 = num1 / num2
	//
	//fmt.Prinft("%f", result)
}
