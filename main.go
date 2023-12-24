package main

import "fmt"

func main() {
	var num1 int = 1
	var num2 int = 2

	//fmt 패키지를 사용하지 않는 경우
	print("Hello goorm!")
	print(num1)
	print(num2)
	print(num1 + num2)
	print("Hello goorm!", num1+num2, "\n")

	println("Hello goorm!")
	println(num1)
	println(num2)
	println(num1 + num2)
	println("Hello goorm!", num1+num2)

	//fmt 패키지를 사용하는 경우
	fmt.Print("Hello goorm!", num1, num2, "\n")

	fmt.Println("Hello goorm!", num1, num2)

	fmt.Printf("num1의 값은:%d num2의 값은:%d\n", num1, num2)
}
