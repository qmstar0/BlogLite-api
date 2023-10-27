package dao

import (
	"blog/domain/articles"
	"blog/infra/e"
	"blog/infra/repository/model"
	"context"
)

func (d *Dao) NewTag(c context.Context, tag *articles.ArticleTags) error {
	var (
		tagModel = &model.ArticleTags{ArticleTags: tag}
	)
	if err := d.db.WithContext(c).Model(&model.ArticleTags{}).Create(tagModel).Error; err != nil {
		return e.NewError(e.DBCreateErr, err)
	}
	return nil
}
func (d *Dao) UptTag(c context.Context, tag *articles.ArticleTags) error {
	var (
		tagModel = &model.ArticleTags{ArticleTags: tag}
	)
	if err := d.db.WithContext(c).Model(&model.ArticleTags{}).Where("id = ?", tagModel.Id).Updates(tagModel).Error; err != nil {
		return e.NewError(e.DBCreateErr, err)
	}
	return nil
}

func (d *Dao) DelTag(c context.Context, tag *articles.ArticleTags) error {
	var (
		tagModel = &model.ArticleTags{ArticleTags: tag}
	)
	if err := d.db.WithContext(c).Model(&model.ArticleTags{}).Where(tagModel).Delete(tagModel).Error; err != nil {
		return e.NewError(e.DBDeleteErr, err)
	}
	return nil
}

func (d *Dao) GetTag(c context.Context, tags []int) ([]*articles.ArticleTags, error) {
	var (
		tagModel = make([]*model.ArticleTags, len(tags))
	)
	result := d.db.WithContext(c).Model(&model.ArticleTags{}).Find(tagModel, &tags)
	if result.RowsAffected == 0 {
		return nil, e.NewError(e.ItemNotExist, nil)
	} else if result.Error != nil {
		return nil, e.NewError(e.DBFindErr, result.Error)
	}
	var artTags = make([]*articles.ArticleTags, len(tagModel))
	for i, t := range tagModel {
		artTags[i] = t.ArticleTags
	}
	return artTags, nil
}

func (d *Dao) AllTag(c context.Context) ([]*articles.ArticleTags, error) {
	var (
		tagModels = make([]*model.ArticleTags, 0)
	)
	if err := d.db.WithContext(c).Model(&model.ArticleTags{}).Find(&tagModels).Error; err != nil {
		return nil, e.NewError(e.DBFindErr, err)
	}
	var tags = make([]*articles.ArticleTags, len(tagModels))
	for i, tag := range tagModels {
		tags[i] = tag.ArticleTags
	}
	return tags, nil
}
