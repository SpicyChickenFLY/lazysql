package commands

import (
	"github.com/sirupsen/logrus"
)

// Table : A docker Table
type Table struct {
	Name          string
	ID            string
	OSCommand     *OSCommand
	Log           *logrus.Entry
	Container     *Container
	DockerCommand LimitedDockerCommand
	SqlCommand    SqlCommand
}

// Stop stops the service's containers
func (s *Table) Create() error {
	// return s.runCommand(s.OSCommand.Config.UserConfig.CommandTemplates.AddTable)
	_, err := s.SqlCommand.Execute("CREATE TABLE")
	return err
}

// Up up's the service
func (s *Table) Delete() error {
	// return s.runCommand(s.OSCommand.Config.UserConfig.CommandTemplates.DeleteTable)
	return nil
}

func (s *Table) runCommand(templateCmdStr string) error {
	// return s.OSCommand.RunCommand(command)
	return nil
}

//
// // ViewLogs attaches to a subprocess viewing the service's logs
// func (s *Table) ViewLogs() (*exec.Cmd, error) {
// 	templateString := s.OSCommand.Config.UserConfig.CommandTemplates.ViewTableLogs
// 	command := utils.ApplyTemplate(
// 		templateString,
// 		s.DockerCommand.NewCommandObject(CommandObject{Table: s}),
// 	)
//
// 	cmd := s.OSCommand.ExecutableFromString(command)
// 	s.OSCommand.PrepareForChildren(cmd)
//
// 	return cmd, nil
// }
//
// // RenderTop renders the process list of the service
// func (s *Table) RenderTop(ctx context.Context) (string, error) {
// 	templateString := s.OSCommand.Config.UserConfig.CommandTemplates.TableTop
// 	command := utils.ApplyTemplate(
// 		templateString,
// 		s.DockerCommand.NewCommandObject(CommandObject{Table: s}),
// 	)
//
// 	return s.OSCommand.RunCommandWithOutputContext(ctx, command)
// }
