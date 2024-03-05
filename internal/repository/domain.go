package repository

import (
	"errors"
	"postfixadmin/internal/model"

	"gorm.io/gorm"
)

type DomainRepository interface {
	FirstById(name string) (*model.Domain, error)
}

func NewDomain(repository *Repository) DomainRepository {
	return &domainRepository{Repository: repository}
}

type domainRepository struct {
	*Repository
}

func (r *domainRepository) FirstById(name string) (*model.Domain, error) {
	var domain model.Domain
	if err := r.DB().Where("domain = ?", name).First(&domain).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("domain not found")
		}
		return nil, err
	}

	return &domain, nil
}
