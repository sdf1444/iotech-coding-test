package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestMain(t *testing.T) {
	// Read the output from output.json
	output, err := ioutil.ReadFile("output.json")
	if err != nil {
		t.Error("Error reading output.json:", err)
	}

	// Unmarshal the output into an Output struct
	var outputStruct Output
	json.Unmarshal(output, &outputStruct)

	// Check that the ValueTotal is correct
	expectedValueTotal := 17 + 50 + 33
	if outputStruct.ValueTotal != expectedValueTotal {
		t.Errorf("Incorrect ValueTotal. Got %d, expected %d", outputStruct.ValueTotal, expectedValueTotal)
	}

	// Check that the UUIDS slice is correct
	expectedUUIDS := []string{
		"22919442-e583-11ec-8fea-0242ac120002",
		"29446300-e583-11ec-8fea-0242ac120002",
		"2d9596e0-e583-11ec-8fea-0242ac120002",
	}
	if !sliceEqual(outputStruct.UUIDS, expectedUUIDS) {
		t.Errorf("Incorrect UUIDS. Got %v, expected %v", outputStruct.UUIDS, expectedUUIDS)
	}
}

// Helper function to check if two slices are equal
func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
