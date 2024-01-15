package main

import (
	"fmt"
	"reflect"
)

// edit time 2024/01/15

func RunTime(str string) {
	println(str)
	println(reflect.TypeOf(new(Pos)))
	fmt.Println("Type:", reflect.TypeOf(new(Pos)))
	println(reflect.TypeOf(RunTime))
	fmt.Println("Type:", reflect.TypeOf(RunTime))
	println(reflect.TypeOf(1))
	fmt.Println("Type:", reflect.TypeOf(1))
	var x float64 = 3.14
	fmt.Println("Type:", reflect.TypeOf(x))

	// 不接收变量指针创建的反射对象，是不具备可写性的 用CanSet()来得知是否可写，Elem()返回指针指向的数据
	name := "Hao"
	println("name>", name)

	v1 := reflect.ValueOf(&name)
	fmt.Println("v1>", v1)
	println(v1.CanSet())

	v2 := reflect.ValueOf(name)
	fmt.Println("v2>", v2)
	println(v2.CanSet())

	v3 := v1.Elem()
	fmt.Println("v3>", v3)
	println(v3.CanSet())

	// 通过反射在运行时动态修改值
	v3.SetString("BuHao")
	println("name>", name)
	fmt.Println("v1>", v1)
	fmt.Println("v2>", v2)
	fmt.Println("v3>", v3)
	// v2没有赋地址，所以新开辟的内存空间指向，把name纯拷贝过来了，v1是直接copy的uintptr

	// 每一个并发活动成为goroutine 给函数就会执行一个独立的并发线程
	// 详细可以看es6(js)委托promise csharp事务 异步

	// go虽然不像py有PASS关键字 但是go有单元测试
	// 单元测试会输出PASS 赢
	// 文件名必须是<name>_test.go格式，函数必须以Test开头，传入参数类型必须是*testing.T

	// 交叉编译：在一个平台编译其他平台的可执行文件
	// GOOS 目标go操作系统operate system GOARCH 目标go架构arch

	// Go语言编译器有两种 一种CGO C语言写的编译器 一种用Go写的Go编译器
	// 可以在Go里面引用C语言的库
	// 但是捏 在1.5版本之后Go逐步脱离CGO 所以如果有C库引用需要CGO_ENABLED=1来启用

	// go install 会把编译好滴放在go的统一环境变量处
	// $GOPATH/pkg $GOPATH/bin

	// go get 可以借助git、svn、hg这些直接拉取代码库，类似maven
	// -u 自动网络寻找依赖包更新进行爆破 笑
}
