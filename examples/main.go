package main

import (
	"fmt"

	sorm "github.com/jiujuan/smileorm"
)

type User struct {
	ID            int
	Username      string
	Email         string
	Password_hash string
}

func main() {
	conn := sorm.NewConnection()
	defer conn.Close()

	//  ======SELECT raw=====
	rs := conn.SelectRaw("SELECT id, username, email, password_hash FROM user WHERE id = ?", 2)

	// fmt.Println(rs.Rows, ";;", rs.Err.Error())

	fmt.Println(rs.Rows.Columns())

	user := User{}
	for rs.Rows.Next() {
		err := rs.Rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password_hash)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user.ID, user.Username, user.Email, user.Password_hash)
	}
}
