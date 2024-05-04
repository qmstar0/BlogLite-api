package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/ports/cli/cmds"
	"go-blog-ddd/internal/ports/cli/util"
)

type Item struct {
	Title string
	Model tea.Model
}

func InitItems(app *application.App, root tea.Model) []Item {

	backhome := util.NewCountinue(lipgloss.NewStyle().
		Foreground(lipgloss.Color("239")).
		PaddingLeft(2).Render("Press any key to return to the home"))
	deleteCategoryHandler := cmds.NewDeleteCategoryModel(app)
	createCategory := cmds.NewCreateCategory(app)
	deletecategory := cmds.NewCategoryList(app)
	modifyCategroyDescHandler := cmds.NewModifyCategroyDesc(app)
	modifyCategroyDesc := cmds.NewCategoryList(app)

	backhome.With(nil, root)
	deleteCategoryHandler.With(deletecategory, backhome)
	createCategory.With(root, backhome)
	modifyCategroyDescHandler.With(modifyCategroyDesc, backhome)

	modifyCategroyDesc.With(root, modifyCategroyDescHandler)
	deletecategory.With(root, deleteCategoryHandler)
	return []Item{
		{"Create category", createCategory},
		{"Delete category", deletecategory},
		{"Modify cateogry description", modifyCategroyDesc},
	}
}
