package testset

type TestSet struct {
	tests []string
}

func Create(tests []string) *TestSet {
	return &TestSet{
		tests: tests,
	}
}

func (ts *TestSet) Split(n int) [][]string {
	result := [][]string{}

	for i := 0; i < n; i++ {
		result = append(result, []string{})
	}

	for i, test := range ts.tests {
		idx := i % n
		result[idx] = append(result[idx], test)
	}

	return result
}
