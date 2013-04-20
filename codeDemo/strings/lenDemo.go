package main

import "fmt"

type UserInfo struct {
	id   int
	name string
}

type Interger int

func main() {
	str := "str"
	arr := [5]int{1, 2, 3}
	slice := make([]int, 5)

	m := make(map[int]string)
	m[2] = "len"

	ch := make(chan int)

	//user :=&UserInfo{id:1,name:"laoYu"}
	//interger :=2

	fmt.Println("len(string)=", len(str)) //3
	fmt.Println("len(array)=", len(arr))  //5invalid argument user (type *UserInfo) for len

	fmt.Println("len(slice)=", len(slice)) //5
	fmt.Println("len(map)=", len(m))       //1
	fmt.Println("len(chat)=", len(ch))     //0
	//fmt.Println("len(my struct)=",len(user))//invalid argument user (type *UserInfo) for len
	//fmt.Println("len(interger)=",len(interger))

	var str2 string
	var arr2 [5]int
	var slice2 []int
	var m2 map[int]string
	var ch2 chan int

	fmt.Println("len(string)=", len(str2)) //3
	fmt.Println("len(array)=", len(arr2))  //5invalid argument user (type *UserInfo) for len

	fmt.Println("len(slice)=", len(slice2)) //5
	fmt.Println("len(map)=", len(m2))       //1
	fmt.Println("len(chat)=", len(ch2))     //0
	//fmt.Println("len(my struct)=",len(user))//invalid argument user (type *UserInfo) for len
	//fmt.Println("len(interger)=",len(interger))


}
