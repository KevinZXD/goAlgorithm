package main

import "fmt"

/*
f（1)=f(2) =1

f(n) = f(n-1)+f(n-2)

log(n) 的复杂度

分析

[f(n),f(n-1)]*[1,1
               1,0]   =[f(n)+f(n-1),f(n)]
*/
func vector(arr1 [][] int, arr2 [][] int) [][] int {
	var result [][] int
	arr1Row := len(arr1)
	arr1Col := len(arr1[0])
	arr2Row := len(arr2)
	arr2Col := len(arr2[0])
	if arr1Row != arr2Col || arr1Col != arr2Row {
		fmt.Printf("invalid parameter")
		return result
	}

		for j := 0; j < arr1Col; j++ {
			for k := 0; k < arr2Row; k++ {
				for i := 0; i < arr1Row; i++ {
					result[j][k]+=arr1[j][i]*arr2[j][k]
				}

			}
		}

}
