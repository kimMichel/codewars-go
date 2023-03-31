package pets

import "testing"

func TestCalculateYears(t *testing.T) {

	verifyError := func(t *testing.T, result, expect [3]int) {
		t.Helper()
		if expect != result {
			t.Errorf("result '%v', expect '%v'", result, expect)
		}
	}

	t.Run("one year", func(t *testing.T) {
		expect := CalculateYears(1)
		result := [3]int{1, 15, 15}

		verifyError(t, result, expect)
	})

	t.Run("two year", func(t *testing.T) {
		expect := CalculateYears(2)
		result := [3]int{2, 24, 24}

		verifyError(t, result, expect)
	})

	t.Run("ten year", func(t *testing.T) {
		expect := CalculateYears(10)
		result := [3]int{10, 56, 64}

		verifyError(t, result, expect)
	})

}

func BenchmarkCalculateYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateYears(10)
	}
}
