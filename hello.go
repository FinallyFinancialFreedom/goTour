package main

import (
	"fmt"
	task "tourBase/subTask"
)

func main() {
	//e := task.MyFloat64(-2.718)
	caoliu := task.MyInt(1024)
	//fmt.Println(e.Abs())
	fmt.Println(caoliu.AbsInt())

	var a task.Abser
	f := task.MyFloat64(2.718)
	a = &f
	fmt.Printf("%T", a)
	fmt.Println()
	fmt.Println(a.AbsInt())
}
