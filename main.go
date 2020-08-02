package main //声明main包，表明当前是一个可执行程序
import "fmt"
//函数外只能防止标识符（变量\常量\函数\类型）的声明
//fmt.Println("Hello")//非法

func main() { //main函数，是程序执行的入口
	fmt.Println("Hello World!") //在终端打印Hello World!
}
