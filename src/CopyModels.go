package src

import (
	"flag"
	"fmt"
	"os"
	"log"
	. "CopyModels/src/exported"
	"os/user"
)

func CopyModels() {
	BoCommand := flag.NewFlagSet("bo", flag.ExitOnError)
	PortCommand := flag.NewFlagSet("portal", flag.ExitOnError)
	TestCommand := flag.NewFlagSet("test", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("bo, text or portal subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "bo":
		BoCommand.Parse(os.Args[2:])
	case "service":
		PortCommand.Parse(os.Args[2:])
	case "test":
		TestCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := dir
	var models []Dir

	if BoCommand.Parsed() {

		path = "/Users/lennartquerter/Documents/go/src/innovam.nl/Applications/src/"

		models = []Dir{
			{
				Source:      "Education/src/Innovam.Education.Api.exported",
				Destination: "Backoffice/src/Innovam.Education.Api.exported",
			},
			{
				Source:      "Email/src/Innovam.Email.Api.exported",
				Destination: "Backoffice/src/Innovam.Email.Api.exported",
			},
		}
	}

	if PortCommand.Parsed() {
		path = "/Users/lennartquerter/Documents/go/src/lenimal.nl/Applications/src/"

		models = []Dir{
			{
				Source:      "Billing/src/Lenimal.Billing.exported",
				Destination: "Portal/src/Lenimal.Billing.exported",
			},
			{
				Source:      "Booking/src/Lenimal.Booking.exported",
				Destination: "Portal/src/Lenimal.Booking.exported",
			},
			{
				Source:      "Mail/src/Lenimal.Mail.exported",
				Destination: "Portal/src/Lenimal.Mail.exported",
			},
		}
	}

	if TestCommand.Parsed() {

		models = []Dir{
			{
				Source:      "/Test/src",
				Destination: "/Test/dest",
			},
		}
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	historyPath := usr.HomeDir + "/.go/copy_models"

	err = MergeDir(path, models, historyPath, false)

	if err != nil {
		log.Fatal(err)
	}
}
