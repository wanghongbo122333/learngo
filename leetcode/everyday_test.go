package every

import "testing"

func Test_findDuplicate(t *testing.T) {
	var (
		in       = []int{1, 2, 3, 4, 5, 1}
		expected = 1
	)
	actual := findDuplicate(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}

}
