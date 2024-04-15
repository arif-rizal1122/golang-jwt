package helper

import "log"

// IfLogError checks if an error exists and logs it as fatal.
func IfLogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
