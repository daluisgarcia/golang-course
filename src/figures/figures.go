package figures

type HasArea interface {
	Area() float64
}

type Square struct {
	Side float64
}

type Rectangle struct {
	Base   float64
	Height float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (r Rectangle) Area() float64 {
	return r.Base * r.Height
}
