package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type DriverTrackerSet struct {
	testName string
	input    []string
	expError error
	original *DriverTracker
	expected *DriverTracker
}

var getNewDriverTrackerTestSets = []DriverTrackerSet{
	{
		testName: "HappyPath",
		input:    []string{"Driver", "George"},
		expError: nil,
		original: NewDriverTracker(),
		expected: &DriverTracker{Drivers: make(map[string]*Driver), ValidDrivers: map[string]bool{"George": false}},
	},
	{
		testName: "InvalidArguments",
		input:    []string{"George"},
		expError: errors.New(fmt.Sprintf("Invalid Driver command. Found 1 arguments, requires 2.")),
		original: NewDriverTracker(),
		expected: NewDriverTracker(),
	},
}

func TestParseDriver(t *testing.T) {
	for _, testSet := range getNewDriverTrackerTestSets {
		errOutput := testSet.original.ParseDriver(testSet.input)
		if !reflect.DeepEqual(testSet.expected, testSet.original) {
			t.Errorf("Test %v didn't match,\nExpected=\n%v\n=====\nGot==\n%v\n=====", testSet.testName, testSet.expected, testSet.original)
		}
		if !reflect.DeepEqual(errOutput, testSet.expError) {
			t.Errorf("Test Error %v didn't match,\nExpected=\n%v\n=====\nGot==\n%v\n=====", testSet.testName, testSet.expError, errOutput)
		}
	}
}

var getParseTripTestSets = []DriverTrackerSet{
	{
		testName: "HappyPath",
		input:    []string{"Trip", "George", "07:15", "07:45", "17.3"},
		expError: nil,
		original: NewDriverTracker(),
		expected: &DriverTracker{
			Drivers: map[string]*Driver{
				"George": {
					Name:         "George",
					MilesDriven:  17.3,
					AverageSpeed: 34.6,
				},
			},
			ValidDrivers: make(map[string]bool, 0),
		},
	},
	{
		testName: "InvalidArguments",
		input:    []string{"George"},
		expError: errors.New(fmt.Sprintf("Invalid Trip command. Found 1 arguments, requires 5.")),
		original: NewDriverTracker(),
		expected: NewDriverTracker(),
	},
}

func TestParseTrip(t *testing.T) {
	for _, testSet := range getParseTripTestSets {
		errOutput := testSet.original.ParseTrip(testSet.input)
		if !reflect.DeepEqual(testSet.expected, testSet.original) {
			t.Errorf("Test %v didn't match,\nExpected=\n%v\n=====\nGot==\n%v\n=====", testSet.testName, testSet.expected, testSet.original)
		}
		if !reflect.DeepEqual(errOutput, testSet.expError) {
			t.Errorf("Test Error %v didn't match,\nExpected=\n%v\n=====\nGot==\n%v\n=====", testSet.testName, testSet.expError, errOutput)
		}
	}
}
