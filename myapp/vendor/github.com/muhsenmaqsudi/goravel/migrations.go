package goravel

import (
	"log"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (c *Goravel) MigrateUp(dsn string) error {
	m, err := migrate.New("file://" + c.RootPath + "/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil {
		log.Println("Error running migration")
		return err
	}
	return nil
}

func (c *Goravel) MigrateDownAll(dsn string) error {
	m, err := migrate.New("file://" + c.RootPath + "/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Down(); err != nil {
		return err
	}

	return nil
}

func (c *Goravel) Steps(n int, dsn string) error {
	m, err := migrate.New("file://" + c.RootPath + "/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Steps(n); err != nil {
		return err
	}

	return nil
}

func (c *Goravel) MigrateForce(dsn string) error {
	m, err := migrate.New("file://" + c.RootPath + "/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Force(-1); err != nil {
		return err
	}

	return nil
}