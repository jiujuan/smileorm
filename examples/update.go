package main

import (
	"fmt"

	sorm "github.com/jiujuan/smileorm"
)

func main() {
	type User struct {
		Username      string
		Email         string
		Password_hash string
	}
	conn := sorm.NewConnection()
	defer conn.Close()

	// ======Update=====
	// where := [][]interface{}{{"id", 7}, {"name", 2}}
	// fmt.Println("u where, ", where)
	update := User{
		Username:      "ttest",
		Email:         "ttst@gogo.me",
		Password_hash: "password123123123",
	}
	insID, err := conn.Table("user").Where("id", 7).Update(update)
	fmt.Println(conn.Table("user").DebugSql())

	if err != nil {
		fmt.Println("error: ", err)
		// return
	}
	fmt.Println("last insert id: ", insID)
}
