package gui

import (
	"github.com/SpicyChickenFLY/lazysql/pkg/commands"
	"github.com/SpicyChickenFLY/lazysql/pkg/gui/panels"
	"github.com/SpicyChickenFLY/lazysql/pkg/gui/presentation"
	"github.com/SpicyChickenFLY/lazysql/pkg/tasks"
)

// Although at the moment we'll only have one project, in future we could have
// a list of projects in the project panel.

func (gui *Gui) getDatabasePanel() *panels.SideListPanel[*commands.Database] {
	return &panels.SideListPanel[*commands.Database]{
		ContextState: &panels.ContextState[*commands.Database]{
			GetMainTabs: func() []panels.MainTab[*commands.Database] {
				return []panels.MainTab[*commands.Database]{
                    // TODO:
					// {
					// 	Key:    "config",
					// 	Title:  gui.Tr.ConfigTitle,
					// 	Render: gui.renderDatabaseConfigTask,
					// },
				}
			},
			GetItemContextCacheKey: func(database *commands.Database) string {
				return "images-" + database.Name
			},
		},
		ListPanel: panels.ListPanel[*commands.Database]{
			List: panels.NewFilteredList[*commands.Database](),
			View: gui.Views.Images,
		},
		NoItemsMessage: gui.Tr.NoImages,
		Gui:            gui.intoInterface(),

		Sort: func(a *commands.Database, b *commands.Database) bool {
			return false // dont sort for now
		},
        GetTableCells: presentation.GetDatabaseDisplayStrings,
		// It doesn't make sense to filter a list of only one item.
		DisableFilter: true,
	}
}

func (gui *Gui) renderDatabaseInfo(_project *commands.Project) tasks.TaskFunc {
	return gui.NewSimpleRenderStringTask(func() string { return "Hello lazysql" })
}
