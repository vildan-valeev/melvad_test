package database_test

import (
	"context"
	"os"
	"testing"

	"github.com/vildan-valeev/melvad_test/pkg/database"
)

func TestNewPooll(t *testing.T) {
	ctx := context.Background()
	dataSourceName := os.Getenv("TN_TEST_DSN")
	opts := []database.Option{
		database.WithLogLevel("debug"),
	}

	got, err := database.NewPool(ctx, dataSourceName, opts...)
	if err != nil {
		t.Error(err)
		return
	}

	if got == nil {
		t.Error("NewPool() return nil")
	}
}
