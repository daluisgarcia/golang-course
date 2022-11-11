package oop

//* INTERFACE
type PrintInfo interface {
	getMessage() string
}

//* TYPE PERSON
type Person struct {
	name string
	age  int
}

// Receiver functions
func (e *Person) SetName(name string) {
	e.name = name
}

func (e *Person) SetAge(age int) {
	e.age = age
}

func (e *Person) GetName() string {
	return e.name
}

func (e *Person) GetAge() int {
	return e.age
}

//* TYPE EMPLOYEE
type Employee struct {
	id int
}

// Receiver functions
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) GetId() int {
	return e.id
}

//* TYPE FULLTIME EMPLOYEE
type FullTimeEmployee struct {
	Person // Anonymous field to inherit from Person
	Employee
}

// Constructor
func NewFullTimeEmployee(id int, name string, age int) *FullTimeEmployee {
	return &FullTimeEmployee{
		Person: Person{
			name: name,
			age:  age,
		},
		Employee: Employee{
			id: id,
		},
	}
}

// Implementing the interface method - Implicit implementation
func (e *FullTimeEmployee) getMessage() string {
	return "Full time employee"
}

// TYPE TEMPORARY EMPLOYEE
type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (e *TemporaryEmployee) SetTaxRate(taxRate int) {
	e.taxRate = taxRate
}

func (e *TemporaryEmployee) GetTaxRate() int {
	return e.taxRate
}

// Implementing the interface method - Implicit implementation
func (e *TemporaryEmployee) getMessage() string {
	return "Temporary employee"
}

func GetMessage(p PrintInfo) string {
	return p.getMessage()
}
