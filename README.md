# Sample code and data for a bug in the fixedwidth package for go

This repository supports an issue report for this repository: https://github.com/ianlopshire/go-fixedwidth/issues/new. The issue is that the Decoder will return `io.EOF` (rather than some error) prior to the end of the file if there is a line that exceeds the Scanner buffer size. 

The file `test_file.txt` contains a sample file that would cause this error. The file `main.go` contains the sample code, based off of the documentation/examples provided for the fixedwidth package.
