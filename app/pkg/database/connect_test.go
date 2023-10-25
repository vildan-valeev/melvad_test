package database_test

import (
	"context"
	"github.com/vildan-valeev/melvad_test/pkg/database"
	"os"
	"testing"
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
