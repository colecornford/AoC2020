package day1

import "testing"

func Test( t *testing.T) {
	
	testData := make(map[int]int)

	for test, expected := range testData {
		result := doFunc(test)
		if result != expected {
			t.Errorf("Failed on: %v. Got: %v. Expected: %v", test, result, expected)
		}
	}
}

