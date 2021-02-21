package main

import "strconv"

func quickSort(array [] int, start int, end int) {
	if start > end {
		return
	}
	index := partition(array, start, end)
	quickSort(array, index+1, end)
	quickSort(array, start, index-1)
}
func partition(array [] int, start int, end int) int {
	pvalue := array[start]
	for start < end {
		for start < end && pvalue < array[end] {
			end--
		}
		array[start] = array[end]
		for start < end && pvalue > array[start] {
			start++
		}
		array[end] = array[start]
	}
	array[start] = pvalue
	return start
}
func ishuiwen(num int) bool {
	numstr := strconv.Itoa(num)
	numstrLen := len(numstr)
	start := 0
	end := numstrLen - 1
	if numstrLen%2 != 0 {
		for {
			if start == end {
				break
			}
			if numstr[start] != numstr[end] {
				return false
			}
			start++
			end--
		}
	} else {
		for {

			if numstr[start] != numstr[end] {
				return false
			}
			if start == end-1 {
				break
			}
			start++
			end--
		}
	}

	return true
}
func main() {

	println(ishuiwen(121))
	println(ishuiwen(123))
	println(ishuiwen(1211))
	println(ishuiwen(1221))

}
