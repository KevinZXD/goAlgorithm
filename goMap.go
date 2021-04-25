package main

import "fmt"
func test11(){
	mapS := make(map[string]string)
	mapS["name"] = "zxd"
	mapS["sex"] = "female"
	val ,ok :=mapS["zxd"]
	if ok{
		println(val)
	}else{
		println("not exist")
	}

	for k, v := range mapS {
		delete(mapS,"name")
		fmt.Printf("%s:%s\n", k, v)

	}

}
func test12(){
	mapS := make(map[int]int)
	go func() {
		_=mapS[1]
	}()
	go func() {
		mapS[2]=1
	}()
	//select {
	//
	//}

}
func main() {
	test11()

}
