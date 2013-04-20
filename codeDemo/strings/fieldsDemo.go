package main

import "fmt"
import "strings"

func main() {
	fmt.Println(strings.Fields("lao Yu Study Go ")) //OutPut: [lao Yu Study Go]
	fmt.Println(strings.Fields("   Go    "))        //[Go]
	fmt.Println(strings.Fields(""))                 //[]
	fmt.Println(strings.Fields(" \n go"))           //[go]

	mySp := func(c rune) bool { return c == '#' }
	fmt.Println(strings.FieldsFunc("lao###Yu#Study####Go#G ", mySp)) //[lao Yu Study Go G<space>]

}
