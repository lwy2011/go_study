package main

import (
	"fmt"
	"time"
)

func main() {
	//初始化条件，判断条件，变化条件
	//i 块级变量
	sum := 0
	for i := 1; i < 10; i++ {
		sum = sum + i
		fmt.Println(i, sum)
	}
	str := "liuliuwan"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
	//元素位置，元素本身，两个返回值！
	for i, data := range str {
		fmt.Println(i, data)
	}
	for i, _ := range str {
		fmt.Println(i, "下标")
	}
	for _, data := range str {
		fmt.Println(data, "数据")
		fmt.Printf("%d\n", data)
		fmt.Printf("%c\n", data)
	}
	//嵌套
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
		}
	}
	//百鸡问题
	count := 0
	count1 := 0
	for man := 0; man <= 20; man++ {
		for woman := 0; woman <= 33; woman++ {
			child := 100 - man - woman
			count++
			if child%3 == 0 && (man*5+woman*3+child/3 == 100) {
				fmt.Printf("man:%d,woman:%d,child:%d\n", man, woman, child)
			}
		}
	}
	for man := 0; man <= 20; man++ {
		for woman := 0; woman <= 33; woman++ {
			for child := 0; child <= 100; child += 3 {
				count1++
				if child+woman+man == 100 && (child/3+woman*3+man*5 == 100) {
					fmt.Printf("man:%d,woman:%d,child:%d\n", man, woman, child)
				}
			}
		}
	}
	fmt.Printf("count=%d,count1=%d\n", count, count1)
	i := 0
	for { //死循环，for后面没有任何的条件就是true
		i++
		time.Sleep(time.Second) //延时
		if i == 3 {
			continue //跳过这次，只能在for循环中使用！
		}
		fmt.Println("i=", i)
		for j := 0; j <= 3; j++ {
			if i == 4 {
				fmt.Println("j is break!")
				break
			}
			time.Sleep(time.Second)
			fmt.Println("j is not break!")
		}
		if i == 5 {
			fmt.Println("break!", i)
			break //终止最近的内循环,嵌套的时候注意啦
		}
		//goto跳转！！
		fmt.Println("不跳")
		goto end
		fmt.Println("要被跳啦！")
	end:
		fmt.Println("跳过来啦，老弟！")
		goto end //再跳回去，死循环！
	}
}
