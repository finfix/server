package ddlHelper

import "strconv"

func BuildJoin(joinTableWithAlias, joinColumnWithPrefix, columnWithPrefix string) string {
	return joinTableWithAlias + " ON " + joinColumnWithPrefix + " = " + columnWithPrefix
}

func WithCustomAlias(table, alias string) string {
	return table + " " + alias
}

func WithCustomPrefix(column, prefix string) string {
	return prefix + "." + column

}

func As(column, newName string) string {
	return column + " AS " + newName
}

func Coalesce(column, defaultValue string) string {
	return "COALESCE(" + column + ", " + defaultValue + ")"
}

func Max(column string) string {
	return "MAX(" + column + ")"
}

func Sum(column string) string {
	return "SUM(" + column + ")"
}

func Plus(column string, value int) string {
	return column + " + " + strconv.Itoa(value)
}

func Minus(column string, value int) string {
	return column + " - " + strconv.Itoa(value)
}

func Desc(column string) string {
	return column + " DESC"
}

func Asc(column string) string {
	return column + " ASC"
}

const SelectAll = "*"
