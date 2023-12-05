package sql

import (
	"bytes"
	"fmt"
	"strings"
)

func replacePlaceholders(sql string) string {
	buffer := bytes.Buffer{}
	var i int
	for {
		p := strings.Index(sql, "?")
		if p > 0 {
			i++
			buffer.WriteString(sql[:p])
			fmt.Fprintf(&buffer, "$%d", i)
			sql = sql[p+1:]
		} else {
			buffer.WriteString(sql)
			break
		}
	}
	if i == 0 {
		return sql
	}
	return buffer.String()
}
