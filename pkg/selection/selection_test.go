package selection

import "testing"

func TestSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6}
	actual := Sort([]int{3, 1, 2, 5, 4, 6})

	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Test failed, expected: %v, got: %v", expected, actual)
			break
		}
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort([]int{3, 1, 2, 5, 4, 6})
	}
}
