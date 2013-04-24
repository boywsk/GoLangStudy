package main

import "fmt"

func main() {
	//切片声明
	s0 := make([]int, 5, 10)
	fmt.Printf("len(s0)=%d,cap(s0)=%d,yields=%d \n", len(s0), cap(s0), s0)

	arr := [5]int{5, 4, 3, 2, 1}

	//切片可以从数组中复制， 
	//array[startIndex:endIndex]表示：获取数组中第startIndex位置到endIndex-1区间的元素. 其切片长度为:endIndex-startIndex
	// startIndex和endIndex为可选项
	//startIndex 不设置时，默认从头开始，endIndex 不设置时，默认截取到末尾

	s1 := arr[:]     //全部，从头到末尾
	s2 := arr[3:]    //从第3+1个到最后
	s3 := arr[:3]    // 从头到 第3+1个(不包括第3+1个)
	s4 := arr[1:3]   //从第1+1个到 第3+1个
	s5 := s1[1:3]    //来自切片s1
	s6 := new([]int) //指针
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("s3=", s3)
	fmt.Println("s4=", s4)

	fmt.Println("s5=", s5)
	fmt.Println("s6=", s6, ",cap=", cap(*s6), ",len=", len(*s6))

}
