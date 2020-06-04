package main

import (
	"fmt"

	sorm "github.com/jiujuan/smileorm"
)

func main() {
	conn := sorm.NewConnection()
	defer conn.Close()

	//  ======DELETE raw=====
	deleteSQL := "DELETE FROM user WHERE id = ? AND username = ?"
	affectedRows, err := conn.DeleteRaw(deleteSQL, 5, "TITA")
	if err != nil {
		fmt.Println("error: ", err)

	}
	fmt.Println("delete affected rows: ", affectedRows)
}
