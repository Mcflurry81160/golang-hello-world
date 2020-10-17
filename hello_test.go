package main

import "testing"

func TestSampleFunctionToTest(t *testing.T) {
	actualResult, err := sampleFunctionToTest()

	if err != nil {
		t.Errorf("Received error")
	}

	if actualResult != "abc" {
		t.Errorf("Did not return expected result")
	}
}