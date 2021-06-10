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

	m := make(map[string]string)

	if u.pwd != "" {
		m["pwd = "]=u.pwd
	}
	if u.fname != "" {
		m["fname = "] = u.fname
	}
	if u.lname != "" {
		m["lname = "]=u.lname
	}
	if u.location != "" {
		m["location = "]=u.location
	}
	if u.status != "" {
		m["status = "]=u.status
	}

	out := buildSQL(table, m)
	fmt.Println(out)

}

// To build a SELECT * FROM <table-name> with optional WHERE
func buildSQL(table string, queries map[string]string) string {
	q := "'"

	var a, b strings.Builder
	a.WriteString("SELECT * FROM " + table)

	if len(queries) > 0 {
		a.WriteString(" WHERE ")
		for k, v := range queries {
			a.WriteString(k+q+v+q+" AND ")
		}
		b.WriteString(strings.TrimRight(a.String(), " AND "))
		return b.String()
	}
	return a.String()
}