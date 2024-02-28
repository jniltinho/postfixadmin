package app

import (
	"embed"
	"os"
	"postfixadmin/util"
	"time"
)

//go:embed scripts/local.toml
var configFile []byte

//go:embed static/*
var FS embed.FS

func InitConfigFile() {
	filename := GetFileNameDate()
	err := os.WriteFile(filename, configFile, 0644)
	if err == nil {
		util.LOG("Config file created %s", filename)
	} else {
		util.LOG("Error creating config file %s", err.Error())
	}
}

func GetFileNameDate() string {
	const layout = "20060102"
	t := time.Now()
	return "local-" + t.Format(layout) + ".toml"
}
