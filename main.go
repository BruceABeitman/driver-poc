package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	// Validate we were given a file
	cmdArgs := os.Args
	if len(cmdArgs) < 2 {
		fmt.Println("Missing required filename. Exiting.")
		return
	}
	filename := cmdArgs[1]

	// Validate we can find the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v. Exiting.\n", err.Error())
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Scan over each line in the file and parse driver data
	lineNumber := 0
	driverTracker := NewDriverTracker()
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) == 0 {
			fmt.Printf("Ignoring empty or incorrectly formatted line(%v)\n", lineNumber)
		}
		command := words[0]
		switch command {
		case "Driver":
			err := driverTracker.ParseDriver(words)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "Trip":
			err := driverTracker.ParseTrip(words)
			if err != nil {
				fmt.Println(err.Error())
			}
		default:
			fmt.Printf("Found unsupported command '%v'. Ignoring line(%v).\n", command, lineNumber)
		}
		lineNumber++
	}
	// Handle file reading errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file: " + err.Error())
	}

	// Sort Driver data
	sortedDrivers := make(SortedDrivers, len(driverTracker.Drivers))
	index := 0
	for _, Driver := range driverTracker.Drivers {
		sortedDrivers[index] = Driver
		index++
	}
	sort.Sort(sortedDrivers)

	// Print sorted, valid, populated Driver data
	// Note, this excludes valid non-populated drivers (e.g. drivers without trip data)
	for _, driver := range sortedDrivers {
		if _, ok := driverTracker.ValidDrivers[driver.Name]; ok {
			fmt.Println(driver.PrintDriver())
			// outputFile.WriteString(fmt.Sprintf("%v\n", driver.PrintDriver()))
			// Mark we've visited the valid driver
			driverTracker.ValidDrivers[driver.Name] = true
		}
	}

	// Print remaining valid drivers not yet visited
	// Note, these will all be 0 miles as we do not have trip data for them
	for driverName, visited := range driverTracker.ValidDrivers {
		if !visited {
			fmt.Printf("%v: 0 miles\n", driverName)
			// outputFile.WriteString(fmt.Sprintf("%v: 0 miles\n", driverName))
		}
	}
}
