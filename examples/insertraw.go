package main

import (
	"fmt"

	sorm "github.com/jiujuan/smileorm"
)

func main() {
	conn := sorm.NewConnection()
	defer conn.Close()

	// ======INSERT raw=====
	insertSQL := "INSERT INTO user(username, email, password_hash) VALUES (?, ?, ?)"
	ins, err := conn.InsertRaw(insertSQL, "tita", "tita@test.com", "tita667788")
	if err != nil {
		fmt.Println("error: ", err)
		// return
	}
	fmt.Println("last insert id: ", ins)
}
