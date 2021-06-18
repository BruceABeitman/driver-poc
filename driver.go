package main

import (
	"fmt"
	"math"
)

type Driver struct {
	Name         string  `json:"name"`
	MilesDriven  float32 `json:"miles"`
	AverageSpeed float32 `json:"mph"`
}

// Creates a new driver
func NewDriver(driverName string, tripTimeSeconds int, tripDistance float32) *Driver {
	tripTimeHours := float32(tripTimeSeconds) / float32(3600)
	tripSpeed := tripDistance / tripTimeHours
	// Discard any trips that average a speed of less than 5 mph or greater than 100 mph.
	if tripSpeed < 5 || tripSpeed > 100 {
		return &Driver{
			driverName,
			0,
			0,
		}
	}
	return &Driver{
		driverName,
		tripDistance,
		tripSpeed,
	}
}

// Updates a driver with new information
func (driver *Driver) UpdateDriver(driverName string, tripTimeSeconds int, tripDistance float32) {
	tripTimeHours := float32(tripTimeSeconds) / float32(3600)
	tripSpeed := tripDistance / float32(tripTimeHours)
	// Discard any trips that average a speed of less than 5 mph or greater than 100 mph.
	if tripSpeed < 5 || tripSpeed > 100 {
		return
	}
	driver.MilesDriven += tripDistance
	driver.AverageSpeed = (driver.AverageSpeed + tripSpeed) / 2
}

// Generates a string for printing the driver data
func (driver *Driver) PrintDriver() string {
	return fmt.Sprintf(
		"%v: %v miles @ %v mph",
		driver.Name,
		math.Round(float64(driver.MilesDriven)),
		math.Round(float64(driver.AverageSpeed)),
	)
}
