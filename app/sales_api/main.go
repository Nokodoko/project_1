package main

import (
	"log"
	"os"
)

func main() {
	//devs should only write to STDOUT
	log := log.New(os.Stdout, "SALES : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(log); err != nil {
		log.Prinln(err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error {

	return nil
}
