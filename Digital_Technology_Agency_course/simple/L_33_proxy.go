package main

import (
	"fmt"
	"simple/proxy"
)

var (
	admin = "administrator"
	user  = "user"
	users = map[string]bool{
		admin: true,
		user:  false,
	}
)

func main() {
	proxy := proxy.ProxyDataBase{
		Users: users,
		Db:    &proxy.DataBase{},
	}
	adminData, err := proxy.GetData(admin)
	fmt.Printf("From [%s] Data:[%v] Err:[%v]\n", admin, adminData, err)
	userData, err := proxy.GetData(user)
	fmt.Printf("From [%s] Data:[%v] Err:[%v]\n", user, userData, err)

}
