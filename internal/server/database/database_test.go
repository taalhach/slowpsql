package database

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taalhach/slowpsql/internal/server/configs"
)

func TestMustConnectDB(t *testing.T) {
	t.Run("must db connect", func(t *testing.T) {
		cfg := configs.DatabaseConfig{
			Name:     "postgres",
			Host:     "localhost",
			Port:     5432,
			Password: "temp123",
		}
		err := MustConnectDB(&cfg)
		require.Nil(t, err)
	})

	t.Run("db connect must failed", func(t *testing.T) {
		cfg := configs.DatabaseConfig{
			Name:     "postgres",
			Host:     "ocalhost",
			Port:     5432,
			Password: "temp123",
		}
		err := MustConnectDB(&cfg)
		require.NotNil(t, err)
	})

}
