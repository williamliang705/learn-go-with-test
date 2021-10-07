package _struct

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name: "rectangle",
			shape: Rectangle{
				Width:  10,
				Height: 10,
			},
			want: 100,
		},
		{
			name: "circle",
			shape: Circle{
				Radius: 10,
			},
			want: 314.1592653589793,
		},
		{
			name: "triangle",
			shape: Triangle{
				Base:   12,
				Height: 6,
			},
			want: 36,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})
	}
}
