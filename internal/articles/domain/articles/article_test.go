package articles_test

import (
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArticleVersion(t *testing.T) {
	t.Parallel()
	var version = articles.Version{
		Title:       "title",
		Description: "description",
		Text:        "text",
		Hash:        "a94a8fe5ccb19ba61c4c0873d391e987982fbbd",
		Source:      "source",
		Note:        "note",
	}
	var version2 = articles.Version{
		Title:       "title",
		Description: "description",
		Text:        "text",
		Hash:        "a94a8fe5ccb19ba61124uh1237h123182ged123",
		Source:      "source",
		Note:        "note",
	}
	article := articles.NewArticle(articles.NewUri("test"), "categoryID")
	assert.Equal(t, article.HasCurrentVersion(), false)
	assert.Equal(t, 1, len(article.Events())) //创建事件

	assert.Equal(t, article.CurrentVersion(), "")
	assert.Error(t, article.SetCurrentVersion(version.Hash))
	assert.Equal(t, 1, len(article.Events())) // 创建

	assert.Error(t, article.RemoveVersion(version.Hash))

	assert.NoError(t, article.AddNewVresion(version))
	assert.Equal(t, 3, len(article.Events()))       // 创建、新增版本、设置版本(没有版本自动设置)
	assert.Error(t, article.AddNewVresion(version)) // 添加已有的版本应该报错

	assert.NoError(t, article.AddNewVresion(version2))
	assert.Equal(t, 4, len(article.Events())) // 创建、新增版本、设置版本(没有版本自动设置), 新增版本

	assert.Error(t, article.SetCurrentVersion(""))             // 设置不存在的版本，应该报错
	assert.NoError(t, article.SetCurrentVersion(version.Hash)) // 设置版本和当前版本一样，应当直接返回，不产生事件
	assert.Equal(t, article.HasCurrentVersion(), true)
	assert.Equal(t, article.CurrentVersion(), version.Hash)
	assert.Equal(t, 4, len(article.Events()))                   // 创建、新增版本、设置版本(没有版本自动设置)，新增版本
	assert.NoError(t, article.SetCurrentVersion(version2.Hash)) // 设置新版本，产生事件
	assert.Equal(t, 5, len(article.Events()))                   // 创建、新增版本、设置版本(没有版本自动设置)，新增版本、设置版本

	article.ChangeVisibility(true)
	assert.Equal(t, 6, len(article.Events())) // 创建、新增版本、设置版本(没有版本自动设置)、新增版本、设置版本、修改可见性

	assert.Error(t, article.RemoveVersion(""))             // 移除没有的版本应该报错
	assert.Error(t, article.RemoveVersion(version2.Hash))  // 移除正在使用的版本应该报错
	assert.NoError(t, article.RemoveVersion(version.Hash)) // 移除未正在使用的版本
	assert.Equal(t, 7, len(article.Events()))              // 创建、新增版本、设置版本(没有版本自动设置)、新增版本、设置版本、修改可见性、删除版本、

	var err error
	group, err := articles.NewTagGroup([]string{"1", "2", "3", "4"})
	assert.NoError(t, err)
	article.ChangeTagGroup(group)
	assert.Equal(t, article.TagGroup(), group)
	assert.Equal(t, 8, len(article.Events())) // 创建、新增版本、设置版本(没有版本自动设置)、新增版本、设置版本、修改可见性、删除版本，修改标签

	article.ChangeCategory("new-categoryID")
	assert.Equal(t, 9, len(article.Events())) // 创建、新增版本、设置版本(没有版本自动设置)、新增版本、设置版本、修改可见性、删除版本，修改标签、更改分类

	article.Delete()
	assert.Equal(t, 10, len(article.Events())) // 创建、新增版本、设置版本(没有版本自动设置)、新增版本、设置版本、修改可见性、删除版本，修改标签、更改分类、删除文章
}
