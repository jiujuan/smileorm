package main

import (
	"fmt"

	sorm "github.com/jiujuan/smileorm"
)

func main() {

	conn := sorm.NewConnection()
	defer conn.Close()

	// ======Delete=====

	insID, err := conn.Table("user").Where("id", 7).Delete()
	fmt.Println(conn.Table("user").DebugSql())

	if err != nil {
		fmt.Println("error: ", err)
		// return
	}
	fmt.Println("last insert id: ", insID)
}
