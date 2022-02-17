package main

import (
	flags "github.com/jessevdk/go-flags"
	"github.com/lucysuddenly/env-cli/internal/config"
	"github.com/lucysuddenly/env-cli/pkg/app"
	"os"
)

func main() {
	// defaults
	var cfg = config.Config{
		RemoteCluster: false,
	}

	_, err := flags.Parse(&cfg)
	if err != nil {
		// TODO: improve error feedback
		os.Exit(1)
	}

	err = app.Run(&cfg)
	if err != nil {
		os.Exit(1)
	}
}
