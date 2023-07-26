package alib

import "testing"

var Debug bool = false

func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("スキップ")
	}
	v := Average([]int{1,2,3,4,5, 10})
	if v != 3 {
		t.Errorf("%v != %v", v, 3)
	}
}

// $ go test ./alib
// ok      udemy-todo/alib 2.871s

// $ go test -v ./alib
// === RUN   TestAverage
// --- PASS: TestAverage (0.00s)
// PASS
// ok      udemy-todo/alib 0.586s

// $ go test -v ./alib
// === RUN   TestAverage
//     alib_test.go:13: 4 != 3
// --- FAIL: TestAverage (0.00s)
// FAIL
// FAIL    udemy-todo/alib 2.163s
// FAIL

// $ go test -v ./... すべてのパッケージのテストを行うことができる
// ?       udemy-todo/foo  [no test files]
// === RUN   TestIsOne
// --- PASS: TestIsOne (0.00s)
// PASS
// ok      udemy-todo      2.505s
// === RUN   TestAverage
//     alib_test.go:13: 4 != 3
// --- FAIL: TestAverage (0.00s)
// FAIL
// FAIL    udemy-todo/alib 0.527s
// FAIL
