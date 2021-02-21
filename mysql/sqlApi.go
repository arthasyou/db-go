package mysql

import (
	"fmt"
	"strings"
)

// =================================================================================
// Api struct
// =================================================================================

// KeyOpVal for where statement used
type KeyOpVal struct {
	Key   string
	Op    string
	Value string
}

// KeyVal for insert replace statement used
type KeyVal struct {
	Key   string
	Value string
}

// =================================================================================
// Api function
// =================================================================================

// Select query
func Select(table string, args ...[]string) string {
	cls := ""
	if len(args) > 0 {
		cls = formatItems(args[0])
	} else {
		cls = "*"
	}
	query := "SELECT " + cls + " FROM " + formatItem(table)
	return query
}

// Insert into database
func Insert(table string, kvs []*KeyVal) string {
	return "INSERT INTO " + formatItem(table) + formatKV(kvs)
}

// Replace into database
func Replace(table string, kvs []*KeyVal) string {
	return "REPLACE INTO " + formatItem(table) + formatKV(kvs)
}

// Update into database
func Update(table string, kvs []*KeyVal) string {
	return "UPDATE " + formatItem(table) + " SET " + formatKeV(kvs)
}

// Deltet from database
func Deltet(table string) string {
	return "DELETE FROM " + formatItem(table)
}

// Where statement
func Where(kovs []*KeyOpVal) string {
	ss := formatKOV(kovs)
	s := " WHERE " + strings.Join(ss, " AND ")
	return s
}

// WhereOr statement
func WhereOr(kovs []*KeyOpVal) string {
	ss := formatKOV(kovs)
	s := " WHERE " + strings.Join(ss, " OR ")
	return s
}

// SubAnd join to where statement
func SubAnd(kovs []*KeyOpVal) string {
	ss := formatKOV(kovs)
	s := " AND " + strings.Join(ss, " AND ")
	return s
}

// SubOr join to where statement
func SubOr(kovs []*KeyOpVal) string {
	ss := formatKOV(kovs)
	s := " OR " + strings.Join(ss, " OR ")
	return s
}

// Order statement
func Order(key string) string {
	return " ORDER BY " + formatItem(key)
}

// OrderDesc statement
func OrderDesc(key string) string {
	return " ORDER BY " + formatItem(key) + " DESC"
}

// Limit statement
func Limit(page uint, column uint) string {
	return fmt.Sprintf(" LIMIT %v, %v", (page-1)*column, column)
}

// =================================================================================
// Internal function
// =================================================================================

func formatItem(c string) string {
	return "`" + c + "`"
}

func formatItems(columns []string) string {
	cls := make([]string, len(columns))
	for i, c := range columns {
		cls[i] = formatItem(c)
	}
	return strings.Join(cls, ",")
}

func formatValue(v string) string {
	return "'" + v + "'"
}

func formatValues(vs []string) string {
	cls := make([]string, len(vs))
	for i, v := range vs {
		cls[i] = formatValue(v)
	}
	return strings.Join(cls, ",")
}

func formatKOV(kovs []*KeyOpVal) []string {
	var ss = make([]string, len(kovs))
	for i, kov := range kovs {
		ss[i] = formatItem(kov.Key) + " " + kov.Op + " " + formatValue(kov.Value)
	}
	return ss
}

func formatKV(kvs []*KeyVal) string {
	length := len(kvs)
	var keys = make([]string, length)
	var values = make([]string, length)
	for i, kv := range kvs {
		keys[i] = formatItem(kv.Key)
		values[i] = formatValue(kv.Value)
	}
	keyStr := strings.Join(keys, ",")
	valueStr := strings.Join(values, ",")
	r := " (" + keyStr + ") VALUES (" + valueStr + ")"
	return r
}

func formatKeV(kvs []*KeyVal) string {
	var ss = make([]string, len(kvs))
	for i, kv := range kvs {
		ss[i] = formatItem(kv.Key) + " = " + formatValue(kv.Value)
	}
	return strings.Join(ss, ",")
}
