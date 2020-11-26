package main

func findArray(array [][] int, target int) bool {
	rows := len(array)
	columns := len(array[0])
	i := rows - 1
	j := 0
	for {

		if array[i][j] == target {
			return true
		}
		if array[i][j] < target {
			j++
		} else {
			i--
		}
		if i < 0 || j >= columns {
			break
		}
	}
	return false

}
func main() {
	target := 7
	array := [][]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}}
	println(findArray(array,target))
	println(findArray(array,1))

	println(findArray(array,0))
}
