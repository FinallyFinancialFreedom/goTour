package subTask

import (
	"fmt"
)

func CreateMap() {
	var m = make(map[int]string)
	m[1] = "hello"
	m[2] = "world"
	fmt.Print(m[1])
}
