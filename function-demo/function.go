package functiondemo

import (
	"fmt"
)

// function as variable
var IsUserAdminFn func(string) bool

func IsUserAdminProduction(username string) bool {
	return false
}

func IsUserAdminMock(username string) bool {
	return username == "admin"
}

func IsUserAdmin(username string) bool {
	return IsUserAdminFn(username)
}

func Run() {
	IsUserAdminFn = IsUserAdminProduction
	fmt.Println("IsUserAdminProduction:", IsUserAdmin("admin"))
	IsUserAdminFn = IsUserAdminMock
	fmt.Println("IsUserAdminMock(admin):", IsUserAdmin("admin"))
	fmt.Println("IsUserAdminMock(user):", IsUserAdmin("user"))
}
