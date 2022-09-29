package main

import (
	"github.com/strwrd/driver-aood/src/business"
	"github.com/strwrd/driver-aood/src/cmd"
	"github.com/strwrd/driver-aood/src/storage"
	"log"
	"os"
)

func main() {
	// verify that command arguments are correctly provided, discard first argument
	if len(os.Args) != 3 {
		log.Fatalln("invalid number of arguments, expected 2 arguments")
	}

	// initialize storage to store current state of program
	storage, err := storage.NewStorage()
	if err != nil {
		log.Fatalln(err)
	}

	ucase := business.NewCase(storage)
	command := cmd.NewCommand(ucase)

	// dispatch proper command
	switch os.Args[1] {
	case "register_driver":
		output, err := command.RegisterDriver(os.Args[2])
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(output)
	case "dispatch_driver_for_a_booking":
		output, err := command.DispatchDriver(os.Args[2])
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(output)
	case "complete_booking":
		output, err := command.CompleteBooking(os.Args[2])
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(output)
	case "drivers_completed_booking_gt":
		output, err := command.CompletedBookingGT(os.Args[2])
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(output)
	default:
		log.Fatalf("invalid `%s` command", os.Args[1])
	}

	// commit current state program as json to file
	storage.Save()
}
