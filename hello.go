package main

import (
	"fmt"
	"strconv"
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

	//notice what happen if there is no * sign below.
	underConcreteValue := a.(*task.MyFloat64)
	fmt.Printf("%T", underConcreteValue)

	fmt.Println()
	str1 := 127
	str2 := 0
	fmt.Print(strconv.Itoa(str1) + "." + strconv.Itoa(str2))

	i, err := parseInt("x")
	fmt.Print(err)
	fmt.Print(i)
}

func parseInt(v string) (int, error) {
	return strconv.Atoi(v)
}
