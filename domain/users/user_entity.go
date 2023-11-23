package users

import (
	"blog/domain/users/valueobject"
	"blog/infra/config"
	"blog/infra/e"
	"blog/infra/jwt"
	"github.com/google/uuid"
	"time"
)

var userAuthTokenLifeTime = time.Duration(config.Conf.User.JwtAuthTokenLifeDay) * time.Hour * 24

type User struct {
	Id             uint                 `json:"-" gorm:"primaryKey"`
	Uid            string               `json:"uid" gorm:"column:uid; uniqueIndex; not null"`
	UserName       string               `json:"user_name" gorm:"column:user_name; type:varchar(255); not null"`
	FormerUserName string               `json:"former_user_name" gorm:"column:former_user_name; type:varchar(255)"`
	Email          valueobject.Email    `json:"user_email" gorm:"column:user_email; type:varchar(255); uniqueIndex; not null"`
	Role           valueobject.Role     `json:"role" gorm:"column:role; type:TINYINT UNSIGNED; not null"`
	Password       valueobject.Passowrd `json:"password" gorm:"column:password; type:varchar(255)"`
	CreateAt       uint                 `json:"create_at" gorm:"column:create_at; not null"`
	UpdateAt       uint                 `json:"update_at" gorm:"column:update_at; default:0; not null"`
	DeleteAt       uint                 `json:"delete_at" gorm:"column:delete_at; default:0; not null"`
}

// NewUser 新建空用户
func NewUser(email string, role uint) (*User, error) {
	newEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}
	password, _ := valueobject.NewPassword("")
	return &User{
		Uid:      uuid.New().String(),
		UserName: newEmail.GetNickName(),
		Email:    newEmail,
		Password: password,
		Role:     valueobject.NewUserRole(role),
	}, nil
}

// ResetPassword 重置密码
func (u *User) ResetPassword(newpwd string) error {
	var (
		err error
	)
	if !u.Password.Check(newpwd) {
		u.Password, err = valueobject.NewPassword(newpwd)
		if err != nil {
			return e.NewError(e.PwdEncryptionErr, nil)
		}
	}
	return nil
}

// UpdateUserName 更新用户
func (u *User) UpdateUserName(newName string) error {
	if newName == u.UserName {
		return e.NewError(e.InvalidUpdate, nil)
	}
	if newName != "" {
		u.FormerUserName = u.UserName
		u.UserName = newName
	}
	return nil
}

func (u *User) GenUserAuthToken() (string, error) {
	data := map[string]any{
		"userId": u.Uid,
		"name":   u.UserName,
		"role":   u.Role.ToUint(),
		"email":  u.Email.ToString(),
	}
	return jwt.Sign(data, userAuthTokenLifeTime)
}
