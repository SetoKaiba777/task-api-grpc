package main

import (
	"context"
	"go-challenger/infrastructure/loader"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()


	var wg sync.WaitGroup
	
	loader.
		NewConfig().
		WithAppConfig().
		InitLogger().
		WithDB().
		WithRepository().
		WithWebServer().
		Start(ctx,&wg)

		wg.Wait()
}
