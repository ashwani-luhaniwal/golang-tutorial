// structure is a user defined type which represents a collection of fields
// It can be used in places where it makes sense to group the data into a single unit rather maintaining 
// each of them as separate types.

/***********************************
	Declaring a structure
***********************************/
/*
	// named structure - used to create structures of type Employee
	type Employee struct {
		firstName	string
		lastName	string
		age			int
	}

	or 

	type Employee struct {
		firstName, lastName	string
		age, salary			int
	}
*/

// Declaring anonymous structures without declaring new type
/*
	var employee struct {
		firstName, lastName	string
		age	int
	}
*/

package main

import (
	"fmt"
	"golang-tutorial/structures/structs/computer"
)

/**********************************
	Creating named structures
**********************************/
type Employee struct {
	firstName, lastName string
	age, salary			int
}

func main() {

	// creating structure using field names
	emp1 := Employee {
		firstName: "Ashwani",
		age:		25,
		salary:		500,
		lastName:	"Luhaniwal",
	}

	// creating structure without using field names
	emp2 := Employee{"Ashwani", "Kumar", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	/************************************
		Creating anonymous structures
	************************************/
	// creates new struct variable and does not define any new struct type
	emp3 := struct {
		firstName, lastName string
		age, salary			int
	} {
		firstName: "Ashwani",
		lastName: "Luhaniwal",
		age: 31,
		salary: 5000,
	}
	fmt.Println("Employee 3", emp3)

	/************************************
		Zero value of structure
	************************************/
	// When a struct is defined and it is not explicitly initialised with any value,
	// the fields of struct are assigned their zero values by default
	var emp4 Employee	// zero valued structure
	fmt.Println("Employee 4", emp4)

	// It is also possible to specify values for some fields and ignore the rest
	// The ignored field names are assigned zero values
	emp5 := Employee {
		firstName: "Ashwani",
		lastName: "Luhaniwal",
	}
	fmt.Println("Employee 5", emp5)

	/*********************************************
		Accessing individual fields of a struct
	*********************************************/
	// the dot . operator is used to access the individual fields of a structure
	emp6 := Employee{"Ashwani", "Luhaniwal", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: %d", emp6.salary)

	// It's also possible to create zero struct and then assign values to its fields later
	var emp7 Employee
	emp7.firstName = "Ashwani"
	emp7.lastName = "Luhaniwal"
	fmt.Println("Employee 7:", emp7)

	/*******************************
		Pointers to a struct
	*******************************/
	// It's also possible to create pointers to struct
	emp8 := &Employee{"Ashwani", "Luhaniwal", 55, 6000}
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)

	// Go lang allows us to use emp8.firstName instead of explicit dereference (*emp8).firstName to access firstName field
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)

	/*******************************
		Anonymous fields
	******************************/
	// It's possible to create structs with fields which contain only a type without the field name
	// These kind of fields are called anonymous fields.
	/*
		type Person struct {
			string
			int
		}
	*/

	type Person struct {
		string
		int
	}

	p := Person{"Ashwani", 50}
	fmt.Println(p)

	// Even though an anonymous fields doesn't have a name, by default name of anonymous field is name of its type
	// Person struct has 2 fields with name string and int
	var p1 Person
	p1.string = "Ashwani"
	p1.int = 50
	fmt.Println(p1)

	/***************************
		Nested structs
	**************************/
	// It's possible a struct contains a field which is another struct
	type Address struct {
		city, state string
	}
	type PersonDetails struct {
		name string
		age int
		address Address
	}

	var p2 PersonDetails
	p2.name = "Ashwani"
	p2.age = 50
	p2.address = Address {
		city: "Bangalore",
		state: "Karnataka",
	}
	fmt.Println("Name:", p2.name)
	fmt.Println("Age:", p2.age)
	fmt.Println("City:", p2.address.city)
	fmt.Println("State:", p2.address.state)

	/******************************
		Promoted fields
	*****************************/
	// Fields that belong to anonymous struct field in structure are called promoted fields
	// since they can be accessed as if they belong to structure which holds the anonymous struct field
	/*
		type Address struct {
			city, state string
		}
		type Person struct {
			name string
			age int
			Address	// anonymous struct field
		}
		// The fields of Address struct namely "city" and "state" are called promoted field since they can be
		// accessed as if they are directly declared in Person struct itself
	*/
	type AddressInfo struct {
		city, state string
	}
	type PersonInfo struct {
		name string
		age int
		AddressInfo
	}

	var p3 PersonInfo
	p3.name = "Ashwani"
	p3.age = 50
	p3.AddressInfo = AddressInfo {
		city: "Bangalore",
		state: "Karnataka",
	}
	fmt.Println("Name:", p3.name)
	fmt.Println("Age:", p3.age)
	fmt.Println("City:", p3.city)	// city is promoted field
	fmt.Println("State:", p3.state)	// state is promoted field

	/**************************************
		Exported structs and fields
	**************************************/
	// If a struct type starts with a capital letter, then it is a exported type and it can be 
	// accessed from other packages. Similarly if the fields of a structure start with caps, 
	// they can be accessed from other packages
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec:", spec)

	/*****************************
		Structs Equality
	*****************************/
	// Structs are value types and are comparable if each of their fields are comparable.
	// Two struct variable considered equal if their corresponding fields are equal
	type name struct {
		firstName string
		lastName string
	}

	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

	// Struct variables are not comparable if they contain fields which are not comparable
	/*
		type image struct {
			data map[int]int
		}

		// image struct type contains a field data which is of type map
		// maps are not comparable, hence image1 and image2 cannot be compared
		image1 := image{data: map[int]int {
			0: 155,
		}}
		image2 := image{data: map[int]int {
			0: 155,
		}}
		if image1 == image2 {	// throws error
			fmt.Println("image1 and image2 are equal")
		}
	*/

}
