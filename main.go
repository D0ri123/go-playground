package main

import "fmt"

func main() {
	var name string
	var gen string
	var school string

	fmt.Print("이름, 성별, 학교를 입력하세요.")

	/**
	공백으로 구분하여 입력한다.
	*/
	//fmt.Scanln(&name, &gen, &school)

	/**
	공백과 개행으로 구분하여 입력한다.
	*/
	fmt.Scan(&name, &gen, &school)
	fmt.Println("이름은 ", name, "성별은 ", gen, "학교는", school)

	/**
	포맷 지정자를 이용하여 개발자가 원하는 형태로 입력한다.
	*/
	var i, j int
	fmt.Print("주민등록번호를 -를 이용해 입력하세요 :")
	fmt.Scanf("%d-%d", &i, &j)

	fmt.Printf("주민등록번호는 %d-%d\n", i, j)
	fmt.Println("앞자리는", i)
	fmt.Println("뒷자리는", j)
}
