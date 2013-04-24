package main

import "fmt"

func main() {

	s0 := []int{5, 4, 3, 2, 1} //切片
	s1 := s0[1:]
	fmt.Println("s1=", s1)

	s0[1] = 1 //切片引用的是地址
	fmt.Println("s1=", s1)

	arr := [...]int{1, 2, 3, 4, 5}
	s2 := arr[:]
	fmt.Printf("before: \n arr=%d,s2=%d \n", arr, s2)
	arr[3] = 3
	fmt.Printf("after set arr[3]=3 \n arr=%d,s2=%d \n", arr, s2)

	fmt.Println("s2=", s2, "s0=", s0, ",arr=", arr)
	s2 = append(s2, s0...)
	s2 = append(s2, arr[:]...) //追加数组元素
	fmt.Println("s2=", s2)

	s3 := append(s1, 7, 9, 11)
	fmt.Println("s3=", s3)

}
