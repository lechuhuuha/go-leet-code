package main

import "fmt"

type IDrawShape interface {
	drawShape(x, y [5]float32)
}

type DrawShape struct {
}

func (d DrawShape) drawShape(x, y [5]float32) {
	fmt.Println("Drawing Shape")
}

type IContour interface {
	drawContour(x, y [5]float32)
	resizeByFactor(f int)
}

type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

func (d DrawContour) drawContour(x, y [5]float32) {
	fmt.Println("Drawing Contour")
	d.shape.drawShape(x, y)
}

func (d DrawContour) resizeByFactor(f int) {
	d.factor = f
}
