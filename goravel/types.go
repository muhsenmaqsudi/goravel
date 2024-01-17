package goravel

import "database/sql"

type initPaths struct {
	rootPath    string
	folderNames []string
}

type databaseConfig struct {
	dsn string
	database string
}

type Database struct {
	DataType string
	Pool *sql.DB
}