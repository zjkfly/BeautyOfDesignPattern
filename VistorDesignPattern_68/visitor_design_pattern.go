package VistorDesignPattern_68

type Visitor interface {
	VisitForTriangle(*Triangle)
	VisitForSquare(square *Square)
}
type shape interface {
	Accept(Visitor)
	GetType()
}

type Triangle struct {
}
type Square struct {
}

func (s Square) Accept(visitor Visitor) {
	panic("implement me")
}

func (s Square) GetType() {
	panic("implement me")
}

func (t Triangle) Accept(visitor Visitor) {
	panic("implement me")
}

func (t Triangle) GetType() {
	panic("implement me")
}

