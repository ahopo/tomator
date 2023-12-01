package helper

import (
	"log"
)

func CheckError(e error, isfatal bool) {
	if isfatal && e != nil {
		log.Fatal(e)
	} else if !isfatal && e != nil {
		log.Println(e)
	}
}
