package main

import "fmt"
import "strings"

func main() {

	fmt.Println(strings.Count("laoYuStudyGo", "o"))          //2
	fmt.Println(strings.Count("laoYuStudyGo", "O"))          //0
	fmt.Println(strings.Count("laoYuStudyGo", ""))           //13=12+1
	fmt.Println(strings.Count("laoYuStudyGo老虞学习Go语言", "虞"))  //1
	fmt.Println(strings.Count("laoYuStudyGo老虞学习Go语言", "Go")) //2
	fmt.Println(strings.Count("laoYuStudyGo老虞学习Go语言", "老虞")) //1
	fmt.Println(strings.Count("", ""))                       //1=0+1
	fmt.Println(strings.Count("aaaaaaaa", "aa"))             //4
	fmt.Println(strings.Count("laoYuStudyGo_n", "\n"))       //0
}
