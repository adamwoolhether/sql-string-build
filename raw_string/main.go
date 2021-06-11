package main

import (
	"fmt"
	"strings"
)

func main() {

	table := "node_role"
	columns := []string{"id", "name", "scope", "description"}
	args := []string{"1234", "NewName", "NewScope", "this is desc"}

	out := getBulkInsertSQL(table, columns, args)
	fmt.Println(out)
}

func getBulkInsertSQL(table string, columns, args []string) string {
	var b strings.Builder

	b.WriteString("INSERT INTO " + table + "(" + strings.Join(columns, ", ") + ") VALUES ")

	b.WriteString("(")
	for i, v := range args {
		b.WriteString("'" + v + "'")
		if i != len(args)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString(")")

	return b.String()
}
