package database

import (
	_ "embed"
	"regexp"

	"postfixadmin/log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//go:embed scripts/postfixadmin.sql
var schema []byte

func CreateSchema(conf *viper.Viper, zlog *log.Logger) {

	query := string(schema)

	lines := splitByEmptyNewline(query)
	for number, line := range lines {
		if len(line) > 0 {
			if err := DB().Exec(line).Error; err != nil {
				zlog.Fatal("Cannot create schema", zap.Error(err), zap.Int("Line:", number))
			}
		}
	}

}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.MustCompile("\r\n").ReplaceAllString(str, "\n")
	return regexp.MustCompile(`\n\s*\n`).Split(strNormalized, -1)
}
