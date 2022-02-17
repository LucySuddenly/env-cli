package app

import (
	"fmt"
	"github.com/lucysuddenly/env-cli/internal/config"
)

func Run(cfg *config.Config) error {
	fmt.Printf("%+v\n", cfg)
	return nil
}

func Shutdown() {
}
