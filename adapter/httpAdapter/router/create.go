package router

import (
	"blog/adapter/httpAdapter/authorize"
	"blog/adapter/httpAdapter/bind"
	"blog/domain/aggregate/categorys"
	"net/http"
)

type CreateCategoryDTO struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	SeoDesc     string `json:"seoDesc"`
}

func CreateCategroy(r *http.Request) (cmd any, err error) {

	claims, err := authorize.ParseToClaims(r)
	if err != nil {
		return nil, err
	}

	var dto CreateCategoryDTO
	err = bind.Decode(r, &dto)
	if err != nil {
		return nil, err
	}

	return categorys.CreateCategoryCommand{
		Uid:         claims.Uid,
		Name:        dto.Name,
		DisplayName: dto.DisplayName,
		SeoDesc:     dto.SeoDesc,
	}, nil
}
