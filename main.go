package main

import "fmt"

func main() {

	const pie = 3
	fmt.Println(pie)

	// note, you can't initialise and then declare constants, i.e. the following won't work. Hence
	// you only declare constants and the initialisation is done for you by golang, behind the scenes.
	// const pi := 3.14

	// we can't change constants, e.i. the following will fail
	// pi = 3.141592

	// constants can't be created during run time. That means constants definition has to be hardcoding in your code, so that
	// they get baked-in during compile time (ie. during the creation of the binary).

	// constants can be used with different data types, e.g. as an integer or float
	fmt.Println(pie + 1)     // works as an integer
	fmt.Println(pie + 0.141) // used as a float

	// however you can restrict constant to be used as a particular data type
	const pi int = 2
	fmt.Println(pi + 1)

	// however this will fail
	// fmt.Println(pi + 1.141)

	// a workaround is by doing a conversion:
	fmt.Println(float32(pi) + 1.141)
}
