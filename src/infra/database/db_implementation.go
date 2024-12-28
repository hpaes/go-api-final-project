package database

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hpaes/go-api-final-project/src/core/domain"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/hpaes/go-api-final-project/src/infra/config"
	"github.com/hpaes/go-api-final-project/src/infra/database/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlConnection struct {
	db *gorm.DB
}

var _ DbUsers = (*SqlConnection)(nil)

func NewSqlConnection(config config.MySQL) (*SqlConnection, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DBName)
	mysqlConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return &SqlConnection{}, err
	}

	return &SqlConnection{db: mysqlConn}, nil
}

// GetByEmail implements DbUsers.
func (s *SqlConnection) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var userEntity entity.User
	result := s.db.Where("email = ?", email).First(&userEntity)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	if userEntity.Email == "" {
		return domain.User{}, custom_errors.NewUserNotFoundError("User", email)
	}

	user, err := userEntity.ToDomain()
	if err != nil {
		return domain.User{}, err
	}

	return *user, nil
}

// Save implements DbUsers.
func (s *SqlConnection) Save(ctx context.Context, user *domain.User) (domain.User, error) {
	userEntity, err := entity.NewUserEntity(user)
	if err != nil {
		return domain.User{}, err
	}

	result := s.db.Create(&userEntity)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	user.Id = userEntity.Id
	return *user, nil
}

// GetById implements DbUsers.
func (s *SqlConnection) GetById(ctx context.Context, id int) (domain.User, error) {
	var userEntity entity.User
	result := s.db.First(&userEntity, id)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	if userEntity.Id == 0 {
		return domain.User{}, custom_errors.NewUserNotFoundError("User", strconv.Itoa(id))
	}

	user, err := userEntity.ToDomain()
	if err != nil {
		return domain.User{}, err
	}

	return *user, nil
}

// Update implements DbUsers.
func (s *SqlConnection) Update(ctx context.Context, user *domain.User) (domain.User, error) {
	userEntity, err := entity.NewUserEntity(user)
	if err != nil {
		return domain.User{}, err
	}

	result := s.db.Save(&userEntity)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return *user, nil
}

// Delete implements DbUsers.
func (s *SqlConnection) Delete(ctx context.Context, id int) error {
	result := s.db.Delete(&entity.User{}, id)
	if result.RowsAffected == 0 {
		return custom_errors.NewUserNotFoundError("User", strconv.Itoa(id))
	}

	return result.Error
}

// List implements DbUsers.
func (s *SqlConnection) List(ctx context.Context, page int) ([]domain.User, error) {
	var userEntities []entity.User
	limit := 10
	offset := (page - 1) * limit

	result := s.db.Limit(limit).Offset(offset).Find(&userEntities)
	if result.Error != nil {
		return nil, result.Error
	}

	var users []domain.User
	for _, userEntity := range userEntities {
		user, err := userEntity.ToDomain()
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}

	return users, nil
}
