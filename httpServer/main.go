package main

import (
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	HTTPAddress string `env:"HTTP_ADDRESS" envDefault:":8080"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	LogFormat   string `env:"LOG_FORMAT" envDefault:"json"`
}

const (
	gracePeriod = 30 * time.Second
)

func main() {
	if err := r(); err != nil {
		log.Error().Err(err).Msg("")
		os.Exit(1)
	}
}

func r() error {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return err
	}
	log.Info().Interface("Config", cfg).Msg("")

	l, err := zerolog.ParseLevel(strings.ToLower(cfg.LogLevel))
	if err != nil {
		return err
	}

	zerolog.SetGlobalLevel(l)
	if strings.ToLower(cfg.LogFormat) == "console" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	s := New()

	return s.ListenAndServe(cfg.HTTPAddress)
}
