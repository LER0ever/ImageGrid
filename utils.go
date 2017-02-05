package main

import (
	"log"
)

func ProcErr(msg string, err error) {
	if err != nil {
		log.Fatalf("Error %s: %s", msg, err.Error())
	}
}
