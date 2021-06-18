package main

// Provide a sorted Driver array
type SortedDrivers []*Driver

func (sDrivers SortedDrivers) Len() int {
	return len(sDrivers)
}

func (sDrivers SortedDrivers) Less(i, j int) bool {
	return sDrivers[i].MilesDriven < sDrivers[j].MilesDriven
}

func (sDrivers SortedDrivers) Swap(i, j int) {
	sDrivers[i], sDrivers[j] = sDrivers[j], sDrivers[i]
}
