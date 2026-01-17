package events

import "log"

// Generic error handling
func CheckError(e error) {
	if e != nil {
		log.Printf("ERROR: %v", e)
		panic(e)
	}
}
