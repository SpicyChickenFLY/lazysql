package gui

import (
	"github.com/SpicyChickenFLY/lazysql/pkg/commands"
	"github.com/SpicyChickenFLY/lazysql/pkg/gui/panels"
	"github.com/SpicyChickenFLY/lazysql/pkg/gui/presentation"
	"github.com/SpicyChickenFLY/lazysql/pkg/tasks"
)

func (gui *Gui) getDatasourcePanel() *panels.SideListPanel[*commands.Datasource] {
	return &panels.SideListPanel[*commands.Datasource]{
		ContextState: &panels.ContextState[*commands.Datasource]{
			GetMainTabs: func() []panels.MainTab[*commands.Datasource] {
				return []panels.MainTab[*commands.Datasource]{
					{
						Key:   "credits",
						Title: gui.Tr.CreditsTitle,
						Render: func(_connection *commands.Datasource) tasks.TaskFunc {
							return gui.NewSimpleRenderStringTask(func() string { return "Hello lazysql" })
						},
					},
				}
			},
			GetItemContextCacheKey: func(connection *commands.Datasource) string {
				return "dsn-" + connection.Name
			},
		},
		ListPanel: panels.ListPanel[*commands.Datasource]{
			List: panels.NewFilteredList[*commands.Datasource](),
			View: gui.Views.Datasources,
		},
		NoItemsMessage: gui.Tr.NoImages,
		Gui:            gui.intoInterface(),

		Sort: func(a *commands.Datasource, b *commands.Datasource) bool {
			return false // dont sort for now
		},
		GetTableCells: presentation.GetDatasourceDisplayStrings,
		// It doesn't make sense to filter a list of only one item.
		DisableFilter: true,
	}
}

func (gui *Gui) reloadDatasources() error {
	if err := gui.refreshStateDatasource(); err != nil {
		return err
	}

	return gui.Panels.Datasources.RerenderList()
}

func (gui *Gui) refreshStateDatasource() error {
	databases, err := gui.SqlCommand.RefreshDatasources()
	if err != nil {
		return err
	}

	gui.Panels.Datasources.SetItems(databases)

	return nil
}

func (gui *Gui) renderDatasourceInfo(_database *commands.Datasource) tasks.TaskFunc {
	return gui.NewSimpleRenderStringTask(func() string { return "Hello lazysql" })
}
