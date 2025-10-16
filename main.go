package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"profile_api/internals/handlers"
)

// Config holds the application configuration parameters.
type Config struct {
	Port     string
	LogLevel string
}

func resolveConfig(cliPort, cliLogLevel string) Config {
	defaultPort := "8080"
	defaultLogLevel := "info"

	cfg := Config{}

	// PORT RESOLUTION
	if cliPort != "" {
		cfg.Port = cliPort
	} else if os.Getenv("APP_PORT") != "" {
		cfg.Port = os.Getenv("APP_PORT")
	} else {
		cfg.Port = defaultPort
	}

	// LOG_LEVEL RESOLUTION
	if cliLogLevel != "" {
		cfg.LogLevel = cliLogLevel
	} else if os.Getenv("LOG_LEVEL") != "" {
		cfg.LogLevel = os.Getenv("LOG_LEVEL")
	} else {
		cfg.LogLevel = defaultLogLevel
	}

	return cfg
}

func setup() {
	var cliPort string
	var cliLogLevel string

	flag.StringVar(&cliPort, "port", "", "The port to listen on. Overrides APP_PORT environment variable.")
	flag.StringVar(&cliLogLevel, "log-level", "", "The log verbosity level (e.g., info, debug, warn). Overrides LOG_LEVEL environment variable.")

	flag.Parse()

	cfg := resolveConfig(cliPort, cliLogLevel)

	log.Printf("Configuration loaded: Port=%s, LogLevel=%s", cfg.Port, cfg.LogLevel)

	servemux := http.NewServeMux()
	servemux.HandleFunc("/me", handlers.ProfileHandler)

	serverAddress := ":" + cfg.Port
	log.Printf("Starting server on %s", serverAddress)

	err := http.ListenAndServe(serverAddress, servemux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
func main() {
	setup()
}
