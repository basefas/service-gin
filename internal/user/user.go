package user

import (
	"basefas.com/service-gin/internal/auth"
	"basefas.com/service-gin/internal/utils/db"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("User not found.")

	ErrIncorrectUsernameOrPassword = errors.New("Incorrect username or password.")

	ErrUsernameOrPasswordNil = errors.New("Username or password can not be null.")

	ErrGenerateTokenFailed = errors.New("Generate token failed")
)

func Create(cu CreateUser) error {
	u := User{Username: cu.Username, Password: cu.Password, Email: cu.Email}

	err := db.Mysql.Table("user").Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func Get(uid string) (*GetUserInfo, error) {
	var u = GetUserInfo{}

	err := db.Mysql.
		Table("user").
		Where("id = ?", uid).
		Where("deleted_at IS NULL").
		Scan(&u).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &u, nil
}

func Update(uid string, uu UpdateUser) error {
	_, err := Get(uid)
	if err != nil {
		return err
	}

	user := make(map[string]interface{})
	if uu.Username != "" {
		user["username"] = uu.Username
	}
	if uu.Password != "" {
		user["password"] = uu.Password
	}
	if uu.Email != "" {
		user["email"] = uu.Email
	}

	err = db.Mysql.
		Debug().
		Table("user").
		Where("id = ?", uid).
		Where("deleted_at IS NULL").
		Updates(user).Error
	return err
}

func Delete(uid string) error {
	_, err := Get(uid)
	if err != nil {
		return err
	}

	var u = User{}

	err = db.Mysql.Table("user").Where("id = ?", uid).Delete(&u).Error
	return err
}

func List() ([]GetUserInfo, error) {
	var users []GetUserInfo

	err := db.Mysql.
		Table("user").
		Where("deleted_at IS NULL").
		Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func Token(pl Login) (string, error) {

	if len(pl.Username) <= 0 || len(pl.Password) <= 0 {
		return "", ErrUsernameOrPasswordNil
	}

	var user User
	err := db.Mysql.
		Where("username = ? ", pl.Username).
		Where("password = ? ", pl.Password).
		Find(&user).Error

	if err != nil {
		return "", ErrIncorrectUsernameOrPassword
	}

	token, tokenErr := auth.GenerateToken(user.ID)
	if tokenErr != nil {
		return "", ErrGenerateTokenFailed
	}
	return token, nil
}
