package main

import "testing"

/*
go test -cover => basic sknowledge
go test -coverprofile=exampletest.out => create a file to know more
	to rum the file created: go tool cover -func=exampletest.out
	to run and see in a browser hmtl: go tool cover -hmtl=exampletest.out
*/
func TestSum(t *testing.T) {
	// total := sum(5, 5)
	// if total != 11 {
	// 	t.Errorf("Sum was incorrect, go %d but expected %d ", total, 11)
	// }
	tables := []struct {
		a int
		b int
		r int
	}{
		{1, 2, 3}, {2, 2, 4}, {25, 26, 51},
	}

	for _, item := range tables {
		total := sum(item.a, item.b)
		if total != item.r {
			t.Errorf("Sum was incorrect, go %d but expected %d", total, item.r)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		r int
	}{
		{4, 8, 8},
		{3, 3, 3},
		{2, 5, 5},
	}
	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.r {
			t.Errorf("Max was incorrect, got %d but expected %d", max, item.r)
		}
	}
}

/*
PROFILING
go test -cpuprofile=cou.out => create a file takes a long time
	to run the file created: go tool pprof cou.out
		% top => shows the details performance
		% list Fibonacci => shows the details performance for Fibonacci[test that you need to see datils in code]
		% web => create a file with more information
		% pdf => generate a PDF file with information about test results - report performance
	to run and see in a browser hmtl: go tool cover -hmtl=exampletest.out
*/
func TestFib(t *testing.T) {
	tables := []struct {
		a int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.r {
			t.Errorf("Fibonacci was incorrect, got %d but expected %d", fib, item.r)
		}
	}
}
