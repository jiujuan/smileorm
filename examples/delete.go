package main

import (
	"fmt"

	sorm "github.com/jiujuan/smileorm"
)

func main() {

	conn := sorm.NewConnection()
	defer conn.Close()

	// ======Delete=====
	wher := [][]interface{}{{"id", 12}, {"username", "tom"}}
	insID, err := conn.Table("user").Delete(wher)
	fmt.Println(conn.Table("user").DebugSql())

	if err != nil {
		fmt.Println("error: ", err)
		// return
	}
	fmt.Println("last insert id: ", insID)
}
