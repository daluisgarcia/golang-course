package main

import "testing"

//* Test Mocking *//
func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "12345678",
			mockFunc: func() {
				// This code will substitute the real code of the original functions
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						id: 1,
					}, nil
				}

				GetPersonByDni = func(dni string) (Person, error) {
					return Person{
						dni: "12345678",
						name: "John",
						age: 30,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					dni:  "12345678",
					name: "John",
					age:  30,
				},
				Employee: Employee{
					id: 1,
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDni

	for _, testCase := range table {
		testCase.mockFunc() // Calling this will make the override of the original functions

		employee, err := GetFullTimeEmployeeById(testCase.id, testCase.dni)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if employee != testCase.expectedEmployee {
			t.Errorf("Expected: %v, but got: %v", testCase.expectedEmployee, employee)
		}

		// Restore the original functions
		GetEmployeeById = originalGetEmployeeById
		GetPersonByDni = originalGetPersonByDni
	}
}
