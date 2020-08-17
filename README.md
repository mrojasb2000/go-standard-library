# go-standard-library

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

