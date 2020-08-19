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