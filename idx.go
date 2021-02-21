package main

import "fmt"

type student struct {
	Name string
	Age  int
}

//for变量是生成一个副本，把slice、map中的值赋值给副本
func test1() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	stus[2].Age += 222
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, *v)
	}
}

// 会编译不成功，map的value不能寻址
func test2() {
	var m = map[string]*student{}
	m["zhou"] = &student{
		Name: "zhou",
		Age:  24,
	}
	m["zhou"].Age += 1
	x := student{Name: "sdf", Age: 55}
	x.Age = 8
	fmt.Println(m["zhou"].Age)
	fmt.Println(x.Age)
}

type People1 struct{}

func (p *People1) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People1) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People1
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

//golang中没有继承的概念。 调用t.ShowA() === t.People.ShowA()
func test3() {
	t := Teacher{}
	t.ShowA()
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//defer 后面函数的值是即时传递
func test4() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	panic(1)
}
func test12() {
	defer println(1)
	defer println(2)
	panic(4)
}

//break只能退出一层循环。channel关闭后可以读，不可以写
func test5() {
	ch := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	for {
		select {
		case item, ok := <-ch:
			if !ok {
				fmt.Println("close")
				break
			}
			fmt.Println(item)
		}
	}

}
func doubleSlice(in []int) {
	for k, v := range in {
		in[k] = v * 2
	}
}
func doubleArray(in [10]int) {
	for k, v := range in {
		in[k] = v * 2
	}
}

//1、golang函数调用都是值传递。2、slice数据结构
func test6() {
	var slice = []int{1, 2, 3, 4, 5}
	doubleSlice(slice)
	fmt.Println(slice)

	var array = [10]int{1, 2, 3, 4, 5}
	doubleArray(array)
	fmt.Println(array)
}

type People interface {
	Speak(string) string
}
type Student struct{}

func (stu Student) Speak(think string) (talk string) {
	if think == "learning" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

//接口实现 Student类型没有实现Speak方法。*Student类型才实现
func test7() {
	var peo People = &Student{}
	think := "learning"
	fmt.Println(peo.Speak(think))
}

func test8() {

	type Books struct {
		title   string
		author  string
		subject string
		book_id int
	}

	book := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}

	// 创建一个新的结构体
	fmt.Println(book.title)
	fmt.Println(&book.title)

}
func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	test12()
}
