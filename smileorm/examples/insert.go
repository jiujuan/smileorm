package main

import (
	"fmt"

	sorm "github.com/jiujuan/smileorm"
)

func main() {
	type User struct {
		ID            int
		Username      string
		Email         string
		Password_hash string
	}
	conn := sorm.NewConnection()
	defer conn.Close()

	// ======INSERT=====
	user := User{
		ID:            7,
		Username:      "jimmy",
		Email:         "jimmy@gogo.me",
		Password_hash: "password123123123",
	}
	insID, err := conn.Table("user").Insert(user)
	fmt.Println(conn.Table("user").DebugSql())

	if err != nil {
		fmt.Println("error: ", err)
		// return
	}
	fmt.Println("last insert id: ", insID)
}
