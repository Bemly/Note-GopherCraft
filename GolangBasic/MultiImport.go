package main

import (
	alias "fmt"
	so "os"
)

func multiImportFunc() {
	alias.Println()
	//fmt.Println()
	// 创建别名之后不能使用本名
	so.Hostname()
}
