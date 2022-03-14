package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_GetConfig(t *testing.T) {
	LoadTestEnv()
	conf := GetConfig()
	require.Equal(t, "testing", conf.Env)
}

func Test_GetRootDir(t *testing.T) {
	dir, err := GetRootDir()
	require.NoError(t, err)
	require.NotEmpty(t, dir)
}

func Test_UpdateTestConfig(t *testing.T) {
	LoadTestEnv()
	conf := GetConfig()
	UpdateTestConfig(func(config *AppConfig) {
		config.DebounceTime = time.Second * 10
	})
	conf = GetConfig()
	require.Equal(t, time.Second*10, conf.DebounceTime)
}

func Test_IsProd(t *testing.T) {
	LoadTestEnv()
	GetConfig()
	require.False(t, IsProd())
}

func Test_IsTestEnv(t *testing.T) {
	LoadTestEnv()
	GetConfig()
	require.True(t, IsTestEnv())
}
