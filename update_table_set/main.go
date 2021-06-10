package main

import (
	"fmt"
	"strings"
)

type variables struct {
	pwd      string
	fname    string
	lname    string
	location string
	status   string
}

func main() {
	u := variables{
		pwd:      "",
		fname:    "adam",
		lname:    "woolhether",
		location: "ca",
		status:   "",
	}
	table := "person"
	target := "user_name = "
	arg := u.fname

	m := make(map[string]string)

	if u.pwd != "" {
		m["pwd = "] = u.pwd
	}
	if u.fname != "" {
		m["fname = "] = u.fname
	}
	if u.lname != "" {
		m["lname = "] = u.lname
	}
	if u.location != "" {
		m["location = "] = u.location
	}
	if u.status != "" {
		m["status = "] = u.status
	}

	out := buildSQL(table, target, arg, m)
	fmt.Println(out)

}

// To build a UPDATE table SET columns WHERE user =
func buildSQL(table, condition, arg string, columns map[string]string) string {

	var a, b strings.Builder
	b.WriteString("UPDATE " + table + " SET ")

	if len(columns) > 0 {
		for k, v := range columns {
			a.WriteString(k + "'" + v + "'" + ", ")
		}
		b.WriteString(strings.TrimRight(a.String(), ", ") + " WHERE " + condition + "'" + arg + "'")
	}
	return b.String()

}
