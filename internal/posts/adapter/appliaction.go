package adapter

import (
	"blog/pkg/mongodb"
	"blog/pkg/rediscache"
	"categorys/application"
	"categorys/application/command"
	"categorys/application/query"
)

func NewApp() *application.App {
	cacher := rediscache.GetCacher()
	db := mongodb.GetDB()
	cateRepo := NewCategoryRepository(db)

	return &application.App{
		Commands: application.Commands{
			CreateCategory: command.NewCreateCategoryHandler(cateRepo),
			UpdateCategory: command.NewUpdataCategoryHandler(cateRepo),
			DeleteCategory: command.NewDeleteCategoryHandler(cateRepo),
		},

		Queries: application.Queries{
			GetCategory:    query.NewGetCategoryHandler(cateRepo, cacher),
			GetAllCategory: query.NewGetAllCategoryHandler(cateRepo, cacher),
		},
	}
}
