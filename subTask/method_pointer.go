package subTask

//declare a method. not class method but type method.and not in type struct
type Vertex struct {
	X, Y float64
}

// the "(v Vertex)" called receiver, which makes a func to a method.
// method can manipulate the receiver
func (v Vertex) I_am_a_method() float64 {
	return v.X + v.Y
}

// equivalent func version
func I_am_a_func(v Vertex) float64 {
	return v.X + v.Y
}

//this method can not modify the original Vertex.
func (v Vertex) I_am_a_method_modify_receiver_butfailed() float64 {
	v.X += 1
	return v.X
}

//this method can not modify the original Vertex.
func (v *Vertex) I_am_a_method_modify_receiver_success() float64 {
	v.X += 1
	return v.X
}

/*
	d := task.Vertex{3, 4}
	fmt.Print(d.I_am_a_method())
	fmt.Print(task.I_am_a_func(d))
*/
