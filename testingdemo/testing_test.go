package testingdemo

import "testing"

// go test -v
// Use the path ./... to test subdirectories recursively:
// go test -v ./...
// To run a specific test, use the -run flag with a regex matching the test name:
// go test -v -run TestSum
func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Sum(2, 3) = %d; want %d", result, expected)
	}
}

// Test methods start with Test
func TestSumWithTable(t *testing.T) {
	// Note that the data variable is of type array of anonymous struct,
	// which is very handy for writing table-driven unit tests.
	data := []struct {
		a, b, res int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{1, -1, 0},
		{2, 3, 5},
		{1000, 234, 1234},
	}

	for _, d := range data {
		if got := Sum(d.a, d.b); got != d.res {
			t.Errorf("Sum(%d, %d) == %d, want %d", d.a, d.b, got, d.res)
		}
	}
}

// go test -bench=.
func BenchmarkSum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = Sum(2, 3)
    }
}

// Run go test as normal, yet with the coverprofile flag. Then use go tool to view the results as HTML.
// go test -coverprofile=c.out
// go tool cover -html=c.out