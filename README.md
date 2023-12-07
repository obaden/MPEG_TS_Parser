# MPEG_TS_Parser
MPEG Transport Stream parser which saves the payload and ensures the stream of a valid format

## How to use

Simply build the code using
```
go build
```

then you can pipe in any file you want using 
```
cat example.ts | ./MPEG_TS_Parser.exe
```

This will give you an error if the example file contains an invalid stream.
If the stream is valid, it will instead print out a list of all the PIDs in hex padded to 4 hex digits

## Tests

to run tests make sure you are in the parser_tests directory and then simply run
```
go test
```
