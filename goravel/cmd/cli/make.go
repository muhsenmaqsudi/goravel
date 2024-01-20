package main

import (
	"errors"
	"fmt"
	"time"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		dbType := gor.DB.DataType
		if arg3 == "" {
			exitGracefully(errors.New("you must give the migration a name"))
		}

		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)

		upFile := gor.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := gor.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

		err := copyFileFromTemplate("templates/migrations/migration." + dbType + ".up.sql", upFile)
		if err != nil {
			exitGracefully(err)
		}

		err = copyFileFromTemplate("templates/migrations/migration." + dbType + ".down.sql", downFile)
		if err != nil {
			exitGracefully(err)
		}
	}

	return nil
}