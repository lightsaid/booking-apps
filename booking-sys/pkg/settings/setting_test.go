package settings

import (
	"testing"
	"toolkit/configs"

	"github.com/stretchr/testify/require"
)

func TestReadConfig(t *testing.T) {
	var appConfig AppConfig
	_, err := configs.NewConfig("config.yaml", &appConfig, "../../configs")
	require.NoError(t, err)
	require.Equal(t, "postgres", appConfig.Database.DBDriver)
	require.True(t, true, len(appConfig.DBSource) > 10)
	require.Equal(t, 5000, appConfig.Server.Port)
	require.Equal(t, "booking-apps#Author@lightsaid", appConfig.JWT.Secret)
}
