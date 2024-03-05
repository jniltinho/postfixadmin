package internal

import (
	"fmt"
	"postfixadmin/internal/repository"

	"github.com/spf13/viper"
)

func TesteInternal(viperViper *viper.Viper) {

	db := repository.NewDB(viperViper)
	repo := repository.NewRepository(db)
	domainRepo := repository.NewDomain(repo)

	domain, _ := domainRepo.FirstById("linuxpro.com.br")
	fmt.Printf("Domain: %v\n", domain)
}
