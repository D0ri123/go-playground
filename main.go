package main

import "fmt"

const username = "lee"

/*
상수를 괄호로 묶어 선언하는 경우
값이 선언되지 않은 상수는 바로 전 상수의 값을 가진다.
iota라는 식별자를 값으로 초기화하면 그 다음 상수들은 순서가 값으로 저장된다.
*/
const (
	c1 = 10
	c2
	c3
	c4
	c5
	c6 = iota
	c7
	c8
	c9 = "earth"
	c10
	c11 = "End"
)

func main() {
	//상수 개별 선언
	const a int = 1
	const b, d = 10, 20
	const c = "goorm"

	fmt.Println(username)
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11)
}
