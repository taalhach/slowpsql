package server

import (
	"fmt"
	"os"

	ini "github.com/nanitor/goini"
	"github.com/taalhach/slowpsql/internal/server/configs"
	"github.com/taalhach/slowpsql/internal/server/database"
)

const (
	envKey = "SETTINGS"
)

var (
	MainConfigs *configs.MainConfig
)

func loadConfigs() (*configs.MainConfig, error) {
	file := os.Getenv(envKey)
	if file == "" {
		fmt.Printf("Missing env variable: %v", envKey)
		os.Exit(1)
	}

	dict, err := ini.Load(file)
	if err != nil {
		return nil, err
	}

	MainConfigs, err = configs.LoadMainConfig(dict)
	if err != nil {
		return nil, err
	}

	// must make database connection or panic
	if err := database.MustConnectDB(MainConfigs.Database); err != nil {
		return nil, err
	}

	return MainConfigs, err
}
