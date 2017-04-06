package subTask

type MyFloat64 float64
type MyInt int

func (v *MyFloat64) Abs() float64 {
	if float64(*v) > 0 {
		return float64(*v)
	} else {
		return float64(*v) * (-1)
	}
}

func (v *MyFloat64) AbsInt() int {
	if float64(*v) > 0 {
		return int(*v)
	} else {
		return int(*v) * (-1)
	}
}

func (v MyInt) AbsInt() int {
	if int(v) < 0 {
		return -int(v)
	} else {
		return int(v)
	}
}

//a method declared a float64 return value can not return a int value
func (v MyInt) Abs() float64 {
	if int(v) < 0 {
		return float64(-v)
	} else {
		return float64(v)
	}
}
