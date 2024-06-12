package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	City   string
	School string
	Sub    string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := &Config{
		City:   os.Getenv("CITY"),
		School: os.Getenv("SCHOOL"),
		Sub:    os.Getenv("SUB"),
	}

	flag.StringVar(&conf.City, "city", conf.City, "Debug mode [CITY]")

	flag.Parse()

	return conf
}
