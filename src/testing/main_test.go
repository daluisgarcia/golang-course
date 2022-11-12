package main

import "testing"

// To execute the test, run the following command: go test

//* Test Coverage *//
// You can also see the coverage of the test by running the following command: go test -cover
// And also you can see which lines are not covered by the test by running the following command: go test -coverprofile=coverage.out
//                                                                                   For Windows: go test --coverprofile=coverage.out
// And then you can see the coverage by running the following command: go tool cover -html=coverage.out or go tool cover -func=coverage.out
//                                                        For Windows: go tool cover --html=coverage.out or go tool cover --func=coverage.out

func TestSum(t *testing.T) {
	// if Sum(5, 5) != 10 {
	// 	t.Error("Sum(5, 5) != 10")
	// }

	tables := []struct {
		a int
		b int
		n int
	}{
		{5, 5, 10},
		{1, 2, 3},
		{2, 2, 4},
	}

	for _, testCase := range tables {
		total := Sum(testCase.a, testCase.b)
		if total != testCase.n {
			t.Errorf("Sum(%d, %d) != %d", testCase.a, testCase.b, testCase.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	// if GetMax(5, 5) != 5 {
	// 	t.Error("GetMax(5, 5) != 5")
	// }

	tables := []struct {
		a int
		b int
		n int
	}{
		{6, 5, 6},
		{3, 2, 3},
		{8, 6, 8},
	} // This test cases does not cover all the possible cases of the function

	for _, testCase := range tables {
		max := GetMax(testCase.a, testCase.b)
		if max != testCase.n {
			t.Errorf("GetMax(%d, %d) != %d", testCase.a, testCase.b, testCase.n)
		}
	}
}

//* Code Profiling *//
// You can see the profiling of the code by running the following command: go test -cpuprofile=cpu.out
//                                                            For Windows: go test --cpuprofile=cpu.out
// And then you can see the profiling by running the following command: go tool pprof cpu.out
// In the cli of the pprof, you can write the command: top. You can see the top 10 functions that take the most time.
// To see the report of a certain function, you can write the command: list <function_name>
// To see the graph of the profiling, you can write the command: web. You can see the graph of the profiling.
// To create a pdf file of the graph, you can write the command: pdf. You can see the pdf file of the graph.

// You can also see the profiling report diagram in the browser: go tool pprof -web cpu.out
//                                For Windows: go tool pprof --web cpu.out

func TestFibonacci(t *testing.T) {

	tables := []struct {
		n int
		f int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, testCase := range tables {
		fib := Fibonacci(testCase.n)
		if fib != testCase.f {
			t.Errorf("Fibonacci(%d) != %d", testCase.n, testCase.f)
		}
	}
}
