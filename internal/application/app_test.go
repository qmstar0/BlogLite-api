package application_test

//TODO 需要重写 App test
//var app *application.App
//var ctx = context.Background()
//
//func init() {
//	config.Init()
//	closeFn := postgresql.Init()
//	utils.OnShutdown(closeFn)
//
//	app = application.NewApp()
//
//}
//
//func TestNewApp(t *testing.T) {
//
//	var (
//		newcategory = fmt.Sprintf("category - %d", rand.Int())
//		newpost     = fmt.Sprintf("posts - %d", rand.Int())
//	)
//
//	var cid, pid uint32
//	var err error
//
//	cid, err = createCategory(newcategory) // 创建分类
//	assert.NoError(t, err, "创建分类")
//	t.Log("new category", newcategory, cid)
//
//	err = createCategoryErr(newcategory) // 尝试重复创建分类
//	assert.NoError(t, err, "尝试重复创建分类")
//
//	err = modifyCategoryDesc(cid) // 更新分类简介
//	assert.NoError(t, err, "更新分类简介")
//
//	pid, err = createPost(newpost)
//	assert.NoError(t, err, "创建Post")
//	t.Log("new post", newpost, pid)
//
//	err = createPostErr(newpost)
//	assert.NoError(t, err, "尝试重复创建Post")
//
//	err = modifyPostHeadCMD(pid)
//	assert.NoError(t, err, "修改Post head")
//
//	err = modifyPostDescCMD(pid)
//	assert.NoError(t, err, "修改Post desc")
//
//	err = modifyPostVisibleCMD(pid)
//	assert.NoError(t, err, "修改Post visible")
//
//	err = modifyPostTagCMD(pid)
//	assert.NoError(t, err, "修改Post tag")
//
//	err = resetPostCategoryCMD(pid, cid)
//	assert.NoError(t, err, "重设PostCategory")
//
//	err = resetPostContentCMD(pid)
//	assert.NoError(t, err, "重设PostContent")
//
//	err = deletePost(pid)
//	assert.NoError(t, err, "删除Post")
//
//	_, err = app.Queries.Posts.FindByID(ctx, pid)
//	assert.Error(t, err, "post已删除，应当查询失败")
//
//	err = deleteCategory(cid)
//	assert.NoError(t, err, "删除分类")
//
//	result, err := app.Queries.Categorys.All(ctx)
//	for _, item := range result.Items {
//		assert.NotEqual(t, item.ID, cid, "category 已删除，不应该出现")
//	}
//}
//
//func createPost(uri string) (uint32, error) {
//	var (
//		id  uint32
//		err error
//	)
//	if err = app.Transaction(ctx, func(tctx context.Context) error {
//		id, err = app.Commands.CreatePost.Handle(tctx, commands.CreatePost{
//			Uri:        uri,
//			MDFilePath: "example.md",
//		})
//
//		return err
//	}); err != nil {
//		return 0, err
//	}
//	return id, nil
//}
//
//func createPostErr(name string) error {
//	var (
//		err error
//	)
//	if err = app.Transaction(ctx, func(tctx context.Context) error {
//		_, err = app.Commands.CreatePost.Handle(tctx, commands.CreatePost{
//			Uri:        name,
//			MDFilePath: "example.md",
//		})
//		return err
//	}); err != nil {
//		return nil
//	}
//	return errors.New("重复创建Post没有报错")
//}
//
//func modifyPostHeadCMD(id uint32) error {
//	var (
//		newTitle = fmt.Sprintf("new title - %d", rand.Int())
//		uri      = fmt.Sprintf("new uri - %d", rand.Int())
//	)
//
//	err := app.Commands.ModifyPostTitle.Handle(ctx, commands.ModifyPostHead{
//		ID:    id,
//		Title: newTitle,
//		Uri:   uri,
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func modifyPostDescCMD(id uint32) error {
//	var (
//		desc = fmt.Sprintf("new desc - %d", rand.Int())
//	)
//
//	err := app.Commands.ModifyPost.Handle(ctx, commands.ModifyPostDesc{
//		ID:   id,
//		Desc: desc,
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func modifyPostVisibleCMD(id uint32) error {
//	err := app.Commands.ModifyPostVisible.Handle(ctx, commands.ModifyPostVisibility{
//		ID:      id,
//		Visible: true,
//	})
//	if err != nil {
//		return err
//	}
//	_, err = app.Queries.Posts.FindByID(ctx, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func modifyPostTagCMD(id uint32) error {
//	var ss = strings.Split("adfasd,fwfj,iuasdfa,jsdfa", ",")
//	err := app.Commands.ModifyPostTags.Handle(ctx, commands.ModifyPostTags{
//		ID:      id,
//		NewTags: ss,
//	})
//	if err != nil {
//		return err
//	}
//	find, err := app.Queries.Posts.FindByID(ctx, id)
//	if err != nil {
//		return err
//	}
//	if len(find.Tags) != len(ss) {
//		return errors.New("数据修改失败")
//	}
//	return nil
//}
//
//func resetPostContentCMD(id uint32) error {
//	find1, err := app.Queries.Posts.FindByID(ctx, id)
//
//	err = app.Commands.ResetPostContent.Handle(ctx, commands.ResetPostContent{
//		ID:         id,
//		MDFilePath: "example2.md",
//	})
//	if err != nil {
//		return err
//	}
//	find2, err := app.Queries.Posts.FindByID(ctx, id)
//	if err != nil {
//		return err
//	}
//	if len(find1.Content) != len(find2.Content) {
//		return nil
//	}
//	return errors.New("数据修改失败")
//}
//func resetPostCategoryCMD(id uint32, cid uint32) error {
//	find1, err := app.Queries.Posts.FindByID(ctx, id)
//
//	err = app.Commands.ResetPostCategory.Handle(ctx, commands.ResetPostCategory{
//		ID:         id,
//		CategoryID: cid,
//	})
//
//	if err != nil {
//		return err
//	}
//	find2, err := app.Queries.Posts.FindByID(ctx, id)
//	if err != nil {
//		return err
//	}
//	if find1.Category != find2.Category {
//		return nil
//	}
//	return errors.New("数据修改失败")
//}
//
//func deletePost(id uint32) error {
//	data, err := app.Queries.Posts.FindByID(ctx, id)
//	if err != nil {
//		return err
//	}
//
//	if err = app.Transaction(ctx, func(tctx context.Context) error {
//		return app.Commands.DeletePost.Handle(tctx, commands.DeletePost{ID: data.ID})
//	}); err != nil {
//		return err
//	}
//	return nil
//}
//
//func createCategory(name string) (uint32, error) {
//	return app.Commands.CreateCategory.Handle(ctx, commands.CreateCategory{
//		Name: name,
//		Desc: "test-category-desc",
//	})
//}
//
//func createCategoryErr(name string) error {
//	_, err := app.Commands.CreateCategory.Handle(ctx, commands.CreateCategory{
//		Name: name,
//		Desc: "test-category-desc",
//	})
//	if err != nil {
//		return nil
//	}
//	return errors.New("重复创建Category没有报错")
//}
//
//func deleteCategory(id uint32) error {
//	return app.Commands.DeleteCategory.Handle(ctx, commands.DeleteCategory{ID: id})
//}
//
//func modifyCategoryDesc(id uint32) error {
//	desc := fmt.Sprintf("new desc - %d", rand.Int())
//
//	err := app.Commands.ModifyCategoryDesc.Handle(ctx, commands.ModifyCategoryDesc{
//		ID:      id,
//		NewDesc: desc,
//	})
//
//	if err != nil {
//		return err
//	}
//	find, err := app.Queries.Categorys.All(ctx)
//	if err != nil {
//		return err
//	}
//	var ok = false
//	for _, data := range find.Items {
//		if data.Desc == desc {
//			ok = true
//		}
//	}
//	if ok {
//		return nil
//	}
//
//	return errors.New("数据修改失败")
//
//}
