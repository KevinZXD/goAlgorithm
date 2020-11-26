package main

func replaceNullChar(str string) string {
	charLen := len(str)
	replaceChar := ""
	for i := 0; i < charLen; i++ {
		if str[i] == ' ' {
			replaceChar = replaceChar + "%20"
		} else {
			replaceChar = replaceChar + string(str[i])
		}
	}
	return replaceChar

}
func main()  {
	testString := " ss dd "
	println(replaceNullChar(testString))

}
