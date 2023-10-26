package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		sig := <-sigs

		fmt.Fprintf(os.Stderr, "Shutting migration. Reason: %s...\n", sig.String())

		cancel()
	}()

	migrator, err := NewMigration(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// get the current migration status
	now, exp, info, err := migrator.Info(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if now < exp {
		// migration is required, dump out the current state
		// and perform the migration
		fmt.Fprintf(os.Stdout, "migration needed, current state: %v\n", info)

		err = migrator.Migrate(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		println("migration successful!")
	} else {
		fmt.Fprintln(os.Stdout, "no database migration needed")
		os.Exit(0)
	}
}
