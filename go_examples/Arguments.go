package main
import (
	"fmt"
	"flag"
)


// 在Go语言的flag包中，命令行参数以指针返回

func main(){
	// 定义一个类型为string，名称为surname的命令行参数
	// 命令行参数名称、默认值、提示
	surname := flag.String("surname", "王", "您的姓")

	// 定义一个类型为String，名称为personalName的命令行参数
	// 除了返回指针类型结果，还可以直接传入变量地址获取参数值
	var personalName string
	flag.StringVar(&personalName, "personalName", "小儿", "您的名")

	// 
	id := flag.Int("id", 0, "your id")

	flag.Parse()
	fmt.Printf("I am %v %v, and my id is %v\n", *surname, personalName, *id)
}