package main

import "fmt"

// the 'iota' key word acts like a counter starting from
const (
	alpha   = iota
	bravo   = iota
	charlie = iota
)

// the iota reset to 0 when moving to the next constang block
const (
	delta   = iota
	echono  = iota
	foxtrot = iota
)

func main() {

	fmt.Println(alpha, bravo, charlie)
	fmt.Println(delta, echono, foxtrot)
}
