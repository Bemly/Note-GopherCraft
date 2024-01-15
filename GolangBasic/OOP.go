package main

// edit time 2024/01/15

func 引用传递传的是地址故修改参数就会改变原本变量值(x *int, y *int) {
	println(x, y)
	// x->a->1 y->b->2
	x, y = y, x
	println(x, y)
	// 前面交换的是指针，是参数的x->b->2 y->a->1
	*x, *y = *y, *x
	// x->b->1 y->a->2
}

func defer1() {
	println(1)
}

func defer2() {
	println(2)
}

func defer3() {
	println(3)
}

func OOP(str string) {
	a, b := 1, 2
	引用传递传的是地址故修改参数就会改变原本变量值(&a, &b)
	println(a, b)

	// 只能先声明后使用
	匿名函数有名了 := func(b bool) string { return "9" }

	// 函数可以作为参数传递
	func(我改名了 func(bool) string) {
		// go还是要区分字符串""和字符''的区别
		println("匿名函数直接使用案例，箭头后面" + 我改名了(true))
	}(匿名函数有名了)

	// defer 函数延迟语句 处理函数结束相关进程
	aver := func(i int) int {
		println("hello")
		// 按照栈stack FILO原则 3->2->1的模式取出执行
		defer defer1()
		defer defer2()
		defer defer3()
		// defer在函数结束之后再运行 意思是比return还靠后哦！！！
		println("defer你先别急")
		defer func() { i++ }()
		return i
	}(10)
	println(aver)

	// go有对象 但是没有class类的概念 我的天
	// 这时候就要用 struct 来代替这些封装 多态 继承的操作了
	// struct的成员方法和struct定义的时候需要分开写
	point := Pos{1, 2, 0, 0}
	println(point.SwitchPos().x)
	println(point.SwitchPos().y)
	// 大写public 小写private orz
	println(point.zone)
	spoint := new(Pos)
	spoint.x = 2
	spoint.y = 3
	spoint.zone = 5
	println(spoint.zone)
	// tmd为啥我全是public 啊啊啊啊 其实这个private不是在class内，即struct内 而是package内啊!!!!!!
	// 所以他也不需要extern这些美好的修饰符！！！！

	// 继承类似Csharp的成员内部类（划）
	// 上面理解错误，应该是 写在构造函数那里的接口 这个继承可以多重继承这些父类和接口
	// 同个参数叠加指明上级就行了
	mpoint := new(MaterialPos)
	println(mpoint.PosTool.y == mpoint.Pos.y)
	// 接口继承下来可以不用强制实现这个函数，但是如果一旦调用就会出现内存地址没有分配的错误情况
	// 抛出signal中断字
	// panic: runtime error: invalid memory address or nil pointer dereference
	//[signal 0xc0000005 code=0x0 addr=0x0 pc=0x6266e1]
	//println(mpoint.GetY)
	println(mpoint.GetX)
	println(mpoint.GetX())

	// interface接口不支持成员方法的实例化(即静态化) 但是支持赋值操作实现和其他类的映射关系
	// TO\DO 已解决 interface query 太难了，前面的区域以后学流处理的时候再来探索吧

	// 接口类型.(Type)

	// 多态重载与java一致

	// 反射：运行时来访问检测修改本身状态的一种能力

}

type Pos struct {
	x, y   int
	zone   int
	Length int
}

// go的成员方法不同于函数，命名是首字母也大写的驼峰 但是形参是这个结构体对象的地址
func (oldPos *Pos) SwitchPos() Pos {
	// 第一个括号负责把类传进来 通过地址绑定的方法来指明这是哪个类的成员方法
	// 第二个括号才是真正负责这个成员方法的名字和参数
	oldPos.x, oldPos.y = oldPos.y, oldPos.x
	return *oldPos
}

type PosTool struct {
	ppp, y int
}

type PosFunc interface {
	GetX()
	GetY()
	PosFunc1
	PosFunc2
}

// 接口组合
type PosFunc1 interface{}
type PosFunc2 interface{}

type MaterialPos struct {
	Pos
	PosTool
	PosFunc
	mass int
}

func (class *MaterialPos) GetX() int {
	return class.x
}
