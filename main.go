// file: main.go
package main

import (
	"Db-DocumentBased/DB"
	"fmt"
)

func main() {
	newDB := DB.NewDB("MyDatabase")

	fields := map[string]string{
		"Age":  "int",
		"Name": "string",
	}
	variables := map[string]interface{}{
		"Age":  12,
		"Name": "Serhat",
	}

	newTable := newDB.NewTable("MyTable", fields)
	retrievedTable := newDB.GetTable("MyTable")
	newDB.InsertInto(newTable, variables)

	fmt.Println("Table Name:", retrievedTable.Variables)
	fmt.Println("Fields:", retrievedTable.Fields)

}
