package main

import (
	"context"
	"sync"

	"github.com/hpaes/go-api-final-project/src/infra/setup"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	setup.
		NewSetup().
		InitLogger().
		WithAppConfig().
		WithDatabase().
		WithRouter().
		WithServer().
		Run(ctx, &wg)

	wg.Wait()
}
