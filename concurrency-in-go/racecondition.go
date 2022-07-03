/**
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.
Submission: Upload your source code for the program along with your written explanation of race conditions.
*/
package main

import (
	"fmt"
	"time"
)

var ctr int

/*
	Interleaving with this instruction is non-deterministic. 'ctr++' is not an atomic instructions at the machine code level which
	consist of READ, WRITE and SAVE the value of 'ctr'.

	In other words, the access to this variable will cause race condition in the sense where the first thread has incremented the value of 'ctr'
	while the other thread might be reading the old value.

	This might result in a value of 1 or 2.

*/
func inc() {
	ctr++

}

func main() {
	go inc()
	go inc()
	time.Sleep(1 * time.Second)
	fmt.Println(ctr)
}

/**
% go run -race racecondition.go
==================
WARNING: DATA RACE
Read at 0x000001277d08 by goroutine 8:
  main.inc()
      ..go/src/coursera/racecondition.go:11 +0x3a

Previous write at 0x000001277d08 by goroutine 7:
  main.inc()
      ..go/src/coursera/racecondition.go:11 +0x56

Goroutine 8 (running) created at:
  main.main()
      ..go/src/coursera/racecondition.go:16 +0x5e

Goroutine 7 (finished) created at:
  main.main()
      ..go/src/coursera/racecondition.go:15 +0x46
==================
2
Found 1 data race(s)
exit status 66
*/
