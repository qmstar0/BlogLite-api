package service

import (
	"github.com/qmstar0/nightsky-api/config"
	"github.com/qmstar0/nightsky-api/internal/apps/query"
	"github.com/qmstar0/nightsky-api/internal/pkg/e"
	"github.com/qmstar0/nightsky-api/pkg/auth"
	"time"
)

type AdminAuthenticationService struct {
	authenticator *auth.JWTAuthenticator
}

func NewAdminAuthenticationService() *AdminAuthenticationService {
	return &AdminAuthenticationService{
		authenticator: auth.NewJWTAuthenticator(
			[]byte(config.Cfg.JWTAuth.AuthKey),
			config.Cfg.JWTAuth.Subject,
			config.Cfg.JWTAuth.Issuer,
			config.Cfg.JWTAuth.Audience...,
		),
	}
}

func (a AdminAuthenticationService) GenerateAdminToken(duration time.Duration) (query.AdminTokenView, error) {
	token, claims, err := a.authenticator.Sign(duration)
	if err != nil {
		return query.AdminTokenView{}, e.AErrSignError.WithError(err)
	}
	issuedAt, _ := claims.GetIssuedAt()
	expirationTime, _ := claims.GetExpirationTime()
	return query.AdminTokenView{
		Token:     token,
		Timestamp: issuedAt.Unix(),
		Exp:       expirationTime.Unix(),
	}, nil
}

func (a AdminAuthenticationService) Verify(tokenStr string) error {
	claims, err := a.authenticator.Parse(tokenStr)
	if err != nil {
		return e.AErrWrongAuthortion.WithError(err)
	}
	if issuer, err := claims.GetIssuer(); err != nil || config.Cfg.JWTAuth.Issuer != issuer {
		return e.AErrWrongAuthortion.WithError(err).WithMessage("无效的Token")
	}
	return nil
}
