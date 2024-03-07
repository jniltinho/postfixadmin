package db

import (
	"context"
	"database/sql"
	"postfixadmin/config"
	"postfixadmin/types"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

// DBInstance is a singleton DB instance
type BunDB struct {
	Conf *viper.Viper
	Bun  *bun.DB
}

var Bun *bun.DB

func createBunDB(cfg *viper.Viper) (*sql.DB, error) {
	//ctx := context.Background()
	DSN := config.GetMysqlDSN(cfg)

	// Open a MySQL 5.7+ database.
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitBun(conf *viper.Viper) error {

	db, err := createBunDB(conf)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	Bun = bun.NewDB(db, mysqldialect.New())
	if conf.GetString("mysql.log_level") == "debug" {
		Bun.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}
	return nil
}

func GetDomainByName(name string) (types.Domain, error) {
	var domain types.Domain
	err := Bun.NewSelect().Model(&domain).Where("domain = ?", name).Scan(context.Background())
	//log.DEBUG("Domain: %v", domain)
	return domain, err
}

func GetAllDomains() ([]types.Domain, error) {
	var domains []types.Domain
	err := Bun.NewSelect().Model(&domains).Scan(context.Background())
	//for _, domain := range domains {
	//	log.DEBUG("Domain: %s - %s -> Active %v", domain.Domain, domain.Description, domain.Active)
	//}
	return domains, err
}
