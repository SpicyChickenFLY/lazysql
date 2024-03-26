package presentation

import "github.com/SpicyChickenFLY/lazysql/pkg/commands"

func GetProjectDisplayStrings(project *commands.Project) []string {
	return []string{project.Name}
}
