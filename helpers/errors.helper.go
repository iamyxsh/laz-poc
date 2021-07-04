package helpers

import "log"

func HandleErr(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}
