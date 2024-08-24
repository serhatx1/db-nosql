package DB

import (
	"fmt"
	"math/rand"
	"time"
)

type DB struct {
	Name   string
	Tables map[string]*Table
}

type Table struct {
	Name      string
	Fields    map[string]string
	Variables map[string]interface{}
}

func NewDB(name string) *DB {
	return &DB{Name: name, Tables: make(map[string]*Table)}
}

func (database *DB) NewTable(name string, fieldMap map[string]string) *Table {
	table := &Table{
		Name:   name,
		Fields: fieldMap,
	}
	database.Tables[name] = table
	return table
}

func (database *DB) GetTable(name string) *Table {
	return database.Tables[name]
}
func (database *DB) InsertInto(table *Table, mapVariables map[string]interface{}) error {
	if table.Variables == nil {
		table.Variables = make(map[string]interface{})
	}

	for k, v := range mapVariables {
		expectedType, ok := table.Fields[k]

		fmt.Println(expectedType)
		if !ok {
			return fmt.Errorf("unexpected key %s", k)
		}

		actualType := fmt.Sprintf("%T", v)

		if actualType == expectedType {
			table.Variables[k] = v
		} else {
			return fmt.Errorf("format error for key %s: expected type %s but got %s", k, expectedType, actualType)
		}
	}

	table.Variables["objectid"] = GenerateID()
	return nil
}

func GenerateID() string {
	timestamp := time.Now().UnixNano()
	random := rand.Int63()
	return fmt.Sprintf("%x-%x", timestamp, random)
}
