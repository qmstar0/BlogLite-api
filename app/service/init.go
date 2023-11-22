package service

import (
	"blog/domain/articles"
	"blog/domain/auth"
	"blog/domain/users"
	"blog/infra/repository/dao"
	"blog/infra/repository/database/mysql"
	"blog/infra/repository/database/redis"
)

var d = dao.NewDao(mysql.GetDB(), redis.GetCacheClient())
var domainSrv = &DomainService{
	articleDTO: articles.NewServiceArticle(d),
	userDTO:    users.NewServiceUser(d),
	authDTO:    auth.NewServiceAuth(),
}

func GetSrv() *DomainService {
	return domainSrv
}
