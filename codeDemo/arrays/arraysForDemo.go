package main

import "fmt"

func main() {
	arr := [5]int{5, 4, 3}

	//遍历数组
	for index, value := range arr {
		fmt.Printf("arr[%d]=%d \n", index, value)
	}

	fmt.Println("------------------------")

	for index := 0; index < len(arr); index++ {
		fmt.Printf("arr[%d]=%d \n", index, arr[index])
	}
}
