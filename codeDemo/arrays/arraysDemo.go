package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2}
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr4 := [5]int{2: 1, 3: 2, 4: 3}
	arr5 := [...]int{2: 1, 4: 3}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)

	arr := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d]=%d \n", i, arr[i])
	}
	for i, value := range arr {
		fmt.Printf("arr[%d]=%d \n", i, value)
	}

	arr6 := arr
	fmt.Printf("arr=%d,arr6=%d \n", arr, arr6)
	arr[0] = 8
	fmt.Printf("arr=%d,arr6=%d \n", arr, arr6)

	arr7 := arr
	arr = arr2
	fmt.Printf("arr7=%d \n", arr7)
}
