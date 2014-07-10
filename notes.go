/* My Go Notes
Sources
	- http://www.golang-book.com
	- http://talks.golang.org/2012/10things.slide
*/

// Every Go program must start with a package declaration. Packages are Go's way of organizing and reusing code.
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
	"strings"
	"bytes"
	"io/ioutil"
	"errors"
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

// Pointers reference a location in memory where a value is stored rather than the value itself
// By using a pointer (*int) the zero function is able to modify the original variable
func makeItZero(x *int) {
	// Uncover the value of the pointer
	// We've to unpack the x because it's a pointer right now
	// then change the value
	// We have to ->>>> DEREFERENCE THE POINTER
	*x = 0
}

// One way to do this is to use a special data type known as a pointer
func pointerExample() {
	x := 4

	// & operator to find the address of a variable
	fmt.Println("Pointer address:", &x)

	fmt.Println("Value:", x)
	makeItZero(&x)
	fmt.Println("Value:", x)
}

func one(x *int) {
	*x = 1
}

// new takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it
func newExample() {
	x := new(int)
	fmt.Println("We created new x", x)
	one(x)
	fmt.Println("Value of x is", *x)
}

// A struct is a type which contains named fields.
type Circle struct {
	x, y, r float64
}

// Anonymous structs do exist
var config struct {
	key  string
	name string
}

// Struct methods
// This function will be then accessible by
// structInstance.area()
// Special type of function known as a method
// In between the keyword func and the name of the function we've added a “receiver”.
// The receiver is like a parameter – it has a name and a type – but by creating the function in this way it allows us to call the function using the . operator.
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// You have to pass the struct in as arguments here
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func structExample() {
	// Ways to declare:
	// 	- var c Circle
	// 	- c := new(Circle)
	// 	- c := Circle{0, 0, 5}
	c := Circle{x: 1, y: 1, r: 5}
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(circleArea(c))
	fmt.Println(c.area())
}

// Defining a person struct
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

/*
	- Now we've defined an Android
	- This would work, but we would rather say an Android is a Person, rather than an Android has a Person. Go supports relationships like this by using an embedded type.
	type Android struct {
	    Person Person
	    Model string
	}
*/

// Known as anonymous fields, embedded types look like this
// Normally you'd do:
// a := new(Android)
// a.Person.Talk()
// This is kind of like inheritance, but it automatically
// inherits all the information without a keyword.
// It is taking multiple relationships and making them accessible
// to the Android struct.
// Here an Android is a Person
// You can do a.Person.Talk() or a.Talk()
//
// The is-a relationship works this way intuitively: People can talk, an android is a person, therefore an android can talk
type Android struct {
	Person
	Model string
}

// Interfaces
// Interface is created using the type keyword, followed by a name and the keyword interface
// But instead of defining fields, we define a “method set”
// Method set is a list of methods that a type must have in order to “implement” the interface

type Shape interface {
	area() float64
}

// Concurrency
// Making progress on more than one task simultaneously is known as concurrency
// Go has rich support for concurrency using goroutines and channels
// A goroutine is a function that is capable of running concurrently with other functions.
//
// To create a goroutine we use the keyword go followed by a function invocation

func fx(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, "::", i)
	}
}

func fx2(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// This program consists of two goroutines
// 	- The first goroutine is implicit and is the goRoutineExample function
// 	- The second goroutine is created when we call go f(0)
// Normally when we invoke a function our program will execute all the statements in a function and then return to the next line following the invocation.
// * With a goroutine we return immediately to the next line and don't wait for the function to complete
// WITHOUT the Scanln the program would just exit without finishing
func goRoutineExample() {
	go fx(0)
	// for i := 0; i < 10; i++ {
	// 	go fx2(i)
	// }
	// var input string
	// fmt.Scanln(&input)
}

// Channels
// Channels provide a way for two goroutines to communicate with one another and synchronize their execution
// Normally channels are synchronous; both sides of the channel will wait until the other side is ready.
// A channel type is represented with the keyword chan followed by the type of the things that are passed on the channel (in this case we are passing strings)
// The <- (left arrow) operator is used to send and receive messages on the channel.
// c <- "ping" means send
// "ping". msg := <- c means receive a message and store it in msg.

func pinger(c chan string) {
	for i := 0; ; i++ {
		// send
		c <- "ping"
	}
}

// Using a channel like this synchronizes the two goroutines
// When pinger attempts to send a message on the channel it will wait until printer is ready to receive the message.
// 	- known as blocking

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		// means receive a message and store it in msg
		msg := <-c
		// Or fmt.Println(<-c)
		fmt.Println(msg)
		time.Sleep(time.Second * 1) // or would go out of control
	}
}

func channelExample() {
	var c chan string = make(chan string)
	go pinger(c)
	// The program will now take turns printing “ping” and “pong”.
	/* because of this -> */ go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

// Channel Direction
// We can specify a direction on a channel type thus restricting it to either sending or receiving.
// For example pinger's function signature can be changed to this:
// 	- func pinger(c chan<- string)
// 		- Now c can only be sent to.
// 	- func printer(c <-chan string)
// 		- Now c can only send
// 	- A channel that doesn't have these restrictions is known as bi-directional
// 		- A bi-directional channel can be passed to a function that takes send-only or receive-only channels, but the reverse is not true.

// Select statement (like switch) for Channels
func selectExample() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

// It's also possible to pass a second parameter to the make function when creating a channel
// This creates a buffered channel with a capacity of 1.
// A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full.
func bufferedChannelExample() {
	c := make(chan int, 1)
	c <- 1
}

// Go includes a large number of functions to work with strings in the strings package
func understandString() {
	strings.Contains("test", "es")
	strings.Count("test", "t")
	strings.HasPrefix("test", "te")
	strings.HasSuffix("test", "st")
	strings.Index("test", "e")
	strings.Join([]string{"a", "b"}, "-")
	strings.Repeat("a", 5)
	strings.Replace("aaaa", "a", "b", 2)
	strings.Split("a-b-c-d-e", "-")
	strings.ToLower("TEST")
	strings.ToUpper("test")
}

// Input Output
// The io package consists of a few functions, but mostly interfaces used in other packages. The two main interfaces are Reader and Writer.
// Readers support reading via the Read method.
// Writers support writing via the Write method.
// Many functions in Go take Readers or Writers as arguments.
// For example the io package has a Copy function which copies data from a Reader to a Writer

// To convert a string to a slice of bytes (and vice-versa)
func stringandBytes(){
	arr := []byte("test")
	fmt.Println(arr)
	str := string([]byte{'t','e','s','t'})
	fmt.Println(str)
}

// To read or write to a []byte or a string you can use the Buffer struct found in the bytes package
// A Buffer doesn't have to be initialized and supports both the Reader and Writer interfaces.
// You can convert it into a []byte by calling buf.Bytes()
// If you only need to read from a string you can also use the strings.
// 	- NewReader function which is more efficient than using a buffer.
func understandBuffer(){
	var buf bytes.Buffer
	buf.Write([]byte("test"))
	fmt.Println("Buffer to Bytes:", buf.Bytes())

	// More easily
	fmt.Println("Buffer to String:", string(buf.Bytes()))
	fmt.Println("Buffer to String()", buf.String())
}

// http://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
func stringConcatenate() {
	var buffer bytes.Buffer
    for i := 0; i < 1000; i++ {
        buffer.WriteString("a")
    }
    fmt.Println(buffer.String())
}

// Files & Folders
// We use defer file.Close() right after opening the file to make sure the file is closed as soon as the function completes.
func filesandFolders() {
	// Load file
	file, err := os.Open("resources/readme.txt")
	defer file.Close()
	if err != nil {
        return
    }

    // File size
    stat, err := file.Stat()
    if err != nil {
        return
    }

    // Read file
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        return
    }

    // Print string
    str := string(bs)
    fmt.Println("Contents", "\"" + str + "\"")
}

func simplerReadFile() {
	bs, err := ioutil.ReadFile("resources/readme.txt")
    if err != nil {
        return
    }
    str := string(bs)
    fmt.Println(str)
}

func simplercreateFile() {
	file, err := os.Create("resources/writtenREADME.txt")
    if err != nil {
        return
    }
    defer file.Close()
    file.WriteString("test")
}

func errorType() {
	err := errors.New("error message")
	if err != nil {
		fmt.Println(err)
	}
}

func randomExample() {
	fmt.Println("My favorite number is", rand.Intn(100))
}

func main() {
	// Lets see defer!
	// It runs at the end of everything else!
	// Argument to go/defer must be function call
	/*
		a := func() {
			fmt.Println("Lets see when defer works. It works at the end!")
		}
		defer a()
	*/

	//Run the other stuff
	//firstPart()
	//goRoutineExample()
	//closureExample()
	//deferExample()
	//panicExample()
	//pointerExample()
	//newExample()
	//structExample()
	//channelExample()
	//selectExample()
	//understandString()
	//stringandBytes()
	//understandBuffer()
	//filesandFolders()
	//simplerReadFile()
	//simplercreateFile()
	errorType()
}
