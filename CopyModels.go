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

	if len(os.Args) < 2 {
		fmt.Println("bo or portal subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "bo":
		BoCommand.Parse(os.Args[2:])
	case "service":
		PortCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	path := "./"
	models := []Model{
		{
			src:  "./output/new_version.txt",
			dest: "./output/old_version.txt",
			past: "./output/past_version.txt",
		},
	}

	//if BoCommand.Parsed() {
	//
	//	path = "/Users/lennartquerter/Documents/go/src/innovam.nl/Applications/src/"
	//
	//	models = []Model{
	//		{
	//			src:  "Education/src/Innovam.Education.Api.Models",
	//			dest: "Backoffice/src/Innovam.Education.Api.Models",
	//		},
	//		{
	//			src:  "Email/src/Innovam.Email.Api.Models",
	//			dest: "Backoffice/src/Innovam.Email.Api.Models",
	//		},
	//	}
	//}
	//
	//if PortCommand.Parsed() {
	//	path = "/Users/lennartquerter/Documents/go/src/lenimal.nl/Applications/src/"
	//
	//	models = []Model{
	//		{
	//			src:  "Billing/src/Lenimal.Billing.Models",
	//			dest: "Portal/src/Lenimal.Billing.Models",
	//		},
	//		{
	//			src:  "Booking/src/Lenimal.Booking.Models",
	//			dest: "Portal/src/Lenimal.Booking.Models",
	//		},
	//		{
	//			src:  "Mail/src/Lenimal.Mail.Models",
	//			dest: "Portal/src/Lenimal.Mail.Models",
	//		},
	//	}
	//}

	err := MergeFiles(path, models, true)

	if err != nil {
		log.Fatal(err)
	}
}
