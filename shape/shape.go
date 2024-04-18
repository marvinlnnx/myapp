package shape

import "math"

type IShape interface {
	Calculate() float64
	CalculateCircleArea() float64
}

type Shape struct {
	Height    float64
	Weight    float64
	ShapeType string
}

type Area struct {
	Radius float64
}

func NewShape(h, w float64, shapeType string) Shape {
	return Shape{
		Height:    h,
		Weight:    w,
		ShapeType: shapeType,
	}
}

func NewArea(r float64) Area {
	return Area{
		Radius: r,
	}
}

func (s Shape) Calculate() float64 {
	switch s.ShapeType {
	case "squre":
		return math.Pow(s.Height, 2.0)
	case "rectangle":
		return s.Height * s.Weight
	default:
		return -1
	}
}

func (s Shape) CalculateCircleArea() float64 {
	return 0.0
}

func (s Area) Calculate() float64 {
	return 0.0
}

func (s Area) CalculateCircleArea() float64 {
	return math.Pow(s.Radius, 2) * 3.12
}
