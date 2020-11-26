package main
//import "github.com/deckarep/golang-set"


func maxCom(i int ,j int) int {
	if i>j {
		return i
	}
	return j
}
// 最长不重复字串
//func maxLengthString(str string) int {
//	strL := len(str)
//	//set := make(map[string]bool)
//	set:=make([]string,strL,strL*2)
//	//max := 0
//	for i := 0; i < strL; i = i + 1 {
//
//		set[i] = string(str[i])
//
//
//	}
//}
