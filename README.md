# go-standard-library

## Chapter 1

### Retrieving the Golang version

### Accessing program arguments

### Creating a program interface with the flag package

- `go build -o util`
- `util -retry 2 -prefix=example -array=1,2`

### Getting and setting environment variables with default values

- `export DB_CONN=db:/user@example && go run get.go`
- `unset DB_CONN && go run lookup.go`
- `go run main.go`

### Retrieving the current working directory

- `go build -o binary`

### Getting the current process PID

- `go run main.go`

### Handling operating system signals

- `go run main.go`

### Calling an external process

- `go run run.go`
- `go run start.go`

### Retrieving child process information

- `go run main_running.go`

### Reading/writing from the child process

- `go run main_read_output.go`
- `go run main_read_stdout.go`
- `go run main.go`

### Shutting down the application gracefully

- `go run main.go`

### File configuration with functional options

- `go run main.go`
- `export CONN_DB=db:/user@example && go run main.go`

## Chapter 2

### Finding the substring in a string

- `go run contains.go`

### Breaking the string into words

- `go run whitespace.go`
- `go run anyother.go`
- `go run specfunction.go`
- `go run regex.go`

### Joining the string slice with a separator

- `go run join.go`
- `go run join_manually.go`

### Concatenating a string with writer

- `go run concat_buffer.go`
- `go run concat_copy.go`
- `go test -bench=.`

### Aligning text with tabwriter

- `go run tabwriter.go`

### Replacing part of the string

- `go run replace.go`
- `go run replacer.go`
- `go run regexp.go`

### Finding the substring in text by regexp pattern

- `go run regexp.go`

### Decoding a string from the non-Unicode charset

- `go run decode.go`
- `go run encode.go`

### Controlling case

- `go run case.go`

### Parsing comma-separated data

- `go run data.go`
- `go run data_uncommon.go`

### Managing white-space in a string

- `go run whitespace.go`

### Indenting a text document

- `go run main.go`

## Chapter 3

### Converting strings to numbers

- `go run main.go`

### Comparing floating-point numbers

- `go run tolerance.go`
- `go run big.go`

### Rouding floating-point numbers

- `go run round.go`

### Floating-point arithmetics

- `go run main.go`

### Formatting numbers

- `go run format.go`
- `go run localized.go`

### Converting between binary, octal, decimal, and hexadecimal

- `go run convert.go`

### Formatting with the correct plurals

- `go run plurals.go`

### Generating random numbers

- `go run rand.go`

### Operating complex numbers

- `go run complex.go`

### Converting between degrees and radians

- `go run radians.go`

### Taking logarithms

- `go run log.go`

### Generating checksums

- `go run checksum.go`
