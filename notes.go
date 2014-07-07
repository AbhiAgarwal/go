// Every Go program must start with a package declaration. Packages are Go's way of organizing and reusing code.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func firstPart() {
	// Newlines, spaces and tabs are known as whitespace
	// Hello, Go
	fmt.Println("Hello, Go")

	// * Numbers
	// Integers
	// Unsigned Integer: uint8, uint16, uint32, uint64, int8, int16, int32 and int64
	// Signed Integer: int
	// byte, uint8, rune, same as int32
	// Go's byte data type is often used in the definition of other types
	// There are also 3 machine dependent integer types: uint, int and uintptr. They are machine dependent because their size depends on the type of architecture you are using

	// Floating Point Numbers
	// single precision and double precision: float32, float64
	// complex numbers: complex64 and complex128

	// Casting works in the same way as any other language
	// float64(len(x))

	fmt.Println("1 + 1 =", 1+1)

	// This is because the character is represented by a byte (remember a byte is an integer). Therefore this would return 72
	fmt.Println("Hello World"[0])

	// Variables can be declared in the following fashion
	var a = "Hello"
	b := " "
	var c string = "World"
	var d string
	d = "."

	fmt.Println(a + b + c + d)

	// Using some of the things we discovered earlier
	// Unsigned Integer 8: The highest it can go is 2^8 - 1 == 255
	var e uint8 = 255
	fmt.Println(e)

	// Declaring constants in this manner
	const f string = "Hello World"
	const g = "X"

	// Declaring multiple variables
	var (
		h = "X"
		i = "Y"
		j = "Z"
	)
	fmt.Println(h + i + j)

	// Taking input
	// Float64 has largest range for floating point so we'll use that
	// incase the user does 2023023203.23
	fmt.Print("Enter a number: ")
	var input float64 = 2
	//fmt.Scanf("%f", &input)
	fmt.Println(input)

	// Time
	// There's Location, Minute, Month, Nanosecond, Second, Year, Weekday, Hour
	t := time.Now()
	fmt.Println("The time is", t)
	fmt.Println("The hour is", t.Hour())
	fmt.Println("The year is", t.Year())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// Control Structures - For
	k := 1
	for k <= 10 {
		k = k + 1
	}

	// Array - fixed length *IMPORTANT here.
	// length of array is len(ARRAY)
	var A [5]int
	fmt.Println("Array", A)
	fmt.Println("Length of array", len(A))

	B := [5]float64{98, 93, 77, 82, 83}
	fmt.Println("Array", B)

	// Slices - A slice is a segment of an array.
	// Like arrays slices are indexable and have a length.
	// *** Unlike arrays this length is allowed to change.
	// Syntax wise they just don't have a size

	var C []float64

	// Make a slice of size 5
	D := make([]float64, 5)

	// Create a slice of size 5, where the underlying array is of size 10
	// Slices are basically pointers pointing to intervals in arrays
	// Say starting pointer points to 0 and ending pointer points to 4
	// But there are still 10 elements in the array
	E := make([]float64, 5, 10)

	fmt.Println(C, D, E)

	// See here where the underlying array is of size 5
	// But using the ":" command we are only selecting the first 2 elements
	// Imagine this being: START:LENGTH
	// START: where to start in the array - the starting index
	// LENGTH: how long you want it to be, ADDING the START
	// so LENGTH = LENGTH YOU WANT + START
	// LENGTH is the index where to end it (but not including the index itself)
	underlyingArray := [5]float64{1, 2, 3, 4, 5}
	F := underlyingArray[0:2]
	fmt.Println(F)

	// If you just do underlyingArray[0:]
	// It will just go from the element specified to the end
	G := underlyingArray[2:]
	fmt.Println(G)

	// Go includes two built-in functions to assist with slices: append and copy.
	// Append creates a new slice by taking an existing slice (the first argument) and appending all the following arguments to it.

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("Slice1", slice1,
		"Slice2", slice2)

	// Copy
	// Since slice2 has room for only two elements only the first two elements of slice1 are copied.

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println("Slice3", slice3,
		"Slice4", slice4)

	// Maps
	// A map is an unordered collection of key-value pairs.
	// Also known as an associative array, a hash table or a dictionary, maps are used to look up a value by its associated key.

	// var H map[string]int
	// var x map[string]int
	// The following will give you: assignment to entry in nil map
	//x["s"] = 4

	// H is a map of strings to ints
	// you've to "make" it
	H := make(map[string]int)

	H["key"] = 10
	fmt.Println(H)

	// Delete function
	delete(H, "key")

	// To check if the element exists
	value, ok := H["key"]
	fmt.Println(value, ok)

	if ok {
		fmt.Println("Exists!")
	}

	// We can also go crazy:
	// map of strings to maps of strings to strings
	I := make(map[string](map[string]string))

	I["H"] = map[string]string{
		"name":  "Hydrogen",
		"state": "gas",
	}

	fmt.Println(I["H"]["name"], I["H"]["state"])
}

// function's signature:
// func, followed by the function's name
// parameters: name type, name type
// return type

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Go is also capable of returning multiple values from a function
// Get return back like: x, y := f()
func f() (int, int) {
	return 5, 6
}

// Variadic Function
// There is a special form available for the last parameter in a Go function
func add(args ...int) int {
	fmt.Println(args)
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// We can also use a slice and unpack it to send into the add
// We make a slice and then fill it with values
// When we do ... after it it unpacks the values
func callAdd() {
	// We can call it this way
	M := add(1, 2, 3)
	fmt.Println(M)

	// Or this way
	X := make([]int, 5)
	X[0] = 1
	X[1] = 4
	add(X...)
}

// create functions inside of functions
func closures() {
	// Add function
	add := func(x, y int) int {
		return x + y
	}

	add(1, 2)

	// Be careful of scope
	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
}

// A better way to do it would be to declare a function
// with the variable in scope and then return the function you want.
// Similar to how they do it in Javascript.
// The variable "i" will remain in scope of the function that we call
// so it will stay in memory, but we will operate on the function, which
// is what is returned. So it's a win win.
func closureCreate() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// We call the function closureCreate, which creates and returns the closure.
// Then the variable exists as the function remains in memory, and is not
// finished in execution. So then we can call the returned function, and
// work with the variable we are given.
func closureExample() {
	nextEven := closureCreate()
	fmt.Print(nextEven(), " ")   // 0
	fmt.Print(nextEven(), " ")   // 2
	fmt.Print(nextEven(), " \n") // 4
}

// defer: schedules a function call to be run after the function completes.
// defer is often used when resources need to be freed in some way.
// - deferred functions are run even if a run-time panic occurs
func deferExample() {
	a := func() {
		fmt.Println("a")
	}
	b := func() {
		fmt.Println("b")
	}
	defer b()
	a()
}

// Panic - function to cause a run time error
// Recover
// 	- handle a run-time panic with the built-in recover function
// 	- recover stops the panic and returns the value that was passed to the call to panic
// A panic generally indicates a programmer error
// 	- for example attempting to access an index of an array that's out of bounds, forgetting to initialize a map, etc.
// 	- or an exceptional condition that there's no easy way to recover from.
func panicExample() {
	/*
		Here the call to recover will never happen in this case because the call to panic immediately stops execution of the function
			panic("PANIC")
			str := recover()
			fmt.Println(str, "x")
	*/

	defer func() {
		str := recover()
		fmt.Println("Panic cause:", str)
	}()
	panic("PANIC")
}

func thirdPart() {
	fmt.Println(os.Open("readme.txt"))
	fmt.Println("My favorite number is", rand.Intn(100))
}

func main() {
	// Lets see defer!
	// It runs at the end of everything else!
	// Argument to go/defer must be function call
	a := func() {
		fmt.Println("Lets see when defer works")
	}
	defer a()

	// Run the other stuff
	firstPart()
	closureExample()
	deferExample()
	panicExample()
}
