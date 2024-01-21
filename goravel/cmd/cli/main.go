package main

import (
	"errors"
	"os"

	"github.com/fatih/color"
	"github.com/muhsenmaqsudi/goravel"
)

const version = "1.0.0"

var gor goravel.Goravel

func main() {
	var message string
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	setup()

	switch arg1 {
	case "help":
		showHelp()
	case "version":
		color.Yellow("Application version: " + version)

	case "migrate":
		if arg2 == "" {
			arg2 = "up"
		}
		err = doMigrate(arg2, arg3)
		if err != nil {
			exitGracefully(err)
		}
		message = "Migrations complete"
	case "make":
		if arg2 == "" {
			exitGracefully(errors.New("make requires a subcommand: (migration|model|handler)"))
		}
		err = doMake(arg2, arg3)
		if err != nil {
			exitGracefully(err)
		}
	default:
		showHelp()
	}

	exitGracefully(nil, message)
}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string

	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}

		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		color.Red("Error: command required")
		showHelp()
		return "","","", errors.New("command required")
	}

	return arg1, arg2, arg3, nil
}

func showHelp() {
	color.Yellow(`Available commands:
	help				  				- show the help commands
	version			  				-	print application version
	migrate			  				- runs all up migrations that have not been run previously
	migrate down  				- reverses the most recent migration
	migrate reset 				- runs all down migrations in reverse order, and then all up migrations
	make migration <name> - create two new up and down migrations in the migrations folder
	make handler <name>		- creates a stub handler in the handlers directory
	make model <name>			- creates a new model in the models directory
	`)
}

func exitGracefully(err error, msg ...string) {
	message := ""
	if len(msg) > 0{
		message = msg[0]
	}

	if err != nil {
		color.Red("Error: %v\n", err)
	}

	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished")
	}

	os.Exit(0)
}