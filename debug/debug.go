package debug

import (
	"log"

	"github.com/Scalingo/acadock-monitoring/config"
)

func Println(args ...interface{}) {
	if config.Debug {
		log.Println(args...)
	}
}

func Printf(format string, args ...interface{}) {
	if config.Debug {
		log.Printf(format, args...)
	}
}
