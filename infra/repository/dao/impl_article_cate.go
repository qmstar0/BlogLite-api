package dao

import (
	"blog/domain/articles"
	"blog/infra/e"
	"blog/infra/repository/model"
	"context"
)

func (d *Dao) NewCate(c context.Context, cate *articles.ArticleCategory) error {
	var (
		cateModel = &model.ArticleCate{ArticleCategory: cate}
	)
	if err := d.db.WithContext(c).Model(&model.ArticleCate{}).Create(cateModel).Error; err != nil {
		return e.NewError(e.DBCreateErr, err)
	}
	return nil
}

func (d *Dao) UptCate(c context.Context, cate *articles.ArticleCategory) error {
	var (
		cateModel = &model.ArticleCate{ArticleCategory: cate}
	)
	if err := d.db.WithContext(c).Model(&model.ArticleCate{}).Where("id = ?", cateModel.Id).Updates(cateModel).Error; err != nil {
		return e.NewError(e.DBUpdateErr, err)
	}
	return nil
}

func (d *Dao) DelCate(c context.Context, cate *articles.ArticleCategory) error {
	var (
		cateModel = &model.ArticleCate{ArticleCategory: cate}
	)
	if err := d.db.WithContext(c).Model(&model.ArticleCate{}).Where(cateModel).Delete(cateModel).Error; err != nil {
		return e.NewError(e.DBDeleteErr, err)
	}
	return nil
}

func (d *Dao) GetCate(c context.Context, cate *articles.ArticleCategory) (*articles.ArticleCategory, error) {
	var (
		cateModel = &model.ArticleCate{ArticleCategory: cate}
	)
	result := d.db.WithContext(c).Model(&model.ArticleCate{}).Where("id = ?", cateModel.Id).Limit(1).Find(cateModel)
	if result.RowsAffected == 0 {
		return nil, e.NewError(e.ItemNotExist, nil)
	} else if result.Error != nil {
		return nil, e.NewError(e.DBFindErr, result.Error)
	}
	return cateModel.ArticleCategory, nil
}

func (d *Dao) AllCate(c context.Context) ([]*articles.ArticleCategory, error) {
	var (
		cateModels = make([]*model.ArticleCate, 0)
	)
	if err := d.db.WithContext(c).Model(&model.ArticleCate{}).Find(&cateModels).Error; err != nil {
		return nil, e.NewError(e.DBFindErr, err)
	}
	var cateVos = make([]*articles.ArticleCategory, len(cateModels))
	for i, cate := range cateModels {
		cateVos[i] = cate.ArticleCategory
	}
	return cateVos, nil
}
