package gocubicsolver

import (
	"fmt"
	"math"
	"testing"
)

func TestSolve(t *testing.T) {
	fmt.Println("Cuadratic")
	got, _ := Solve(0, 1, 1, -1)
	want := []float64{0.6180339887498949, -1.618033988749895}
	check(got, want, t, "1")

	fmt.Println("Cuadratic")
	got, _ = Solve(0, 2, 1, -1)
	want = []float64{0.5, -1}
	check(got, want, t, "2")

	fmt.Println("Cubic")
	got, _ = Solve(1, -6, 11, -6)
	want = []float64{3, 1, 2}
	check(got, want, t, "3")

	fmt.Println("Cubic")
	got, _ = Solve(2, -4, -22, 24)
	want = []float64{4, -3, 1}
	check(got, want, t, "4")

	fmt.Println("Cubic")
	got, _ = Solve(1, -3, 3, -1)
	want = []float64{1, 1, 1}
	check(got, want, t, "5")
}
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve(1, -3, 3, -1)
	}
}
func check(got, want []float64, t *testing.T, s string) {
	fmt.Println("Got:", got)
	fmt.Println("Want:", want)
	for i := 0; i < len(want); i++ {
		if math.Abs(got[i]-want[i]) > 1e-3 {
			t.Errorf("got %v, wanted %v, Case %s", got[i], want[i], s)
		}
	}
}
