package subTask

//interface,contains a set of METHOD signature, not FUNC signature.
//although it conatains methods, variable of it's type only hold value that implements those methods.
//notice that value assign to this interface type must implements all the methods it contains. Like java implements a interface should implements all the methods define in interface
type Abser interface {
	Abs() float64
	AbsInt() int
}
