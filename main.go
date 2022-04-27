package main

import (
	"context"
	"os"
	"os/signal"

	"gitea.com/gitea/act_runner/cmd"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// trap Ctrl+C and call cancel on the context
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	// run the command
	cmd.Execute(ctx)
}