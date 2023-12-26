package main

import "fmt"

func main() {
	//이항 연산자 사용(java와 거의 유사함)
	num1, num2 := 17, 5
	str1, str2 := "Hello", "goorm!"

	fmt.Println("num1 + num2 =", num1+num2)
	fmt.Println("str1 + str2 =", str1+str2)
	fmt.Println("num1 - num2 =", num1-num2)
	fmt.Println("num1 * num2 =", num1*num2)
	fmt.Println("num1 / num2 =", num1/num2)
	fmt.Println("num1 % num2 =", num1%num2)

	/*
		증감 연산자 사용
		- 증감 연산자 사용 동시에 대입할 수 없다.
		- 전위 연산을 할 수 없다.
	*/
	count1, count2 := 1, 10.4
	count1++
	count2--

	fmt.Println("count1++ :", count1)
	fmt.Println("count2-- :", count2)

	//할당 연산자
	a := 2
	var num int

	num = a
	fmt.Println("num = a :", num)

	num += 4
	fmt.Println("num += 4 :", num)

	num -= 2
	fmt.Println("num -= 2 :", num)

	num *= 5
	fmt.Println("num *= 5 :", num)

	num /= 2
	fmt.Println("num /= 2 :", num)

	num %= 3
	fmt.Println("num %= 3 :", num)

	num = 3
	num &= 2
	fmt.Printf("num &= 2 : %08b, %d\n", num, num)

	num |= 5
	fmt.Printf("num |= 5 : %08b, %d\n", num, num)

	num ^= 4
	fmt.Printf("num ^= 4 : %08b, %d\n", num, num)

	num &^= 2
	fmt.Printf("num &^ = 2 : %08b, %d\n", num, num)

	num <<= 9
	fmt.Printf("num <<= 9 : %08b, %d\n", num, num)

	num >>= 8
	fmt.Printf("num >>= 8 : %08b, %d\n", num, num)

	//논리 연산자
	var c bool = true
	d := false

	fmt.Println("0 && 0 : ", d && d)
	fmt.Println("0 && 1 : ", d && c)
	fmt.Println("1 && 1 : ", c && c)
	fmt.Println("0 || 0 : ", d || d)
	fmt.Println("0 || 1 : ", d || c)
	fmt.Println("1 || 1 : ", c || c)

	fmt.Println("!1 ", !true)
	fmt.Println("!0 ", !false)

	//관계 연산자 사용
	fmt.Println("13 == 13 : ", 13 == 13)
	fmt.Println("13 == 23 : ", 13 == 23)
	fmt.Println("13 != 13 : ", 13 != 13)
	fmt.Println("3 != 5 : ", 3 != 5)
	fmt.Println("0 < 1 : ", 0 < 1)
	fmt.Println("0 > 1 : ", 0 > 1)
	fmt.Println("0 >= 1 : ", 0 >= 1)
	fmt.Println("0 <= 1 : ", 0 <= 1)

	//비트 연산자 사용
	num3 := 15
	num4 := 20

	fmt.Printf("num1 & num2 : %08b, %d\n", num3&num4, num3&num4)
	fmt.Printf("num1 | num2 : %08b, %d\n", num3|num4, num3|num4)
	fmt.Printf("num1 ^ num2 : %08b, %d\n", num3^num4, num3^num4)
	fmt.Printf("num1 &^ num2 : %08b, %d\n", num3&^num4, num3&^num4)

	fmt.Printf("num1 << 4 : %08b, %d\n", num3<<4, num3<<4)
	fmt.Printf("num2 >> 2 : %08b, %d\n", num4>>2, num4>>2)

	//채널 연산자 - 채널을 사용할 때 쓰는 연산자
	ch := make(chan int) // 정수형 채널 생성

	go func() {
		ch <- 10
	}()

	result := <-ch
	fmt.Println("result :", result)

	/*
		포인터 연산자 - &와 *연산자를 이용해 메모리 접근 가능
		포인터 산술 기능(포인터에 더하고 빼는 기능)은 제공하지 않음
		포인터 변수의 기본값은 nil
	*/
	var number int = 5
	var pnum = &num

	fmt.Println("number :", number) //number 값
	fmt.Println("pnum :", pnum)     //number의 메모리 주소
	fmt.Println("pnum :", *pnum)    //number의 주소로 메모리에 할당된 값 접근

	*pnum++
	fmt.Println("number :", number)
	fmt.Println("pnum :", *pnum)
}
