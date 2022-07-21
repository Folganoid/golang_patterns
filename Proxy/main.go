package main

import (
	"./pkg"
	"fmt"
)

var (

	admin = "administrator"
	user = "guest"

	users = map[string]bool{
		admin: true,
		user: false,
	}
)

func main() {

	var service pkg.Service

	service = pkg.ProxyDataBase{
		Users: users,
		Db: &pkg.DataBase{},
	}

	adminData, err := service.GetData(admin)
	fmt.Printf("From [%s] Data: [%v] Err: [%v]\n", admin, adminData, err)

	userData, err := service.GetData(user)
	fmt.Printf("From [%s] Data: [%v] Err: [%v]\n", user, userData, err)

}