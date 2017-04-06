package subTask

import (
	"fmt"
)

func DemostrateSliceAndArray() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	slice := array[3:5]   // fixed to this length, can not be extended anymore.
	fmt.Print(cap(slice)) // capacity count from first number of the slice to the end of the array.
}

/*
slices, on the other hand, is a dynamically-sized version of array
means that slice is another view of array.It can not exist with out a array. [a:b] will create a slice,but it will not copy the array data to a new place.
in fact, it point to the original data.
*/
func BasicSlice() {
	var slices = []int{1, 2, 3, 4}                  //notice the difference of slices and arrays,no number in the square brackets.
	fmt.Print(slices[1:4])                          // slices[a,b] like java subString, from a(included) to b(excluded).more mathmatically,[a,b)
	var dynamicSlice []int                          //this slice is nonsence
	var appendSlice = append(dynamicSlice, 1, 2, 3) //except with append function
	fmt.Println(appendSlice)
}

/*array is a fixed length data structure like other compiled language such as C,C++,Java*/
func Arrays() {
	var array = [2]string{"Hello", "World"}
	var b [2]string //no initialized,no = sign.
	array[0] = "Hello"
	fmt.Print(array[0])
	fmt.Print(b[1])
}

func InsideType() {
	type Vertex struct {
		X int
		Y int
	}
	v := &Vertex{1, 2}
	fmt.Print(v.X)
}

func Pointers() {
	//go has pointer, yes, the sophisticate C pointer.
	var p1 *int
	i := 12
	p2 := &i //& sign generate a pointer to its operand.
	//p1 and p2 both are pointers to a int value. that's the * and & means in pointer semanics.
	//look at the other mean of * sign
	fmt.Println(*p2) // this will retrive the data which stored at the loation that p2 pointed
	fmt.Println(p1)
	//fortunately go does not support pointer arithmetic as C.
}
