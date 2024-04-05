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
		positionQ := strings.Index(sql, "?")
		if positionQ > 0 {
			i++
			buffer.WriteString(sql[:positionQ])
			_, err := fmt.Fprintf(&buffer, "$%d", i)
			sql = sql[positionQ+1:]
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
