// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"fmt"

	"github.com/lovexiaoe/golangpros/types/exportUnexport/exportUnexport2/entities"
)

// main is the entry point for the application.
func main() {
	// Create a value of type Admin from the entities package.
	a := entities.Admin{
		Rights: 10,
	}

	// even though the inner type "user" is unexported,the fields declared within the inner type are exported.
	// since the identifiers from inner type are promoted to the outer type。
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
