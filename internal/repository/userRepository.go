package repository

import (
	"errors"
	"log"

	"go-ecommerce-app2/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// in this file there is port(interface => public), adapter(struct => private) and constructor returning structure

/* port is used by service and constructor will be called in handler to inject in service which is already taking port
there it will connect to adapter
*/

// UserRepository
// @Description: Creating Port(interface)
type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserByID(id uint) (domain.User, error)
	UpdateUser(id uint, user domain.User) (domain.User, error)

	// bank
	CreateBankAccount(e domain.BankAccount) error

	// miscellaneous
	DeleteAllUser() error
}

// userRepository
// @Description: Creating Adapter(struct) returned by constructor, can't call directly that's why it is private
type userRepository struct {
	db *gorm.DB
}

// CreateBankAccount implements UserRepository.
func (ur userRepository) CreateBankAccount(e domain.BankAccount) error {
	result := ur.db.Create(&e)
	if result.Error != nil {
		return errors.New("failed to create bank account")
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		db: db,
	}
}

func (ur userRepository) CreateUser(user domain.User) (domain.User, error) {
	result := ur.db.Create(&user)
	if result.Error != nil {
		return domain.User{}, errors.New("failed to create user")
	}
	return user, nil
}

func (ur userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User
	err := ur.db.First(&user, "email=?", email).Error
	if err != nil {
		log.Println("user find error:", err)
		return domain.User{}, err
	}
	return user, nil
}

func (ur userRepository) FindUserByID(id uint) (domain.User, error) {
	var user domain.User
	err := ur.db.First(&user, id).Error
	if err != nil {
		log.Println("user find error", err)
		return domain.User{}, err
	}
	return user, nil
}

func (ur userRepository) UpdateUser(id uint, user domain.User) (domain.User, error) {
	err := ur.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(user).Error
	if err != nil {
		log.Println("error on update", err)
		return domain.User{}, err
	}
	return user, nil
}

func (ur userRepository) DeleteAllUser() error {
	ur.db.Exec("DELETE FROM users")
	log.Println("Delete All user function triggered")
	return nil
}
