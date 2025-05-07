package utils

import (
	"fmt"
	"strings"
)

type PatchField struct {
	Name  string
	Value any
	Skip  bool
}

func BuildPatchQuery(
	table string,
	idField string,
	id any,
	fields []PatchField,
) (string, []any) {
	setClauses := []string{}
	args := []any{}

	for _, field := range fields {
		if !field.Skip {
			setClauses = append(setClauses, fmt.Sprintf("%s = ?", field.Name))
			args = append(args, field.Value)
		}
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s = ?",
		table,
		strings.Join(setClauses, ", "),
		idField,
	)
	args = append(args, id)

	return query, args
}

func UpdateToNowQuery(
	table string,
	fieldToUpdate string,
	row string,
	value any,
) (string, []any) {
	return fmt.Sprintf("UPDATE %s SET %s = datetime('now') WHERE %s = ?", table, fieldToUpdate, row), []any{value}
}

func UpdateUpdatedAtQuery(
	table string,
	row string,
	value any,
) (string, []any) {
	return UpdateToNowQuery(table, "updatedAt", row, value)
}
