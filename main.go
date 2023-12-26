package main

import "fmt"

func main() {
	var num1, num2, num3 int

	/*
		Scanln은 빈칸으로 값을 구분하고, 엔터를 누르면 입력 종료된다.
		입력받는 변수에 &연산자를 붙여 입력받는다.
	*/
	fmt.Print("정수 3개를 입력하세요 :")
	fmt.Scanln(&num1, &num2, &num3)
	fmt.Println(num1, num2, num3)
}
