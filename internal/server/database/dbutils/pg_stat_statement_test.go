package dbutils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taalhach/slowpsql/internal/server/configs"
	"github.com/taalhach/slowpsql/internal/server/database"
	"github.com/taalhach/slowpsql/pkg/forms"
)

func TestFindStatements(t *testing.T) {
	form := forms.BasicList{
		Limit: 10,
		Page:  1,
	}
	t.Run("connect db", func(t *testing.T) {
		cfg := configs.DatabaseConfig{
			Name:     "postgres",
			Host:     "localhost",
			Port:     5432,
			Password: "temp123",
		}
		err := database.MustConnectDB(&cfg)
		require.Nil(t, err)
	})

	t.Run("Page limit test", func(t *testing.T) {
		items, total, err := FindStatements(&form)
		require.Nil(t, err)
		require.GreaterOrEqual(t, int(total), int(0))
		require.NotNil(t, items)
	})

	t.Run("Select Filter test", func(t *testing.T) {
		form.Filters = []string{"query:starts_with:select"}
		items, total, err := FindStatements(&form)
		require.Nil(t, err)
		require.GreaterOrEqual(t, int(total), int(0))
		require.NotNil(t, items)
		require.Greater(t, len(items), 0)
		for _, item := range items {
			// check if query start with update
			has := strings.HasPrefix(strings.ToLower(item.Query), "update")
			// should be false because we have passed start_with select filter
			require.False(t, has)
		}
	})

	t.Run("Create Filter test", func(t *testing.T) {
		form.Filters = []string{"query:starts_with:create"}
		items, total, err := FindStatements(&form)
		require.Nil(t, err)
		require.GreaterOrEqual(t, int(total), int(0))
		require.NotNil(t, items)
		require.Greater(t, len(items), 0)
		for _, item := range items {
			// check if query start with update
			has := strings.HasPrefix(strings.ToLower(item.Query), "update")
			// should be false because we have passed start_with select filter
			require.False(t, has)
		}
	})
}
