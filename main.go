package main

import "fmt"

func main() {
	age, name := 24, "길동"

	fmt.Printf("안녕 난 %s이야\n", name)
	fmt.Printf("나이는 %d살이야\n", age)
	fmt.Printf("반가워")

	n := 14
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("내 동생은 %d살이야\n", n)
	fmt.Printf("%d \n", arr)
	fmt.Printf("%d %d", n, arr)
}
