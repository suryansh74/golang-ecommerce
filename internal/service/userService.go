package service

import (
	"errors"
	"go-ecommerce-app2/internal/domain"
	"go-ecommerce-app2/internal/dto"
	"go-ecommerce-app2/internal/helper"
	"go-ecommerce-app2/internal/repository"
	"gorm.io/gorm"
	"log"
	"time"
)

// UserService
// @Description: Inside service inject interface(port) not repository(adapter) domain UserRepo
type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (us UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := us.Repo.FindUser(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // user does NOT exist
		}
		return nil, err // real DB error
	}
	return &user, nil // user found
}

func (us UserService) SignUp(input dto.UserSignUp) (string, error) {
	// in param we are not using email, phone, password explicitly because in
	// future there could be other input fields as well

	// check if user email is already exists
	log.Println("sign up function is called")
	existingUser, err := us.findUserByEmail(input.Email)
	if existingUser != nil {
		log.Println("user already exists")
		return "", errors.New("user already exists")
	}

	hashedPassword, err := us.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := us.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	})
	if err != nil {
		return "", err
	}

	//generate token
	return us.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (us UserService) Login(email string, password string) (string, error) {
	// check if user is present or not
	user, err := us.Repo.FindUser(email)
	if err != nil {
		return "", errors.New("user does not exist, please provide valid email address")
	}

	// check password is valid or not
	err = us.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}
	// generate token and return

	return us.Auth.GenerateToken(user.ID, user.Email, user.UserType)

}

func (us UserService) isVerifiedUser(id uint) bool {
	currentUser, err := us.Repo.FindUserByID(id)
	return err == nil && currentUser.Verified
}

func (us UserService) GetVerificationCode(user domain.User) (int, error) {
	// check if user is already verified
	if us.isVerifiedUser(user.ID) {
		return 0, errors.New("user already verified")
	}

	// generate verification code
	log.Println("user is not verified confirmed")
	code, err := us.Auth.GenerateCode()

	// ADD THIS DEBUG LINE
	log.Printf("GenerateCode returned: code=%d, err=%v", code, err)

	if err != nil {
		return 0, err
	}

	// ADD THIS DEBUG LINE TOO
	log.Printf("Code before update: %d", code)

	updatedUser := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = us.Repo.UpdateUser(user.ID, updatedUser)
	if err != nil {
		return 0, errors.New("unable to update verification code")
	}

	// ADD THIS DEBUG LINE
	log.Printf("Returning code: %d", code)

	return code, nil
}

func (us UserService) VerifyCode(id uint, code int) error {
	if us.isVerifiedUser(id) {
		return errors.New("user is already verified")
	}

	// get user from id
	user, err := us.Repo.FindUserByID(id)
	if err != nil {
		return err
	}

	// check if code matched with user's code
	if user.Code != code {
		return errors.New("verification code doesn't match")
	}

	// check if code expired or not
	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	// update user
	udpateUser := domain.User{
		Verified: true,
	}

	_, err = us.Repo.UpdateUser(id, udpateUser)
	if err != nil {
		return errors.New("unable to update user")
	}
	return nil

}

func (us UserService) CreateProfile(id uint, input any) error {

	return nil
}

func (us UserService) GetProfile(input any) (*domain.User, error) {

	return nil, nil
}

func (us UserService) UpdateProfile(id int, input any) (string, error) {

	return "", nil
}

func (us UserService) BecomeSeller(id int, input any) (string, error) {

	return "", nil
}

func (us UserService) FindCart(id uint) ([]interface{}, error) {

	return nil, nil
}

func (us UserService) CreateCart(input any, user domain.User) ([]interface{}, error) {

	return nil, nil
}

func (us UserService) CreateOrder(user domain.User) (int, error) {

	return 0, nil
}

func (us UserService) GetOrders(user domain.User) ([]interface{}, error) {

	return nil, nil
}

func (us UserService) GetOrderByID(id uint, uID int) (interface{}, error) {

	return nil, nil
}
