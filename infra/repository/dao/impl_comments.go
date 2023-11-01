package dao

import (
	"blog/domain/articles"
	"blog/infra/repository/model"
	"context"
)

// NewComment 创建分类
func (d *Dao) NewComment(c context.Context, comment *articles.Comments) error {
	var (
		err          error
		commentModel = &model.Comments{Comments: comment}
	)
	if err = d.db.WithContext(c).Model(&model.Comments{}).Create(commentModel).Error; err != nil {
		return err
	}
	return nil
}

// DelComment 删除分类
func (d *Dao) DelComment(c context.Context, comment *articles.Comments) error {
	var (
		err          error
		commentModel = &model.Comments{Comments: comment}
	)
	if err = d.db.WithContext(c).Model(&model.Comments{}).Delete(commentModel).Error; err != nil {
		return err
	}
	return nil
}
