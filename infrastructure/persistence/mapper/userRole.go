package mapper

import (
	"blog/domain/aggregate/userRole"
	"blog/infrastructure/persistence/model"
)

func UserRoleDomainToModel(role userRole.UserRole) *model.UserM {
	r := role.(*userRole.UserRoleImpl)

	return &model.UserM{
		Uid:  r.Uid,
		Role: r.Role,
	}
}

func UserRoleModelToDomain(user *model.UserM) *userRole.UserRoleImpl {
	return &userRole.UserRoleImpl{
		Uid:  user.Uid,
		Role: user.Role,
	}
}
