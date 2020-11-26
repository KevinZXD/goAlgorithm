package main

//func findMidValue(array1 [] int,array2[] int)  {
//	length:=len(array1)+len(array2)
//	midIndex:=length/2
//
//
//}
//最长回文字串
func longestPalindrome(s string) string {
	if len(s)< 2 {
		return s
	}
	maxLen:=1
	res:=s[0:1]
	var maxLenStr string
	for i:=1;i<len(s)-1;i=i+1{
		oldStr:=centerSp(s,i,i)
		evenStr:=centerSp(s,i,i+1)
		if len(oldStr) > len(evenStr){
			maxLenStr = oldStr

		}else{
			maxLenStr = evenStr
		}
		if len(maxLenStr) > maxLen {
			maxLen=len(maxLenStr)
			res=maxLenStr
		}
	}
	return res
}
func centerSp(s string,left int,right int) string {
	length:=len(s)
	i:=left
	j:=right
	for i>=0 && j<length{
		if s[i]==s[j]{
			i++
			j--
		}else{
			break
		}
	}
	return s[i+1:j]
}