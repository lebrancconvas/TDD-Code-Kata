package kata01

import "testing"

type TestCase struct {
	Name string
	Input string
	Want string
}

func TestDNAtoRNA(t *testing.T) {
	testCases := []TestCase{
		{Name: "GCAT should return GCAU", Input: "GCAT", Want: "GCAU"},
		{Name: "ACTG should return ACUG", Input: "ACTG", Want: "ACUG"},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			got := DNAtoRNA(tc.Input)

			if got != tc.Want {
				t.Errorf("[ERROR] %v: got %v, but want %v", tc.Name, got, tc.Want)
			}
		})
	}
}
