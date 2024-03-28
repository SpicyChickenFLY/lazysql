package commands

import (
	"github.com/google/uuid"

	"github.com/SpicyChickenFLY/lazysql/pkg/config"
	"github.com/sirupsen/logrus"
)

type SqlCommand interface {
	Connect(urlstr string) error
	TestConnection(urlstr string) error
	GetDatabases() ([]string, error)
	// CreateTable(database string, table string) ()
	GetTables(database string) (map[string][]string, error)
	GetTableColumns(database, table string) ([][]string, error)
	// UpdateTable(database string, table string) error
	// DeleteTable(database string, table string) error
	GetConstraints(table string) ([][]string, error)
	GetForeignKeys(table string) ([][]string, error)
	GetIndexes(table string) ([][]string, error)
	GetRecords(table, where, sort string, offset, limit int) ([][]string, int, error)
	UpdateRecord(table, column, value, primaryKeyColumnName, primaryKeyValue string) error
	DeleteRecord(table string, primaryKeyColumnName, primaryKeyValue string) error
	Execute(query string) (string, error)
	Query(query string) ([][]string, error)
	ExecutePendingChanges(changes []DbDmlChange, inserts []DbInsert) error
	SetProvider(provider string)
	GetProvider() string
}

// NewOSCommand os command runner
func NewSqlCommand(log *logrus.Entry, config *config.AppConfig) SqlCommand {
	return &MySQLCommand{}
}

type DbDmlChange struct {
	Type                 string
	Table                string
	Column               string
	Value                string
	PrimaryKeyColumnName string
	PrimaryKeyValue      string
	Option               int
}

type DbInsert struct {
	Table           string
	Columns         []string
	Values          []string
	Option          int
	PrimaryKeyValue uuid.UUID
}

type DbTableColumn struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}
