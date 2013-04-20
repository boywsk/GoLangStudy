package main

import "fmt"

func main() {
	str := "laoYu老虞"
	str2 := "laoYu"
	fmt.Println("len(", str, ")=", len(str))   //len=11=5+6,一个汉字在UTF-8中占3个字节
	fmt.Println("len(", str2, ")=", len(str2)) //len=5

	fmt.Println("str[0]=", str[0]) //l

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
	for i, s := range str {
		fmt.Println(i, "Unicode(", s, ") string=", string(s))
	}

	r := []rune(str)
	fmt.Println("rune=", r)
	for i := 0; i < len(r); i++ {
		fmt.Println("r[", i, "]=", r[i], "string=", string(r[i]))
	}
}
