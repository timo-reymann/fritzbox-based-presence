package log

import (
	"log"
)

// Print a log message
func Print(component string, message string) {
	log.Printf("[%s] %s\n", component, message)
}
