# Mutexes
Mutexes in golang offer synchronization utility.

Mutexes help prevent race conditions when multiple go routines try to read and write the same elemet at the same time.

# How to run test files:

for test run -
go test -v main.go main_test.go

for Race Condition Test, run - 
go test main.go main_test.go --race


