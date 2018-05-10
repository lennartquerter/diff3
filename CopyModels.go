package main

import (
	"flag"
	"fmt"
	"os"
	"log"
)

func main() {
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

	path := "./"
	var models []Model

	if BoCommand.Parsed() {

		path = "/Users/lennartquerter/Documents/go/src/innovam.nl/Applications/src/"

		models = []Model{
			{
				src:  "Education/src/Innovam.Education.Api.Models",
				dest: "Backoffice/src/Innovam.Education.Api.Models",
			},
			{
				src:  "Email/src/Innovam.Email.Api.Models",
				dest: "Backoffice/src/Innovam.Email.Api.Models",
			},
		}
	}

	if PortCommand.Parsed() {
		path = "/Users/lennartquerter/Documents/go/src/lenimal.nl/Applications/src/"

		models = []Model{
			{
				src:  "Billing/src/Lenimal.Billing.Models",
				dest: "Portal/src/Lenimal.Billing.Models",
			},
			{
				src:  "Booking/src/Lenimal.Booking.Models",
				dest: "Portal/src/Lenimal.Booking.Models",
			},
			{
				src:  "Mail/src/Lenimal.Mail.Models",
				dest: "Portal/src/Lenimal.Mail.Models",
			},
		}
	}

	if TestCommand.Parsed() {
		path = "./Test/"

		models = []Model{
			{
				src:  "Dir1",
				dest: "Dir2",
			},
		}
	}

	err := MergeDir(path, models, true)

	if err != nil {
		log.Fatal(err)
	}
}
