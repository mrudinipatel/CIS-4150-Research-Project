package domain

import "math/rand"

type TestSet struct {
	tests []string
}

func NewTestSet(tests []string) *TestSet {
	return &TestSet{
		tests: tests,
	}
}

func (ts *TestSet) Split(n int) [][]string {
	result := make([][]string, n)

	for i := 0; i < n; i++ {
		result = append(result, []string{})
	}

	tests := make([]string, len(ts.tests))

	rand.Shuffle(len(ts.tests), func(i int, j int) {
		tests[i] = ts.tests[j]
	})

	for i, test := range tests {
		idx := i % n
		result[idx] = append(result[idx], test)
	}

	return result
}
