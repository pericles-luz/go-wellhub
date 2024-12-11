package config

import (
	"os"
)

type Base struct {
	raw []byte
}

func (cfg *Base) ReadConfigurationFile(fileName string) ([]byte, error) {
	if _, err := os.Stat(fileName); nil != err {
		return nil, err
	}
	json, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	cfg.raw = json
	return cfg.raw, nil
}
