package presentation

import (
	"github.com/SpicyChickenFLY/lazysql/pkg/commands"
)

func GetDatabaseDisplayStrings(database *commands.Database) []string {
	return []string{
		database.Name,
		string(database.TableNum),
	}
}
