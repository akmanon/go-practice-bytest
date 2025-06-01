package structmethodinterface

import (
	"testing"
)

func TestArea(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Circle", Circle{radius: 5.0}, 78.53981633974483},
		{"Rectangle", Rectangle{height: 5.0, width: 7.0}, 35.0},
		{"Triangle", Triangle{base: 12.0, height: 6.0}, 36.0},
	}
	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %.2f wnat %.2f", tt.shape, got, tt.hasArea)
		}
	}
}
