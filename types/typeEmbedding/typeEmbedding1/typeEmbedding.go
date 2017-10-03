// 这里的类型嵌套和匿名字段相同，主要说明类型嵌套的一些用法

package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method that can be called via
// a value of type user.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin represents an admin user with privileges.
type admin struct {
	user  // Embedded Type
	level string
}

// main is the entry point for the application.
func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly.
	ad.user.notify()

	// The inner type's method is promoted.
	ad.notify()
}
