package main

import (
	"reflect"
	"testing"
)

type DriverTestSet struct {
	testName        string
	driverName      string
	tripTimeSeconds int
	tripDistance    float32
	expectedDriver  *Driver
}

var getNewDriverTestSets = []DriverTestSet{
	{"happyPath", "George", 1800, 20, &Driver{Name: "George", MilesDriven: 20, AverageSpeed: 40}},
	{"empty", "", 3600, 60, &Driver{Name: "", MilesDriven: 60, AverageSpeed: 60}},
	{"AlmostSlowTrip", "SlowGuy", 3600, 5, &Driver{Name: "SlowGuy", MilesDriven: 5, AverageSpeed: 5}},
	{"RealSlowTrip", "ReallySlowGuy", 3600, 4, &Driver{Name: "ReallySlowGuy", MilesDriven: 0, AverageSpeed: 0}},
	{"AlmostFastTrip", "FastGuy", 3600, 100, &Driver{Name: "FastGuy", MilesDriven: 100, AverageSpeed: 100}},
	{"RealFastTrip", "ReallyFastGuy", 3600, 101, &Driver{Name: "ReallyFastGuy", MilesDriven: 0, AverageSpeed: 0}},
}

func TestNewDriver(t *testing.T) {
	for _, testSet := range getNewDriverTestSets {
		output := NewDriver(testSet.driverName, testSet.tripTimeSeconds, testSet.tripDistance)
		if !reflect.DeepEqual(output, testSet.expectedDriver) {
			t.Errorf("Test %v didn't match,\nExpected=\n%v\n=====\nGot==\n%v\n=====", testSet.testName, testSet.expectedDriver, output)
		}
	}
}

type UpdateDriverTestSet struct {
	testName        string
	driverName      string
	tripTimeSeconds int
	tripDistance    float32
	originalDriver  *Driver
	expectedDriver  *Driver
}

var getUpdateDriverTestSets = []UpdateDriverTestSet{
	{"happyPath", "George", 900, 5, &Driver{Name: "George", MilesDriven: 20, AverageSpeed: 40}, &Driver{Name: "George", MilesDriven: 25, AverageSpeed: 30}},
	{"tooFast", "GeorgeF", 900, 200, &Driver{Name: "GeorgeF", MilesDriven: 20, AverageSpeed: 40}, &Driver{Name: "GeorgeF", MilesDriven: 20, AverageSpeed: 40}},
	{"tooSlow", "GeorgeS", 900, 1, &Driver{Name: "GeorgeS", MilesDriven: 20, AverageSpeed: 40}, &Driver{Name: "GeorgeS", MilesDriven: 20, AverageSpeed: 40}},
}

func TestUpdateDriver(t *testing.T) {
	for _, testSet := range getUpdateDriverTestSets {
		testSet.originalDriver.UpdateDriver(testSet.driverName, testSet.tripTimeSeconds, testSet.tripDistance)
		if !reflect.DeepEqual(testSet.originalDriver, testSet.expectedDriver) {
			t.Errorf("Test %v didn't match,\nExpected=\n%v\n=====\nGot==\n%v\n=====", testSet.testName, testSet.expectedDriver, testSet.originalDriver)
		}
	}
}
