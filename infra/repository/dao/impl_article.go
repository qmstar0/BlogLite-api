package dao

import (
	"blog/domain/articles"
	"blog/domain/articles/valueobject"
	"blog/infra/e"
	"blog/infra/repository/model"
	"context"
)

func (d *Dao) NewArticle(c context.Context, art *articles.ArticleMate) error {
	var (
		err          error
		artMateModel = &model.ArticleMate{ArticleMate: art}
	)
	if err = d.db.WithContext(c).Model(&model.ArticleMate{}).Create(artMateModel).Error; err != nil {
		return e.NewError(e.DBCreateErr, err)
	}
	return nil
}

func (d *Dao) UptArticle(c context.Context, art *articles.ArticleMate) error {
	var (
		artMateModel = &model.ArticleMate{ArticleMate: art}
	)
	if err := d.db.WithContext(c).Model(&model.ArticleMate{}).Where("aid = ?", artMateModel.Aid).Updates(artMateModel).Error; err != nil {
		return e.NewError(e.DBUpdateErr, err)
	}
	return nil
}

func (d *Dao) DelArticle(c context.Context, art *articles.ArticleMate) error {
	var (
		artMateModel = &model.ArticleMate{ArticleMate: art}
	)
	if err := d.db.WithContext(c).Model(&model.ArticleMate{}).Where(artMateModel).Delete(artMateModel).Error; err != nil {
		return e.NewError(e.DBDeleteErr, err)
	}
	return nil
}

func (d *Dao) GetArticle(c context.Context, art *articles.ArticleMate) (*articles.ArticleMate, error) {
	var (
		artMateModel = &model.ArticleMate{ArticleMate: art}
	)
	if result := d.db.WithContext(c).Model(&model.ArticleMate{}).Where(artMateModel).Limit(1).Find(artMateModel); result.Error != nil {
		return nil, e.NewError(e.DBFindErr, result.Error)
	} else if result.RowsAffected == 0 {
		return nil, e.NewError(e.ItemNotExist, nil)
	}
	return artMateModel.ArticleMate, nil
}

func (d *Dao) AllArticle(c context.Context, limit int, offset int, isDraft, isTrash bool) ([]*articles.ArticleMate, error) {
	var (
		articleModels = make([]*model.ArticleMate, 0)
		tx            = d.db.WithContext(c).Model(&model.ArticleMate{})
	)
	tx.Limit(limit).Offset(offset)
	if isTrash {
		tx.Where("status = ?", valueobject.Deleted)
	} else if isDraft {
		tx.Where("status = ?", valueobject.Draft)
	} else {
		tx.Where("status = ?", valueobject.Published)
	}
	if err := tx.Find(&articleModels).Error; err != nil {
		return nil, e.NewError(e.DBFindErr, err)
	}
	var articlesEntity = make([]*articles.ArticleMate, len(articleModels))
	for i, art := range articleModels {
		articlesEntity[i] = art.ArticleMate
	}
	return articlesEntity, nil
}
