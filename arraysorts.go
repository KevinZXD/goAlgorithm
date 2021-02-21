package main

//假设数组 A 和 B 都是有序的，并且 A 有足够多余的空间可以合并 B。要求不使用额外的空间，把 B 合并进 A 并保证最终的数组有序

func arraySorts(A [] int, size1 int, B [] int) {
	allSize := len(A) - 1
	size2 := len(B) - 1
	size1 = size1 - 1
	for {
		if size2 < 0 || allSize <= 0 {
			break
		}
		if B[size2] >= A[size1] {
			A[allSize] = B[size2]
			size2--
		} else {
			A[allSize] = A[size1]
			size1--
		}
		allSize--
	}
	if A[0] > B[0] {
		A[0] = B[0]
	}


}

//给出两个字符串 S 和 T，要求在O(n)的时间复杂度内在 S 中找出最短的包含 T 中所有字符的子串。例如：
//S ="XDOYEZODEYXNZ"
//T ="XYZ"
//找出的最短子串为"YXNZ"
//注意：
//如果 S 中没有包含 T 中所有字符的子串，返回空字符串 “”；
//满足条件的子串可能有很多，但是题目保证满足条件的最短的子串唯一。
func findShortestSubStr(S string, T string) string {
	sL := len(S)
	indexArray := make([] int, sL)
	for j := 0; j < len(T); j++ {
		for i := 0; i < len(S); i++ {
			if T[j] == S[i] {
				indexArray[i] = j + 1
			}
		}
	}

	return ""

}
func minwindow(S string, T string) string {
	mapT := make(map[string]int)
	Ta := []byte(T)
	Sa := []byte(S)
	d := 99999
	for i := 0; i < len(T); i++ {
		mapT[string(Ta[i])] ++
	}
	begin := 0
	end := 0
	count := len(T)
	head := 0
	for end < len(S) {
		c := Sa[end]
		end++
		if mapT[string(c)] > 0 {
			count--
		}
		mapT[string(c)]++
		for count == 0 {
			if end-begin < d {
				d = end - begin
				head = begin
			}
			e := Sa[begin]
			if mapT[string(e)] == 0 {
				count++
			}
			begin++
			mapT[string(e)]++
		}
	}
	if d == 99999 {
		return ""
	}
	return ""

}
func main() {
	//A := [] int{2, 5, 8, 0, 0, 0,}
	//B := []int{3, 4, 9}
	//arraySorts(A, 3, B)
	S := "XDOYEZODEYXNZ"
	T := "XYZ"
	findShortestSubStr(S, T)

}
