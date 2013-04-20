package main

import "fmt"
import "strings"

func main() {
	str := "laoYuStudyGotrue是否包含某字符串"
	fmt.Println(strings.Contains(str, "go"))    //false
	fmt.Println(strings.Contains(str, "Go"))    //true
	fmt.Println(strings.Contains(str, "laoyu")) //false
	fmt.Println(strings.Contains(str, "是"))     //true

	fmt.Println(Contains(str, "go", true))  //true
	fmt.Println(Contains(str, "go", false)) //false

	fmt.Println(strings.Contains(str, "")) //true
}

//在字符串s中是否包含字符串substr,ignoreCase表示是否忽略大小写
func Contains(s string, substr string, ignoreCase bool) bool {
	return Index(s, substr, ignoreCase) != -1

}

//字符串subst在字符串s中的索引位置,ignoreCase表示是否忽略大小写
func Index(s string, sep string, ignoreCase bool) int {
	n := len(sep)
	if n == 0 {
		return 0
	}

	//to Lower
	if ignoreCase == true {
		s = strings.ToLower(s)
		sep = strings.ToLower(sep)
	}

	c := sep[0]
	if n == 1 {
		// special case worth making fast
		for i := 0; i < len(s); i++ {
			if s[i] == c {
				return i
			}
		}
		return -1
	}
	// n > 1
	for i := 0; i+n <= len(s); i++ {
		if s[i] == c && s[i:i+n] == sep {
			return i
		}
	}
	return -1

}
