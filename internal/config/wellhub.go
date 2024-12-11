package config

import "encoding/json"

type Wellhub struct {
	file    Base
	LinkAPI string `json:"linkAPI"`
	Token   string `json:"token"`
}

func (w *Wellhub) Load(file string) error {
	raw, err := w.file.ReadConfigurationFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(raw), w)
	if err != nil {
		return err
	}
	return nil
}

func (w *Wellhub) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"linkAPI": w.LinkAPI,
		"token":   w.Token,
	}
}

func NewWellhub(filePah string) *Wellhub {
	result := &Wellhub{}
	if err := result.Load(filePah); err != nil {
		return nil
	}
	return result
}
