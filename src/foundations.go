package main

import (
	"fmt"
	fig "golang-course/src/figures"
	pk "golang-course/src/mypackage"
	"sync"
)

func exampleFunction(message string) {
	fmt.Println(message)
}

func tripleArgumentsFunctions(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValuesFunction(a, b int) int {
	return a + b
}

func doubleReturnValuesFunction(a, b int) (int, int) {
	return a + b, a - b
}

func testAreaInterface(figure fig.HasArea) {
	fmt.Println(figure.Area())
}

// To be executed by a goroutine
func say(message string, wg *sync.WaitGroup) {

	defer wg.Done() // Will be executed when the function finishes

	fmt.Println(message)
}

func sayWithChannels(message string, channel chan<- string) { // Channel defined as only input, channel <-chan string is only output

	channel <- message // Send message to channel
}

func emailsWithChannels(email string, channel chan string) {
	channel <- email
}

func main() {
	// Constants declaration
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Integer varaibles declarations
	base := 12 // Parsing of the type
	var height int = 14
	var area int

	fmt.Println(base, height, area)

	// Zero values - Default values for variables
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Square area
	const baseSquare = 10
	squareArea := baseSquare * baseSquare
	fmt.Println("Square area:", squareArea)

	//* Arithmetic operators *//
	x := 10
	y := 50

	// Sum
	result := x + y
	fmt.Println("x + y =", result)

	// Subtraction
	result = y - x
	fmt.Println("y - x =", result)

	// Multiplication
	result = x * y
	fmt.Println("x * y =", result)

	// Division
	result = y / x
	fmt.Println("y / x =", result)

	// Module
	result = y % x
	fmt.Println("y % x =", result)

	// Increment
	x++
	fmt.Println("x++ =", x)

	// Decrement
	x--
	fmt.Println("x-- =", x)

	// Relational operators
	fmt.Println("x > y", x > y)

	//* Primitive data types *//
	// Integer numbers
	// int = Depending on OS (32 or 64 bits)
	// int8 = 8bits = -128 to 127
	// int16 = 16bits = -2^15 to 2^15-1
	// int32 = 32bits = -2^31 to 2^31-1
	// int64 = 64bits = -2^63 to 2^63-1

	// Optimize memory when we know that the data will always be positive
	// uint = Depending in OS (32 or 64 bits)
	// uint8 = 8bits = 0 to 127
	// uint16 = 16bits = 0 to 2^15-1
	// uint32 = 32bits = 0 to 2^31-1
	// uint64 = 64bits = 0 to 2^63-1

	// Decimal numbers
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	// Text and boolean
	// string = ""
	// bool = true or false

	// Complex numbers
	// Complex64 = Real and Imaginary float32
	// Complex128 = Real and Imaginary float64
	// Example : c:=10 + 8i

	//* fmt package *//
	var helloMessage string = "Hello"
	var worldMessage string = "World"
	// Println
	fmt.Println(helloMessage, worldMessage)
	// Printf
	name := "John"
	age := 32
	fmt.Printf("%s is %d years old\n", name, age)
	fmt.Printf("%v is %v years old\n", name, age) // %v = for any type
	// Sprintf
	var message = fmt.Sprintf("%s is %d years olds", name, age)
	fmt.Println(message)
	// Know the type of a variable
	fmt.Printf("helloMessage: %T\n", message)

	//* Functions *//
	exampleFunction("Inside exampleFunction")
	tripleArgumentsFunctions(1, 2, "c")
	fmt.Println(returnValuesFunction(1, 2))
	value1, value2 := doubleReturnValuesFunction(1, 2)
	fmt.Println("value1: ", value1, "value2: ", value2)
	_, value3 := doubleReturnValuesFunction(1, 2) // _ = ignore the value
	fmt.Println("value3: ", value3)

	//* Loops *//
	// Conditional for
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	fmt.Printf("\n")

	// For While
	counter := 0
	for counter < 10 {
		fmt.Println("Counter:", counter)
		counter++
	}

	// For infinite
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	//* Logical operators *//
	bool1 := true
	bool2 := false
	num1 := 10
	num2 := 20
	// AND
	fmt.Println("bool1 && bool2:", bool1 && bool2)

	// OR
	fmt.Println("bool1 || bool2:", bool1 || bool2)

	// NOT
	fmt.Println("!bool1:", !bool1)

	// Equal
	fmt.Println("num1 == num2:", num1 == num2)

	// Different
	fmt.Println("num1 != num2:", num1 != num2)

	// Greater than
	fmt.Println("num1 > num2:", num1 > num2)

	// Less than
	fmt.Println("num1 < num2:", num1 < num2)

	// Greater than or equal
	fmt.Println("num1 >= num2:", num1 >= num2)

	// Less than or equal
	fmt.Println("num1 <= num2:", num1 <= num2)

	//* If *//
	// If
	if num1 == 10 {
		fmt.Println("num1 is 10")
	}

	// If else if

	if num1 == 10 {
		fmt.Println("num1 is 10")
	} else if num1 == 20 {
		fmt.Println("num1 is 20")
	} else {
		fmt.Println("num1 is not 10 or 20")
	}

	// Switch
	switch num1 {
	case 10:
		fmt.Println("num1 is 10")
	case 20:
		fmt.Println("num1 is 20")
	default:
		fmt.Println("num1 is not 10 or 20")
	}

	// Defer
	defer fmt.Println("This is a defer") // This line will be executed at the end of the function

	// Break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("i:", i)
	}

	// Continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("i:", i)
	}

	//* Arrays *// - Immutable
	var array [5]int // Array of 5 zeros as zero values
	fmt.Println("array:", array)

	array[2] = 100
	fmt.Println("array:", array, "Length:", len(array), "Max capacity", cap(array))

	//* Slices *// - Mutable
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice:", slice, "Length:", len(slice), "Max capacity", cap(slice))

	// Slicing
	fmt.Println("slice[1]:", slice[1])
	fmt.Println("slice[:3]:", slice[:3])
	fmt.Println("slice[2:4]:", slice[2:4])
	fmt.Println("slice[3:]:", slice[3:])

	// Append
	slice = append(slice, 6)
	fmt.Println("slice:", slice, "Length:", len(slice), "Max capacity", cap(slice))

	// Append new slice
	newSlice := []int{7, 8, 9}
	slice = append(slice, newSlice...)
	fmt.Println("slice:", slice, "Length:", len(slice), "Max capacity", cap(slice))

	//* Range to iterate *//
	for index, value := range slice {
		fmt.Println("Index:", index, "Value:", value)
	}
	fmt.Printf("\n")
	for index, value := range "Hello World" {
		fmt.Println("Index:", index, "Value:", string(value))
	}

	//* Maps *//
	dictionary := make(map[string]int)
	dictionary["one"] = 1
	dictionary["two"] = 2
	fmt.Println("dictionary:", dictionary)

	// Other to instantiate a map
	dictionary2 := map[string]int{"one": 1, "two": 2}
	fmt.Println("dictionary2:", dictionary2)

	// Iterate over a map
	for key, value := range dictionary { // Will not iterate in the same order
		fmt.Println("Key:", key, "Value:", value)
	}

	// Find value
	value, ok := dictionary["number not in"] // The second value is a boolean that describes if the key is present
	fmt.Println(value, ok)                   // The value is equal to the zero value

	//* Structs *//
	myCar := pk.Car{Brand: "Ford", Model: 2019}
	fmt.Println(myCar)

	// Other way to instantiate a struct
	var otherCar pk.Car
	otherCar.Brand = "Ferrari"
	fmt.Println(otherCar) // model attr as zero value

	//* Pointers *//
	variable := 50
	pointerToVariable := &variable

	fmt.Println("variable:", variable, "pointerToVariable:", pointerToVariable)
	fmt.Println("Value inside th memory address:", *pointerToVariable)

	*pointerToVariable = 100
	fmt.Println("variable:", variable, "pointerToVariable:", pointerToVariable)

	var myPc pk.Pc = pk.Pc{Ram: 16, Disk: 500, Brand: "Lenovo"}
	fmt.Println(myPc)
	myPc.Ping()
	myPc.DuplicateRAM()
	fmt.Println(myPc)
	myPc.DuplicateRAM()
	fmt.Println(myPc)

	//* Interfaces *//
	square := fig.Square{Side: 12}
	rectangle := fig.Rectangle{Base: 6, Height: 13}
	testAreaInterface(square)
	testAreaInterface(rectangle)

	//* Concurrency - Goroutines *// - Go routines by themself cant trade data between them
	var wg sync.WaitGroup // Allows to interect with set of gorutines and wait for them to finish
	fmt.Println("Hello")  //say("Hello")

	wg.Add(1)
	go say("World", &wg) // This will be executed in a different thread

	// Usually gorutines are used with anonymous functions
	go func(text string) {
		fmt.Println(text)
	}("Hello from anonymous function")

	//time.Sleep(1 * time.Second) // This is to wait for the goroutine to execute - Not efficient

	wg.Wait() // Wait for the goroutines to finish

	//* Concurrency - Channels *// - Channels are used to trade data between goroutines
	channel := make(chan string, 1) // The second parameter defines how many values wil be stored in the channel, if not defined it will be dynamic

	// The channels are less efficient than goroutines, but are easier to use
	go sayWithChannels("Hello from channel", channel)

	fmt.Println(<-channel) // Read from channel

	// Close and range
	channel2 := make(chan string, 2)

	channel2 <- "Message1"
	channel2 <- "Message2"

	fmt.Println(len(channel2), cap(channel2))

	close(channel2) // Close the channel to receive more data

	for message := range channel2 { // To iterate over a channel is recommened to close it first
		fmt.Println(message)
	}

	// Select - Used to select which channel will be read
	email1 := make(chan string)
	email2 := make(chan string)

	go emailsWithChannels("email1", email1)
	go emailsWithChannels("email2", email2)
	// At this point we dont know which channel will be read first, we use select to know it

	for i := 0; i < 2; i++ {
		select {
		case message1 := <-email1:
			fmt.Println("Message from email1:", message1)
		case message2 := <-email2:
			fmt.Println("Message from email2:", message2)
		}
	}

	close(email1)
	close(email2)

	//* Go modules management *//
	// To download a module use the command: go get -v -u <module url>
}
