package main

import (
	"fmt"
	"strings"
)

func main() {
	table := "node_role"
	//condition := "id = "

	out := buildSQLDelete(table, "", "test")
	fmt.Println(out)

	out2 := buildSQLDeleteQuota("disk_quota","name = ", "quotaname")
	fmt.Println(out2)
}

func buildSQLDelete(table, condition, arg string) string {
	var b strings.Builder
	if len(condition) == 0 || len(arg) == 0 {
		return ""
	}
	b.WriteString("DELETE FROM " + table + " WHERE " + condition + "'" + arg + "'")
	return b.String()
}

func buildSQLDeleteQuota(table, condition, arg string) string {
	var b strings.Builder
	if len(condition) == 0 || len(arg) == 0 {
		return ""
	}
	b.WriteString("DELETE FROM " + table + " WHERE " + condition + "'" + arg + "'" + " AND enabled IS NOT TRUE")
	return b.String()
}