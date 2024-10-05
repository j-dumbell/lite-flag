package pg

import (
	"database/sql"
)

func ToNullString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{
			Valid: false,
		}
	}

	return sql.NullString{
		String: *value,
		Valid:  true,
	}
}

func FromNullString(nullString sql.NullString) *string {
	if !nullString.Valid {
		return nil
	}

	return &nullString.String
}

func ToNullBool(value *bool) sql.NullBool {
	if value == nil {
		return sql.NullBool{
			Valid: false,
		}
	}

	return sql.NullBool{
		Bool:  *value,
		Valid: true,
	}
}

func FromNullBool(nullBool sql.NullBool) *bool {
	if !nullBool.Valid {
		return nil
	}

	return &nullBool.Bool
}
