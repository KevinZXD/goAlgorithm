package main

import "fmt"

func main() {
	//mapS := map[string]string{}
	mapS := make(map[string]string)
	mapS["name"] = "zxd"
	mapS["sex"] = "female"
	val ,ok :=mapS["zxd"]
	if ok{
		println(val)
	}else{
		println("not exist")
	}
	delete(mapS,"name")
	for k, v := range mapS {
		fmt.Printf("%s:%s\n", k, v)
	}

}
