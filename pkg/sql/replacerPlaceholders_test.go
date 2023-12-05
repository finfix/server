package sql

import (
	"testing"
)

func TestReplacingPlaceholders(t *testing.T) {

	for _, tt := range []struct {
		message string
		sql     string
		result  string
	}{
		{"1.Обычный запрос с несколькими параметрами",
			`SELECT * FROM users WHERE id = ? AND name = ?`,
			`SELECT * FROM users WHERE id = $1 AND name = $2`,
		},
		{"2.Запрос без параметров",
			`SELECT * FROM users`,
			`SELECT * FROM users`,
		},
		//{"3.Запрос с параметрами в виде вопросительного знака и доллара",
		//	`SELECT * FROM users WHERE id = $1 AND name = ?`,
		//	`SELECT * FROM users WHERE id = $1 AND name = $2`,
		//},
		{"4.Запрос в виде доллара",
			`SELECT * FROM users WHERE id = $1 AND name = $2`,
			`SELECT * FROM users WHERE id = $1 AND name = $2`,
		},
		{
			"5.Запрос с параметрами для INSERT",
			`INSERT INTO users (id, name) VALUES (?, ?)`,
			`INSERT INTO users (id, name) VALUES ($1, $2)`,
		},
	} {
		t.Run(tt.message, func(t *testing.T) {
			if tt.result != replacePlaceholders(tt.sql) {
				t.Fatalf("\n\nОжидалось: %v\nПолучено: %v\n\n", tt.result, replacePlaceholders(tt.sql))
			}
		})
	}
}
