package fundamentalstest

import (
	"math"
	"testing"
)

type Shape interface {
	Area() float32
}

type Rectangle struct {
	Width  float32
	Height float32
}

func (r *Rectangle) Area() float32 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float32
}

func (c *Circle) Area() float32 {
	return (c.Radius * c.Radius) * math.Pi
}

type Triangle struct {
	Base   float32
	Height float32
}

func (t *Triangle) Area() float32 {
	return (t.Base * t.Height) * 0.5
}

func TestPerimeter(t *testing.T) {
	testCases := []struct {
		desc         string
		shape        Shape
		expectedArea float32
	}{
		{
			desc:         "test rectangle area 1",
			shape:        &Rectangle{10, 10},
			expectedArea: 100.0,
		},
		{
			desc:         "test rectangle area 2",
			shape:        &Rectangle{20, 10},
			expectedArea: 200.0,
		},
		{
			desc:         "test circle area 1",
			shape:        &Circle{10},
			expectedArea: 314.15927,
		},
		{
			desc:         "test circle area 2",
			shape:        &Circle{20},
			expectedArea: 1256.6371,
		},
		{
			desc:         "test triangle area 2",
			shape:        &Triangle{20, 10},
			expectedArea: 100,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			area := tC.shape.Area()
			if area != tC.expectedArea {
				t.Error("expected:", tC.expectedArea, "got:", area)
			}
		})
	}
}
