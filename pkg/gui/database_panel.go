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
					{
						Key:   "credits",
						Title: gui.Tr.CreditsTitle,
						Render: func(_database *commands.Database) tasks.TaskFunc {
							return gui.NewSimpleRenderStringTask(func() string { return "Hello lazysql" })
						},
					},
				}
			},
			GetItemContextCacheKey: func(database *commands.Database) string {
				return "databases-" + database.Name
			},
		},
		ListPanel: panels.ListPanel[*commands.Database]{
			List: panels.NewFilteredList[*commands.Database](),
			View: gui.Views.Databases,
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

func (gui *Gui) reloadDatabases() error {
	if err := gui.refreshStateDatabase(); err != nil {
		return err
	}

	return gui.Panels.Databases.RerenderList()
}

func (gui *Gui) refreshStateDatabase() error {
	databases, err := gui.SqlCommand.RefreshDatabases()
	if err != nil {
		return err
	}

	gui.Panels.Databases.SetItems(databases)

	return nil
}

func (gui *Gui) renderDatabaseInfo(_database *commands.Database) tasks.TaskFunc {
	return gui.NewSimpleRenderStringTask(func() string { return "Hello lazysql" })
}
