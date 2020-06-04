package main

import (
	"fmt"

	sorm "github.com/jiujuan/smileorm"
)

func main() {
	conn := sorm.NewConnection()
	defer conn.Close()

	//  ======UPDATE raw=====
	updateSQL := "UPDATE user SET username = ?, email = ? WHERE id = ?"
	affectedRows, err := conn.UpdateRaw(updateSQL, "jimmy", "jimmy007@test.com", 3)
	if err != nil {
		fmt.Println("error: ", err)
		// return
	}
	fmt.Println("update affected rows: ", affectedRows)
}
