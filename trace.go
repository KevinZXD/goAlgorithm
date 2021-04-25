package main

func traceCap(tr [][] int, cap int) bool {
	var x [10] int
	row := len(tr)
	for i := 0; i < row; i++ {
		x[tr[i][1]] = tr[i][0]
		x[tr[i][2]] = -tr[i][0]
	}
	rowX := len(x)
	sum := 0
	for i := 0; i < rowX; i++ {
		sum = sum + x[i]
		if sum > cap {
			return false
		}
	}
	return true
}
func main() {
	var tr = [][]int{{2, 1, 5}, {3, 3, 7}}
    print(traceCap(tr,5))
}
