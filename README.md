# Driver-POC

This is a POC for tracking driver statistics from a file. It reads a file line by line looking for 2 keywords `Driver` & `Trip`. Example format:

- `Driver Kumi`
- `Trip Dan 07:15 07:45 17.3`

Each line is parsed by splitting on one (or more consecutive) whitespace characters, and parsing each line accordingly.

- `Driver` `{driver name}`
- `Trip` `{driver name}` `{start time}` `{stop time}` `{miles driven}`

_Note that any trips that average a speed of less than 5 mph or greater than 100 mph are discarded._

Finally, **valid** Drivers are printed in descending order of miles driven to the console in the following format:

- `{driver_name}: {total miles driven} miles @ {average driver speed} mph`

Example: `Lauren: 42 miles @ 34 mph`

Any **valid** driver that has no data is printed as such:

`Kumi: 0 miles`

### Valid Driver Details

A valid driver and how it is handled in the code is a little complex. A valid driver is delineated by having a dedicated `Driver` command in the input file. We obviously have a complication around a `Trip` command for an invalid driver or vice-versa.

The way the code approaches this is treating each type of command separately. This is done to allow us to gather all the information with a single pass over the file. When parsing a `Driver` command, the name is kept in a `validDriver` map as a key and a boolean is populated as a value (this is useful later). When parsing a `Trip` command, the name is kept in a `driverMap` as a key (regardless of it being "valid" or not) and all data is tracked in a `Driver Object` as a value. At the end, we sort the `Driver Object`s and print their data if they exist in the `validDriver` map. Finally (using that boolean from the `validDriver` map) we print any `validDriver`s we did not have any `Trip` data for, and print generic data for them.

_Note when printing the valid non-populated driver data the ordering will be random due to Golang's map implementation. While this meets requirements, creating tests for this might be complex._

# Getting it running

This was built in `go1.12.1`, although is expected to be compatible with most go versions.

To build, navigate to this project's directory and run:

`go build ./...`

`./driver-poc input.txt`

And to test:

`go test ./...`

# Next Steps

This code is a good foundation, but additional considerations should be made for extension and production prep.

- The current tests overall are rather minimal. Additional testing should be considered for typos and incorrect formatting in the input file.
- The `main` file has some core functionality that is not currently covered in tests. This file should be covered in an integration test. Additionally, some of the `main` functionality could be moved depending on how the project is expected to be extended (in particular the keyword parsing).
- The `sortedDrivers` is currently not covered by tests either. It can be relatively easily tested via local integration, but should be considered for unit tests especially if expanding the features of the `Driver` model.
