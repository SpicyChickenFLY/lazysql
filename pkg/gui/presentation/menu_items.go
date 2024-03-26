package presentation

import "github.com/SpicyChickenFLY/lazysql/pkg/gui/types"

func GetMenuItemDisplayStrings(menuItem *types.MenuItem) []string {
	return menuItem.LabelColumns
}
