package main

import (
	"fmt"
	"reflect"
)

/*
The 'iota' key word acts like a counter starting from 0, it's the equivanlent of
doing "++1". Consequently
the constants default to the integer data type.
*/
const (
	alpha   = iota
	bravo   = iota
	charlie = iota
)

/*
the iota reset to 0 when moving to the next constang block.
Iota is always used in conjunction with constants.
*/
const (
	delta   = iota
	echono  = iota
	foxtrot = iota
)

// the iota reset to 0 when moving to the next constang block
const (
	golf  = iota + 8  // equals to 8
	hotel = iota      // equals to 1
	india = 9 << iota /* iota squared x 9 = 36
						  "<<" is a bitwise operator and means left-shift
	                      https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827
	*/
	juliet = 5 << iota // iota is 3. so 3squared x 5 is 45. But shows as 40, not sure why.
	kilo   = iota      // equals to 3
)

// This starts and continues the counter from a certain number
const (
	lima     = iota + 8 // equals to 8
	mike                // equals 9
	november            // equals 10
)

func main() {

	fmt.Println(alpha, bravo, charlie)  // 0 1 2
	fmt.Println(delta, echono, foxtrot) // 0 1 2
	fmt.Println(reflect.TypeOf(bravo))  // int

	fmt.Println(golf, hotel, india, juliet, kilo) // 8 1 36 40 4
	fmt.Println(lima, mike, november) // 8 9 10

}
