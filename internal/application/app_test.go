package application_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-blog-ddd/config"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/application/query"
	"go-blog-ddd/internal/domain/commands"
	"math/rand/v2"
	"strings"
	"testing"
)

func init() {
	config.Init()
}

func TestNewApp(t *testing.T) {
	app := application.NewApp()
	ctx := context.Background()

	getcategorys := func() query.CategoryListView {
		categorys, err := app.Queries.Categorys.GetCategorys(ctx)
		if err != nil {
			panic(err)
		}
		return categorys
	}

	getpost := func(id uint32) query.PostView {
		byID, err := app.Queries.Posts.FindByID(ctx, id)
		if err != nil {
			panic(err)
		}
		return byID
	}

	gettags := func() query.TagListView {
		tags, err := app.Queries.Tags.GetTags(ctx)
		if err != nil {
			panic(err)
		}
		return tags
	}

	var cid, pid uint32
	var err error

	cid, err = createCategory(ctx, app) // 创建分类
	assert.NoError(t, err, "创建分类")

	assert.Equal(t, getcategorys().Count, 1)
	t.Log("new category", cid)

	err = createCategoryErr(ctx, app) // 尝试重复创建分类
	assert.NoError(t, err, "尝试重复创建分类")

	err = modifyCategoryDesc(ctx, app, cid) // 更新分类简介
	assert.NoError(t, err, "更新分类简介")

	pid, err = createPost(ctx, app)
	assert.NoError(t, err, "创建Post")

	t.Log("new post", pid)
	getpost(pid)

	err = createPostErr(ctx, app)
	assert.NoError(t, err, "尝试重复创建Post")

	err = modifyPostInfoCMD(ctx, app, pid)
	assert.NoError(t, err, "修改Post info")

	err = modifyPostVisibleCMD(ctx, app, pid)
	assert.NoError(t, err, "修改Post visible")

	err = modifyPostTagCMD(ctx, app, pid)
	assert.NoError(t, err, "修改Post tag")
	t.Log("post tags", getpost(pid).Tags)
	assert.Equal(t, gettags().Count, 4)

	err = resetPostCategoryCMD(ctx, app, pid, cid)
	assert.NoError(t, err, "重设PostCategory")
	t.Log("post category", getpost(pid).Category)

	err = resetPostContentCMD(ctx, app, pid)
	assert.NoError(t, err, "重设PostContent")

	err = deletePost(ctx, app, pid)
	assert.NoError(t, err, "删除Post")

	err = deleteCategory(ctx, app, cid)
	assert.NoError(t, err, "删除分类")

	assert.Equal(t, getcategorys().Count, 0)
	assert.Equal(t, gettags().Count, 0)
}

func createPost(ctx context.Context, app *application.App) (uint32, error) {
	var (
		id  uint32
		err error
	)
	if err = app.Transaction(ctx, func(tctx context.Context) error {
		id, err = app.Commands.CreatePost.Handle(tctx, commands.CreatePost{
			Uri:        "test-post",
			MDFilePath: "example.md",
		})

		return err
	}); err != nil {
		return 0, err
	}
	return id, nil
}

func createPostErr(ctx context.Context, app *application.App) error {
	var (
		err error
	)
	if err = app.Transaction(ctx, func(tctx context.Context) error {
		_, err = app.Commands.CreatePost.Handle(tctx, commands.CreatePost{
			Uri:        "test-post",
			MDFilePath: "example.md",
		})
		return err
	}); err != nil {
		return nil
	}
	return errors.New("重复创建Post没有报错")
}

func modifyPostInfoCMD(ctx context.Context, app *application.App, id uint32) error {
	var (
		newTitle = fmt.Sprintf("new title - %d", rand.Int())
		desc     = fmt.Sprintf("new desc - %d", rand.Int())
	)

	err := app.Commands.ModifyPostInfo.Handle(ctx, commands.ModifyPostInfo{
		ID:    id,
		Title: newTitle,
		Desc:  desc,
	})
	if err != nil {
		return err
	}
	find, err := app.Queries.Posts.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if find.Title != newTitle || find.Desc != desc {
		return errors.New("数据修改失败")
	}
	return nil
}

func modifyPostVisibleCMD(ctx context.Context, app *application.App, id uint32) error {
	err := app.Commands.ModifyPostVisible.Handle(ctx, commands.ModifyPostVisibility{
		ID:      id,
		Visible: true,
	})
	if err != nil {
		return err
	}
	find, err := app.Queries.Posts.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if !find.Visible {
		return errors.New("数据修改失败")
	}
	return nil
}
func modifyPostTagCMD(ctx context.Context, app *application.App, id uint32) error {
	var ss = strings.Split("adfasd,fwfj,iuasdfa,jsdfa", ",")
	err := app.Commands.ModifyPostTags.Handle(ctx, commands.ModifyPostTags{
		ID:      id,
		NewTags: ss,
	})
	if err != nil {
		return err
	}
	find, err := app.Queries.Posts.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if len(find.Tags) != len(ss) {
		return errors.New("数据修改失败")
	}
	return nil
}

func resetPostContentCMD(ctx context.Context, app *application.App, id uint32) error {
	find1, err := app.Queries.Posts.FindByID(ctx, id)

	err = app.Commands.ResetPostContent.Handle(ctx, commands.ResetPostContent{
		ID:         id,
		MDFilePath: "example2.md",
	})
	if err != nil {
		return err
	}
	find2, err := app.Queries.Posts.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if len(find1.Content) != len(find2.Content) {
		return nil
	}
	return errors.New("数据修改失败")
}
func resetPostCategoryCMD(ctx context.Context, app *application.App, id uint32, cid uint32) error {
	find1, err := app.Queries.Posts.FindByID(ctx, id)

	err = app.Commands.ResetPostCategory.Handle(ctx, commands.ResetPostCategory{
		ID:         id,
		CategoryID: cid,
	})

	if err != nil {
		return err
	}
	find2, err := app.Queries.Posts.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if find1.Category != find2.Category {
		return nil
	}
	return errors.New("数据修改失败")
}

func deletePost(ctx context.Context, app *application.App, id uint32) error {
	data, err := app.Queries.Posts.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err = app.Transaction(ctx, func(tctx context.Context) error {
		return app.Commands.DeletePost.Handle(tctx, commands.DeletePost{ID: data.ID})
	}); err != nil {
		return err
	}
	return nil
}

func createCategory(ctx context.Context, app *application.App) (uint32, error) {
	return app.Commands.CreateCategory.Handle(ctx, commands.CreateCategory{
		Name: "test-category",
		Desc: "test-category-desc",
	})
}

func createCategoryErr(ctx context.Context, app *application.App) error {
	_, err := app.Commands.CreateCategory.Handle(ctx, commands.CreateCategory{
		Name: "test-category",
		Desc: "test-category-desc",
	})
	if err != nil {
		return nil
	}
	return errors.New("重复创建Category没有报错")
}

func deleteCategory(ctx context.Context, app *application.App, id uint32) error {
	return app.Commands.DeleteCategory.Handle(ctx, commands.DeleteCategory{ID: id})
}

func modifyCategoryDesc(ctx context.Context, app *application.App, id uint32) error {
	desc := fmt.Sprintf("new desc - %d", rand.Int())

	err := app.Commands.ModifyCategoryDesc.Handle(ctx, commands.ModifyCategoryDesc{
		ID:      id,
		NewDesc: desc,
	})

	if err != nil {
		return err
	}
	find, err := app.Queries.Categorys.GetCategorys(ctx)
	if err != nil {
		return err
	}
	var ok = false
	for _, data := range find.Items {
		if data.Desc == desc {
			ok = true
		}
	}
	if ok {
		return nil
	}

	return errors.New("数据修改失败")

}
