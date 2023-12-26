package main

import "fmt"

func main() {
	//백틱 안에 있으면 개행문자열(\n)이 특별한 의미로 인식되지 않음
	//백틱은 [option + ~]를 누르면 됨
	var rawLiteral string = `바로 실행해보면서 배우는 \n Golang`

	//아무리 길게 치고, 엔터를 쳐도 한 줄로 표현된다. 이 때 \n이 특별한 역할을 함
	var interLiteral string = "바로 실행해보면서 배우는 \nGolang"

	plusString := "구름 " + "EDU\n" + "Golang"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
	fmt.Println()
	fmt.Println(plusString)
}
