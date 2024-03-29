package main

import "fmt"

var b = 8 //全局变量   ，小心变量名冲突,其他文件中的全局变量名
func t() {
	a := 5 //局部变量，函数内部的变量优先，也是分层级的，区块内部是唯一的,跟全局变量名不怕冲突，因为冲突了，也是先用局部的变量。
	a++
}

var a = 10 //全局变量
func t6() {
	b = 10
	b++
}
func t7() {
	fmt.Println(a) //这里不遵循js里的作用域链的知识，它找的是声明时的区块那里，而不是，执行时的区块那里
}
func main() {
	//fmt.Println(a)访问不到a，声明不会提升的，不同于js。
	fmt.Printf("b=%d\n", b) //全局变量可以随意调用
	a := 9                  //局部变量
	t()
	t6()
	fmt.Println(a, b)
	for i := 0; i < 10; i++ {
	}
	//fmt.Println(i) 访问不到i，所以for循环里面的i是块级变量
	t7()            //期望中是希望打印执行域中的局部变量，实际上是函数声明那里的区块里的
	fmt.Println(t7) //函数的地址,代码区。 函数的内容信息放在栈取
	fmt.Println(&b) //全局变量的地址，数据区的地址。
	fmt.Println(&a) //局部变量的地址，栈区，因为在函数内部跟函数信息的地址一样的
}
