package main

import "fmt"

// edit time 2024/01/15
// 本代码是 GoWeb编程实战派 的笔记

// 全局变量
var golbal string

func main() {
	fmt.Println("Hello World~")
	fmt.Println()

	multiImportFunc()

	//var a, b *int
	//	.\helloworld.go:11:6: a declared and not used
	//	.\helloworld.go:11:9: b declared and not used
	// 这玩意必须要使用才不报错
	fmt.Println("Hello World~")

	var name string
	fmt.Print(name)

	var (
		age         int
		notStructOh struct {
			name string "1"
			// 这玩意结构不会出现变量未使用的报错oh
		}
		sex bool
	)
	fmt.Print(age)
	fmt.Print(sex)
	fmt.Print(notStructOh.name)

	var language string = "GoGoGo"
	fmt.Print(language + "ZX")

	// 声明的简单形式：
	// 只能在函数内部，不能声明全局变量，不能提供数据类型
	a := 1
	fmt.Println(a)

	// 隐藏类型
	var (
		b = 2
		c = 3
		d = "自动识别"
	)
	fmt.Print(b + c)
	fmt.Print(d)

	// 聚合简短形式
	e, f, g := 1, true, "nb"
	fmt.Print(e)
	fmt.Print(f)
	fmt.Println(g)

	// 一行交换基本变量 bce before:231 after:312
	// 不需要中间变量
	b, c, e = e, b, c
	fmt.Printf("b:%d\nc:%d\ne:%d\n", b, c, e)

	// 全局变量
	golbal = "hao"
	fmt.Println(golbal)

	//e, f = f, e
	//.\helloworld.go:67:9: cannot use f (variable of type bool) as int value in assignment
	//.\helloworld.go:67:12: cannot use e (variable of type int) as bool value in assignment
	// 只能同类

	// 常量constant
	const pi = 3.14
	// 常量的值必须是编译时可以确定的 比如5/2 如果是函数什么的运行时才能确定的是无法赋值的

	const ipv4Lenth = 4
	var p [ipv4Lenth]byte
	fmt.Println(p)

	// iota 常量生成器 => 枚举类型
	type pos int
	const (
		x1 pos = iota
		y1
		z1
	)
	const (
		x2 pos = 0
		y2
		z2
	)

	// go没有while until(do while)
	res := 1
	for i := 1; i < 20; i++ {
		res *= i
		//fmt.Print(i++)
		//.\helloworld.go:97:14: syntax error: unexpected ++ in argument list; possibly missing comma or )
		// 不能在参数里面自增自减
		fmt.Print(i)
	}
	fmt.Printf("\n%d\n", res)

	//while:
	i := 0
	for {
		i++
		if i > 10 {
			fmt.Println(i)
			break
		}
	}
	// 简化
	for i < 100 {
		i++
		fmt.Print(".")
	}
	fmt.Println()

	// 选择层级break 支持for所有
BreakFor:
	for {
		for {
			for {
				for {
					i++
					break BreakFor
					fmt.Print("我不打扰了")
				}
			}
		}
	}

	// for-range 迭代结构 遍历数组、切片、字符串、Map、channel 类似PHP foreach
	for num, char := range golbal {
		fmt.Printf("[%d]=%c\\0x%x\n", num, char, char)
	}

	println(6 / 5)
	//println(6 // 5) 无
	println(6.0 / 5)

	// 切片 数组 列表 list 有序！！不是集合Set
	// 这里匿名声明
	for iota, value := range []int{99, 12, 243, 3} {
		fmt.Printf("iota:%d value:%d\n", iota, value)
	}
	// 对象 map 字典 json 此操作遍历是无序的！！
	for key, value := range map[string]int{
		"go":  20,
		"web": 90,
	} {
		fmt.Printf("key:%d value:%d\n", key, value)
	}
	// 这个格式和json末尾多了个逗号！！！！
	// [void]类型 => 切片
	// [key类型]value类型 => Map

	// channel
	ch := make(chan int) // 创建整型通道
	go func() {
		ch <- 7
		ch <- 8
		ch <- 9
		close(ch)
	}() //匿名函数传给goroutine启动，把789推进通道
	for chvar := range ch {
		fmt.Print(chvar) // 打印存放进通道的数据流
	}

	// 变量_是占位符

	// match switch
	golbal = "aa"
	switch golbal {
	case "aa":
		fmt.Println(golbal)
	case "bb": //"自带break，不会执行到下一行，但是失去很多and操作"
	default:
		fmt.Println("最后的，一样用法")
	}
	// 但是失去很多and操作 =》 是不可能的，但是case的类型和java一致，强类型
	switch golbal {
	case "bb", "aa":
		fmt.Println(golbal)
	}
	// 和python 一样的操作 可以塞表达式在case 前提是switch
	e = 1
	switch {
	case i > 20 && i < 30:
		fmt.Println(golbal)
	}

	// Go Data Type
	// basic Type:
	// bool(true, false)
	// int/uint(uint8/byte,uint16,uint32,uint64,
	//		int8,int16,int32/rune,int64,float32,float64取代double
	//		complex64,complex128 实数虚数, uintptr 指针！！！)
	// string/char(UTF-8存储)
	// complex Type:
	// List/数组/切片/Map/结构体

	// &&比||优先级高！！！！！！！！！！不需要再像其他同优先级语言一样加括号了

	// bool有点特殊，不能和数字其他转换
	fmt.Println(string(100) + "1")
	//fmt.Println(100 + "1") 错 没拼接跨类型
	fmt.Println(`

	不支持转义符哦 类似py''' '''\n

	`)

	// 字符串不可变 但是字符串支持级联+ 追加+=
	str := string(97) + string(98)
	str += string(99) + string(100)
	fmt.Println(str)
	// 字符串分割 [:::]比py少一个步长
	fmt.Println(str[1:2])    //索引:长度 X 0索引:1索引 V
	fmt.Println([]rune(str)) // string(rune) 码点值转char char转码点值数组
	fmt.Println([]byte(str)) // rune是聚合字符 byte聚合字节
	fmt.Println([]int32(str))
	fmt.Println([]uint8(str))

	// 字符串遍历除了range 还可以使用[]rune(str)
	// 因为狗语言是使用UTF-8储存压缩的，一些字符是单字节一些是多字节
	// 所以索引法不一定可靠，我们需要用指针步长或者是转换为单字符！（没说单字节）索引
	str = "那n没m事s了l"
	fmt.Println(string(str[0])) // 这里只提取到了单字节,中文占三字节
	goto 那                      //go语言就是要用goto！！！！！！！
那:
	fmt.Println(str[0])
	fmt.Println(str[1])
	fmt.Println(str[2])
	fmt.Println([]byte{str[0], str[1], str[2]})
	fmt.Println(string([]byte{str[0], str[1], str[2]}))
	fmt.Println([]rune("那"))
	goto n

n:
	fmt.Println(string(str[3])) // 根据UTF-8 字母属于单字节ASCII系

	// 只用单次追加可以+= 多次用类似java的stringbuilder TODO:多次用类似java的stringbuilder

	// 修改和上面提取一致 自己考虑rune还是byte

	j := &e
	println(*j == e)
	fmt.Printf("%p %p %p\n", j, *j, &j)

	*j, e = e, *j
	fmt.Printf("%p %p %p\n", j, *j, &j)

	var arr [10]int
	fmt.Printf("arr是数组=>%d", arr)
	var arr2 [10]int = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var arr3 []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var arr4 = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var arr5 = [...]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var arr6 = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//var arr7 [...]int
	//.\helloworld.go:264:11: invalid use of [...] array (outside a composite literal)
	fmt.Printf("\n%d\n%d\n%d\n%d\n%d\n", arr2, arr3, arr4, arr5, arr6)

	var parr1 = [][10]int{
		arr2, arr4, arr5, arr,
	} //这个最后逗号换行就要 不换就不要 [10]和不写[]都行 但是不属于同类
	var parr2 = [][]int{arr3, arr6}
	fmt.Printf("\n%d\n%d\n", parr1, parr2)

	// struct 可以套娃 甚至父级可以套父级

	// slice 切片 = C:数组 Py:List
	// 切片比指针更安全更强大 原话 不是我说的哈

	// 函数里面不能再套函数了
	println(变参(1, 2, 3, 4, 5, 6))

	// goto OOP
	OOP("面向对象章节")
	RunTime("反射 goroutine go编译 单元测试章节")
}
func 变参(arg ...int) int {
	count := 0
	for range arg {
		count++
	}
	return count
}
