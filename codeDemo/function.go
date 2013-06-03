package main

import (
	"fmt"
)

//定义新的类型double，主要目的是给float64类型扩充方法
type double float64

//判断a是否等于b
func (a double) IsEqual(b double) bool {
	var r = a - b
	if r == 0.0 {
		return true
	} else if r < 0.0 {
		return r > -0.0001
	}
	return r < 0.0001
}

//判断a是否等于b
func IsEqual(a, b float64) bool {
	var r = a - b
	if r == 0.0 {
		return true
	} else if r < 0.0 {
		return r > -0.0001
	} else {
		return r < 0.0001
	}
}

func Log(title string, GetMsg func() string) {

	//如果开启日志记录,则记录日志
	if true {
		fmt.Println(title, ":", GetMsg())
	}
}

func main() {
	var a double = 1.999999
	var b double = 1.9999998
	fmt.Println(a.IsEqual(b))
	fmt.Println(a.IsEqual(3))
	fmt.Println(IsEqual((float64)(a), (float64)(b)))

	fc := func(msg string) {
		fmt.Println("you say :", msg)
	}

	fmt.Printf("%T \n", fc)
	fc("hello,my love")
	func(msg string) {
		fmt.Println("say :", msg)
	}("I love to code")

	count := 0
	msg := func() string {
		count++
		return "您没有即使提醒我,已触犯法律"
	}
	Log("error", msg)
	Log("warring", msg)
	Log("info", msg)
	fmt.Println(count)
}
