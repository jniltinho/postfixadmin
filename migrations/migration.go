package migrations

import (
	_ "embed"
	"log/slog"
	"regexp"

	"gorm.io/gorm"
)

//go:embed scripts/postfixadmin.sql
var schema []byte

func InitMigrations(db *gorm.DB) {

	query := string(schema)

	lines := splitByEmptyNewline(query)
	for number, line := range lines {
		if len(line) > 0 {
			if err := db.Exec(line).Error; err != nil {
				slog.Error("Cannot create schema", err, "Line:", number)
			}
		}
	}

}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.MustCompile("\r\n").ReplaceAllString(str, "\n")
	return regexp.MustCompile(`\n\s*\n`).Split(strNormalized, -1)
}
