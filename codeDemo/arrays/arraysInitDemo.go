package main

import "fmt"
import "reflect"

func main() {
	//声明数组
	arr1 := [2]int{1, 2}               //两个长度
	arr2 := [5]int{1, 2}               //5个长度
	arr3 := [...]int{1, 2, 3, 4, 5, 6} //根据元素个数动态确定的数组长度
	arr4 := []int{1, 2, 3}             //注意：这里定义的不是数组而是切片slice(类似数组) 
	fmt.Println(reflect.TypeOf(arr1), arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println("type is =", reflect.TypeOf(arr4), arr4) //slice

}
