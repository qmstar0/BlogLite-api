package commandConstructor

import (
	"blog/adapter/httpAdapter/authorize"
	"blog/adapter/httpAdapter/bind"
	"blog/apps/dto"
	"blog/domain/aggregate/categorys"
	"net/http"
)

func CreateCategory(r *http.Request) (cmd any, err error) {
	var req = dto.CategoryCreateD{}

	err = bind.Decode(r, &req)
	if err != nil {
		return nil, err
	}

	claims, err := authorize.ParseToClaims(r)
	if err != nil {
		return nil, err
	}

	cmd = categorys.CreateCategoryCommand{
		Uid:         claims.Uid,
		Name:        req.Name,
		DisplayName: req.DisplayName,
		SeoDesc:     req.SeoDesc,
	}

	return cmd, nil
}
