package adapters

import (
	"blog/pkg/jwtAuth"
	"common/auth"
	"users/domain/user"
)

type userAuthService struct {
	authCli *jwtAuth.JwtAuth[*user.User]
}

func NewUserAuthServce() user.UserAuthService {
	return &userAuthService{
		authCli: auth.NewAuthClient[*user.User](),
	}
}

func (u userAuthService) SignToken(user *user.User) (string, error) {
	return u.authCli.Sign(user)
}

func (u userAuthService) ParseToken(token string) (*user.User, error) {
	claims, _, err := u.authCli.Parse(token)
	return claims, err
}

func (u userAuthService) IsAdmin(user *user.User) bool {
	//TODO implement me
	panic("implement me")
}
