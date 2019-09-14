package main

import "fmt"

func main() {

	// the 'iota' key word acts like a counter starting from
	const (
		alpha   = iota
		bravo   = iota
		charlie = iota
	)

	fmt.Println(alpha, bravo, charlie)

	// the iota reset to 0 when moving to the next constang block
	const (
		delta   = iota
		echono  = iota
		foxtrot = iota
	)

	fmt.Println(delta, echono, foxtrot)

}
