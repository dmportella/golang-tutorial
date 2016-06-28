package chapter5

import "testing"

func TestFibonacci(t *testing.T) {
	input := 7
	expected := 13

	actual := Fibonacci(input)

	if actual != expected {
		t.Errorf("Fibonacci(%d): expected %d, actual %d", input, expected, actual)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(10)
	}
}
