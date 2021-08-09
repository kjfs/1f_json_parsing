package config

import (
	"log"
)

// Config holds configuration data for this program.
type Config struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
