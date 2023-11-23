package handlers

import (
	"blog/domain/users"
	"blog/infra/repository/dao"
	"blog/infra/repository/database/mysql"
	"blog/infra/repository/database/redis"
)

type DomainService struct {
	*users.ServiceUser
}

var (
	Srv = DomainService{
		users.NewServiceUser(),
	}
	Dao = dao.NewDao(mysql.GetDB(), redis.GetCacheClient())
)
