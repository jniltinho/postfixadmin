package database

import (
	_ "embed"
	"log/slog"
	"regexp"

	"github.com/spf13/viper"
)

//go:embed scripts/postfixadmin.sql
var schema []byte

func CreateSchema(conf *viper.Viper) {

	query := string(schema)

	lines := splitByEmptyNewline(query)
	for number, line := range lines {
		if len(line) > 0 {
			if err := DB().Exec(line).Error; err != nil {
				slog.Error("Cannot create schema", err, "Line:", number)
			}
		}
	}

}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.MustCompile("\r\n").ReplaceAllString(str, "\n")
	return regexp.MustCompile(`\n\s*\n`).Split(strNormalized, -1)
}
