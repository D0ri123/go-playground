package main

import "fmt"

var globalA = 5

/*
*
변수 선언 방식 => var [변수이름] [변수형] = [초기값]
별다른 형 선언 없이(func 내에서만 사용가능) => [변수이름] := [초기값]
*/
func main() {
	var a string = "goorm"
	fmt.Println(a)

	var b int = 10
	fmt.Println(b)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	fmt.Println(globalA)

	// 여러 개의 변수를 한 번에 선언하고 초기화할 수 있다.
	var num1, num2 int = 10, 20
	fmt.Println(num1, num2)

	i, j, k := 1, 2, 3
	fmt.Println(i, j, k)

	var str1, str2 string = "Hello", "goorm"
	fmt.Println(str1, str2)

}
