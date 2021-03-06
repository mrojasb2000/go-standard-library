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
- `go run sha_panic.go`

## Chapter 3

### Finding today's date

Obtaining the current date is a very common task for any system or application.

- `go run today.go`

### Formating date to string

In case the representation textual of a time value is needed, usually, certain formatting is expected. The Time type of the time package provides the ability to create the string output in the given format.

- `go run format.go`

### Parsing the string into date

The same concept as the one used in date formatting is also used by date parsing. The same reference date and layout principles can be used.

- `go run parse.go`

### Converting dates to epoch and vice versa

The epoch es the universal system to describe the point in time. The beginning of epoch time is defined as 00:00:00 1 Jan 1970 UTC. The value of epoch is the amount of seconds since the timestamp, minus the amount of leap seconds since then.

- `go run epoch.go`

### Retrieve time units from the date

The Time type also provides the API to retrieve time units from the instance. The means you are to find out what day in a month or what hour in a day the instance represents.

- `go run units.go`

### Date arithmetics

The Time type of the time package also allows you to perform basic arithmetic on the given date and time. This way, you can find out past and future dates.

- `go run arithmetics.go`

### Finding the difference between two dates

Finding the difference between two dates is not an unusual task. For the operation, the Go standard package time, respectively the Time type, provides supporting methods.

### Converting between time zones

Dealing with time zones is hard. A good way to handle the different time zones is keep one timezone as referential in the system and convert the others if needed.

- `go run timezones.go`

### Running the code block periodically

Basides the code and time operations, the time package also provides support for periodic and delayed code execution. Typicaly, the application health cheks, activy checks, or many periodic job can be implemented this way.

`go run ticker.go`

### Waiting a certain amount of time

The previous recipe descrives how to execute the code periodically. This recipe will show you how to to execute the code with a delay.

`go run delay.go`

### Timeout long-running operations

The previous recipe decribes the concept of executing the code with some delay. The same concept can be used to implement the timeout for long running operations.

`go run timeout.go`

### Serializing the time and date

When serializing the date and time information, it is necesary tho choose the proper format. This recipe will illustrate how the time package helps to choose one and do the serialization properly.

`go run serialize.go`

## Chapter 5

### Reading standard input

Every process owns its standard input, output, and error file descriptor. The stdin serves as the input of the process. This recipe describes how to read the data from the stdin.

`go run fmt.go`
`go run scanner.go`
`echo 'Go is awesome!' | go run reader.go`

### Writing standard output and error

As the previous recipe describes, each process has stdin, a stdout and stderr file descriptors. The standard approach is the use of stdout as a process output and stderr as process error output. As these are the file descriptors, the destination where the data is written could be anything, from the console to the socker. This recipe will show you how to write the stdout and sterr.

`go run stdouterr.go`

### Opening a file by name

File access is a veru common operation used to store or read the data. This recipe illustrates how to open a file by its name and path, using the standard library.

`go run openfile.go`

### Reading the file into a string

In the previous recipes, we saw the reading from Stdin and
the opening of the file. In this recipe, we will combine these two a little bit and how to read the file into a string.

`go run readfile.go`

### Reading/writing a different charset

It is not an exception that the input from various sources could come in various charsets. Note that a lot of systems use the Windows operating system but there are others. Go, by default, expects that the string used in the program are UTF-8 based. If they are not, then decoding from the given charset must be deno to be able to work with the string. This recipe will show the reading and writing of the file in a charset other than UTF-8.

`go run charset.go`

### Seeking a position within a file

In some cases, you need to read from or write to a particular location in a file, such as an indexed file. The recipe will show you how to use the position seeking in the context of flat file operations.

`go run fileseek.go`
`xxd temp/flatfile.txt`


### Reading and writing binry data

This recipe describes how to write and read any type in then binary form.

`go run rwbinary.go`


### Wrting to multiple writers at once

When you need to write the same output into maro than one target, there is a helping hand available in the built-in package. The recipe show how to implement writing simultaneously into multiple targets.

`go run multiwr.go`


### Piping between writer and reader

The pipes between processes are the easy way to use the output of the first process as the input of other processes. The same concept could be done in Go, for example, to pipe data from one socket to another socket, to create the tunneled connection. This recipe will show you how to to create the pipe with use of the Go built-in library.

`go run pipe.go`


### Serializing objects to binary format

Besides the well-known JSON and XML, Go also offers the binary format, gob. This recipe goes throught the basic concept of how to the gob package.

`go run gob.go`

### Reading and writing ZIP files

ZIP compression is a widely used compression format. It is usual to use the ZIP format for an application to upload set or, on the other hand, export zipped files as output. This recipe will show you how to handle ZIP files programmatically with the use of the standard library.

`go run zip.go`