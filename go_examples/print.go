package main

import "fmt"
import "bytes"

func main()  {
	fmt.Print("Hello", " ", "World")
	// 格式化输出
	fmt.Printf("Hello, %s!\n", "World")
	// 末尾自动添加换行符
	fmt.Println("hello, world!")


	// fmt.Sprintf系列
	// 这些函数不会直接打印，而是将格式化后的结果作为字符串返回。
	result := fmt.Sprint("Hello"," ","world")
	fmt.Println(result)
	result = fmt.Sprintf("Hello, %s!", "World")
	fmt.Println(result)
	result = fmt.Sprintln("Hello", "World")
	fmt.Print(result)

	// fmt.Fprint系列
	// fmt.Fprint：将格式化后的内容写入到指定的 io.Writer
	var buf bytes.Buffer
	fmt.Fprint(&buf, "Hello", " ", "World")
	fmt.Println(buf.String())

	// fmt.Fprintf：支持格式化，并将结果写入指定的 io.Writer
	// var buf bytes.Buffer
	fmt.Fprintf(&buf, "Hello, %s!", "World")
	fmt.Println(buf.String())

	// fmt.Fprintln：在 fmt.Fprint 基础上追加换行符
	// var buf bytes.Buffer
	fmt.Fprintln(&buf, "Hello", "World")
	fmt.Print(buf.String())

	// fmt.Error 系列
	// fmt.Errorf：创建一个带格式化消息的 error 对象
	err := fmt.Errorf("an error occurred: %s", "something went wrong")
	fmt.Println(err)
	// 输出: an error occurred: something went wrong




}