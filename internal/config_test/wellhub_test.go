package config_test

import (
	"testing"

	"github.com/pericles-luz/go-wellhub/internal/config"
	"github.com/stretchr/testify/require"
)

func TestWellhubMustImportConfigFile(t *testing.T) {
	configPath := "../../config/wellhub.json"
	configWellhub := config.NewWellhub(configPath)
	require.NotNil(t, configWellhub)
	require.NotEmpty(t, configWellhub.LinkAPI)
	require.NotEmpty(t, configWellhub.Token)
}
