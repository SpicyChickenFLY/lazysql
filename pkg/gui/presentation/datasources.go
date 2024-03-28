package presentation

import (
	"github.com/SpicyChickenFLY/lazysql/pkg/commands"
)

func GetDatasourceDisplayStrings(datasource *commands.Datasource) []string {
	return []string{
		datasource.Name,
	}
}
