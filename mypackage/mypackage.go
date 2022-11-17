package mypackage

import "fmt"

type Car struct {
	Brand      string // First letter in mayus to set the attr as public
	Model      int
	kilometers int // First letter in minus to set the attr as private
}

type car struct { // Private struct
	brand string
}

func privateFunction() {
	// This function is private
}

func PublicFunction() {
	// This function is public
}

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPc Pc) Ping() { // Pc is a copy of the object instance
	fmt.Println(myPc.Brand, "Pong")
}

func (myPc *Pc) DuplicateRAM() { // *pc is a pointer to pc so we can modify the value of the object instance
	myPc.Ram = myPc.Ram * 2
}

func (myPc Pc) String() string { // Function used by fmt.Println to print the object (stringers)
	return fmt.Sprintf("A %s PC with %d GB of RAM and %d GB of disk", myPc.Brand, myPc.Ram, myPc.Disk)
}
