package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	// run the migration command
	switch arg2 {
	case "up":
		err := gor.MigrateUp(dsn)
		if err != nil {
			return err
		}
	case "down":
		if arg3 == "all" {
			err := gor.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := gor.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
	case "reset":
		err := gor.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		err = gor.MigrateUp(dsn)
		if err != nil {
			return nil
		}

	default:
		showHelp()
	}

	return nil
}