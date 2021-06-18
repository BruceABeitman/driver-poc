package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type DriverTracker struct {
	Drivers      map[string]*Driver `json:"drivers"`
	ValidDrivers map[string]bool    `json:"validDrivers"`
}

// Creates a new DriverTracker
func NewDriverTracker() *DriverTracker {
	return &DriverTracker{
		Drivers:      make(map[string]*Driver),
		ValidDrivers: make(map[string]bool, 0),
	}
}

// Handles parsing a Driver command
func (driverTracker *DriverTracker) ParseDriver(input []string) error {
	if len(input) != 2 {
		return errors.New(fmt.Sprintf("Invalid Driver command. Found %v arguments, requires 2.", len(input)))
	}
	driverName := input[1]
	driverTracker.ValidDrivers[driverName] = false
	return nil
}

// Handles parsing a Trip command
func (driverTracker *DriverTracker) ParseTrip(input []string) error {
	if len(input) != 5 {
		return errors.New(fmt.Sprintf("Invalid Trip command. Found %v arguments, requires 5.", len(input)))
	}
	driverName := input[1]
	tripStart := input[2]
	tripEnd := input[3]
	tripDistance := input[4]

	timeStart, err := time.Parse("15:04", tripStart)
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid time value: %v. Error: %v", tripStart, err))
	}
	timeEnd, err := time.Parse("15:04", tripEnd)
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid time value: %v. Error: %v", tripEnd, err))
	}

	tripDurationSeconds := int(timeEnd.Sub(timeStart) / time.Second)
	fmtTripDistance, err := strconv.ParseFloat(tripDistance, 32)
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid trip distance: %v. Error: %v", tripDistance, err))
	}

	if driver, ok := driverTracker.Drivers[driverName]; ok {
		// Update existing driver
		driver.UpdateDriver(driverName, tripDurationSeconds, float32(fmtTripDistance))
	} else {
		// Add new driver
		driverTracker.Drivers[driverName] = NewDriver(driverName, tripDurationSeconds, float32(fmtTripDistance))
	}
	return nil
}
