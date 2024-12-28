package entity

import (
	"github.com/hpaes/go-api-final-project/src/core/domain"
)

type (
	User struct {
		Id    int    `gorm:"type:int:primaryKey"`
		Name  string `gorm:"type:varchar(100);not null"`
		Email string `gorm:"type:varchar(100);not null"`
		Age   int    `gorm:"type:int;not null"`
	}
)

func NewUserEntity(p *domain.User) (*User, error) {
	if p.Id > 0 {
		return &User{
			Id:    p.Id,
			Name:  p.Name,
			Email: p.Email.Value(),
			Age:   p.Age,
		}, nil
	}
	return &User{
		Name:  p.Name,
		Email: p.Email.Value(),
		Age:   p.Age,
	}, nil
}

func (u *User) ToDomain() (*domain.User, error) {
	user, err := domain.Restore(u.Id, u.Name, u.Email, u.Age)
	if err != nil {
		return nil, err
	}

	return user, nil
}
