package articles

import (
	"blog/app/dto"
	"blog/domain/articles/valueobject"
	"context"
)

type RepoArticle interface {
	RepoArticleMate
	RepoCate
	RepoTags
}

// ServiceArticle 文章服务
type ServiceArticle struct {
	repo RepoArticle
}

// NewServiceArticle 构造方法
func NewServiceArticle(d RepoArticle) *ServiceArticle {
	return &ServiceArticle{repo: d}
}

func (s ServiceArticle) GetArticleDetailList(c context.Context, limit int, offset int, isDraft bool, isTrash bool) ([]dto.ArticleListDisplay, error) {
	var err error
	arts, err := s.repo.AllArticle(c, limit, offset, isDraft, isTrash)
	if err != nil {
		return nil, err
	}
	var (
		artD       dto.ArticleDisplay
		cateD      dto.CateDisplay
		artDateilD []dto.ArticleListDisplay
	)
	for _, art := range arts {
		artD = dto.ArticleDisplay{
			Id:        0,
			Aid:       art.Aid,
			Uid:       art.Uid,
			Title:     art.Title,
			Summary:   art.Summary,
			Content:   art.Content,
			PublishAt: art.PublishAt,
			CreatedAt: art.CreateAt,
			UpdatedAt: art.UpdateAt,
		}
		tags, err := s.repo.GetTag(c, art.TagIds)
		if err != nil {
			return nil, err
		}
		var tagD []dto.TagDisplay
		for _, tag := range tags {
			tagD = append(tagD, dto.TagDisplay{
				Id:          0,
				Name:        tag.Name,
				DisplayName: tag.DisplayName,
				SeoDesc:     tag.SeoDesc,
				Num:         tag.Num,
			})
		}
		cate, err := s.repo.GetCate(c, &ArticleCategory{Id: art.CategoryId})
		if err != nil {
			return nil, err
		}
		cateD = dto.CateDisplay{
			Id:          0,
			Name:        cate.Name,
			DisplayName: cate.DisplayName,
			SeoDesc:     cate.SeoDesc,
		}

		artDateilD = append(artDateilD,
			dto.ArticleListDisplay{
				Article:  artD,
				Tags:     tagD,
				Category: cateD,
				Author:   dto.UserDisplay{},
			})
	}
	return artDateilD, nil
}

func (s ServiceArticle) GetArticleDetail(c context.Context, aid string) (dto.ArticleListDisplay, error) {
	art, err := s.repo.GetArticle(c, &ArticleMate{Aid: aid})
	if err != nil {
		return dto.ArticleListDisplay{}, err
	}
	aid = art.Aid
	artD := dto.ArticleDisplay{
		Id:        art.Id,
		Aid:       art.Aid,
		Uid:       art.Uid,
		Title:     art.Title,
		Summary:   art.Summary,
		Content:   art.Content,
		PublishAt: art.PublishAt,
		CreatedAt: art.CreateAt,
		UpdatedAt: art.UpdateAt,
	}
	tags, err := s.repo.GetTag(c, art.TagIds)
	if err != nil {
		return dto.ArticleListDisplay{}, err
	}
	var tagD = make([]dto.TagDisplay, len(tags))
	for i, tag := range tags {
		tagD[i] = dto.TagDisplay{
			Id:          tag.Id,
			Name:        tag.Name,
			DisplayName: tag.DisplayName,
			SeoDesc:     tag.SeoDesc,
			Num:         tag.Num,
		}
	}
	cate, err := s.repo.GetCate(c, &ArticleCategory{Id: art.CategoryId})
	if err != nil {
		return dto.ArticleListDisplay{}, err
	}
	cateD := dto.CateDisplay{
		Id:          cate.Id,
		Name:        cate.Name,
		DisplayName: cate.DisplayName,
		SeoDesc:     cate.SeoDesc,
	}
	artDateil := dto.ArticleListDisplay{
		Article:  artD,
		Tags:     tagD,
		Category: cateD,
		Author:   dto.UserDisplay{},
	}
	return artDateil, nil
}

func (s ServiceArticle) GetArticle(c context.Context, aid string) (dto.ArticleDisplay, error) {
	art, err := s.repo.GetArticle(c, &ArticleMate{Aid: aid})
	if err != nil {
		return dto.ArticleDisplay{}, err
	}
	return dto.ArticleDisplay{
		Id:        art.Id,
		Aid:       art.Aid,
		Uid:       art.Uid,
		Title:     art.Title,
		Summary:   art.Summary,
		Content:   art.Content,
		PublishAt: art.PublishAt,
		CreatedAt: art.CreateAt,
		UpdatedAt: art.UpdateAt,
	}, nil
}

func (s ServiceArticle) GetTag(c context.Context, tagIds []int) ([]dto.TagDisplay, error) {
	tag, err := s.repo.GetTag(c, tagIds)
	if err != nil {
		return nil, err
	}
	var tags = make([]dto.TagDisplay, len(tag))
	for i, t := range tag {
		tags[i] = dto.TagDisplay{
			Id:          t.Id,
			Name:        t.Name,
			DisplayName: t.DisplayName,
			SeoDesc:     t.SeoDesc,
			Num:         t.Num,
		}
	}
	return tags, nil
}

func (s ServiceArticle) AllTag(c context.Context) ([]dto.TagDisplay, error) {
	tags, err := s.repo.AllTag(c)
	if err != nil {
		return nil, err
	}
	var tagD = make([]dto.TagDisplay, len(tags))
	for i, tag := range tags {
		tagD[i] = dto.TagDisplay{
			Id:          tag.Id,
			Name:        tag.Name,
			DisplayName: tag.DisplayName,
			SeoDesc:     tag.SeoDesc,
			Num:         tag.Num,
		}

	}
	return tagD, nil
}

func (s ServiceArticle) GetCate(c context.Context, cateId int) (dto.CateDisplay, error) {
	cate, err := s.repo.GetCate(c, &ArticleCategory{Id: cateId})
	if err != nil {
		return dto.CateDisplay{}, err
	}
	return dto.CateDisplay{
		Id:          cate.Id,
		Name:        cate.Name,
		DisplayName: cate.DisplayName,
		SeoDesc:     cate.SeoDesc,
	}, nil
}

func (s ServiceArticle) GetCateByParentId(c context.Context, ParentId int) (dto.CateDisplay, error) {
	cate, err := s.repo.GetCate(c, &ArticleCategory{ParentId: ParentId})
	if err != nil {
		return dto.CateDisplay{}, err
	}
	return dto.CateDisplay{
		Id:          cate.Id,
		Name:        cate.Name,
		DisplayName: cate.DisplayName,
		SeoDesc:     cate.SeoDesc,
	}, nil
}

func (s ServiceArticle) AllCate(c context.Context) ([]dto.CateDisplay, error) {
	cateAll, err := s.repo.AllCate(c)
	if err != nil {
		return nil, err
	}
	var cateD = make([]dto.CateDisplay, len(cateAll))
	for i, cate := range cateAll {
		cateD[i] = dto.CateDisplay{
			Id:          cate.Id,
			Name:        cate.Name,
			DisplayName: cate.DisplayName,
			SeoDesc:     cate.SeoDesc,
		}
	}
	return cateD, nil
}

func (s ServiceArticle) NewArticle(c context.Context, uid string, store dto.ArticleStore) error {
	art := NewArticleMate(uid)
	art.SetTitle(store.Title)
	art.SetSummary(store.Summary)
	art.SetStatus(valueobject.Draft)
	art.SetCategory(store.Category)
	art.SetTagIDs(store.Tags)
	if err := art.SetContent(store.Content); err != nil {
		return err
	}
	if err := s.repo.NewArticle(c, art); err != nil {
		return err
	}
	return nil
}

func (s ServiceArticle) UpdateArticle(c context.Context, aid string, store dto.ArticleStore) error {
	art := &ArticleMate{Aid: aid}
	art.SetTitle(store.Title)
	art.SetSummary(store.Summary)
	art.SetStatus(valueobject.Draft)
	art.SetCategory(store.Category)
	art.SetTagIDs(store.Tags)
	if err := art.SetContent(store.Content); err != nil {
		return err
	}
	if err := s.repo.UptArticle(c, art); err != nil {
		return err
	}
	return nil
}

func (s ServiceArticle) DeleteArtcle(c context.Context, aid string) error {
	return s.repo.DelArticle(c, &ArticleMate{Aid: aid})
}

func (s ServiceArticle) PublishArticle(c context.Context, aid string) error {
	art := &ArticleMate{Aid: aid}
	art.Status.SetPublished()
	return s.repo.UptArticle(c, art)
}

func (s ServiceArticle) NewTag(c context.Context, store dto.TagStore) error {
	tag := NewArticleTags(store.Name, store.DisplayName, store.SeoDesc)
	return s.repo.NewTag(c, tag)
}

func (s ServiceArticle) UpdateTag(c context.Context, tagId int, store dto.TagStore) error {
	tag := &ArticleTags{
		Id:          tagId,
		Name:        store.Name,
		DisplayName: store.DisplayName,
		SeoDesc:     store.SeoDesc,
	}
	return s.repo.UptTag(c, tag)
}

func (s ServiceArticle) DeleteTag(c context.Context, tagId int) error {
	return s.repo.DelTag(c, &ArticleTags{Id: tagId})
}

func (s ServiceArticle) NewCate(c context.Context, store dto.CateStore) error {
	cate := NewArticleCategory(store.Name, store.DisplayName, store.SeoDesc)
	if store.ParentId != 0 {
		cate.SetParentId(store.ParentId)
	}
	return s.repo.NewCate(c, cate)
}

func (s ServiceArticle) UpdateCate(c context.Context, cateId int, store dto.CateStore) error {
	cate := &ArticleCategory{
		Id:          cateId,
		Name:        store.Name,
		DisplayName: store.DisplayName,
		SeoDesc:     store.SeoDesc,
		ParentId:    store.ParentId,
	}
	return s.repo.UptCate(c, cate)
}

func (s ServiceArticle) DeleteCate(c context.Context, cateId int) error {
	return s.repo.DelCate(c, &ArticleCategory{Id: cateId})
}
