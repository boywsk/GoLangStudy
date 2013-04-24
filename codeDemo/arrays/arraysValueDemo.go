package main

import "fmt"

func main() {
	arr2 := [5]int{1, 2} //5个长度
	arr5 := arr2         //数组赋值，是复制数组，而不是引用指针
	arr5[0] = 5
	arr2[4] = 2
	fmt.Printf(" arr5= %d \n arr2=%d \n arr5[0]==arr2[0]= %v \n", arr5, arr2, arr5[0] == arr2[0])

}
