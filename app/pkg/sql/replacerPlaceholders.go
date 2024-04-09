package sql

import (
	"bytes"
	"fmt"
	"strings"

	"server/app/pkg/errors"
	"server/app/pkg/logging"
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
			if err != nil {
				logging.GetLogger().Error(errors.InternalServer.Wrap(err))
			}
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
