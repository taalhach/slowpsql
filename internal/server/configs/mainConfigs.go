package configs

import (
	ini "github.com/nanitor/goini"
)

type MainConfig struct {
	Database *DatabaseConfig
}

func LoadMainConfig(dict ini.Dict) (*MainConfig, error) {
	var err error

	ret := &MainConfig{}
	ret.Database, err = DatabaseConfigsFromDict(dict)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
