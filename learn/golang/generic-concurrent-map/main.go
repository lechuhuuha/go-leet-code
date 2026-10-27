package main

import (
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// useMapErr(ctx, cancel)
	// useBatch(ctx)
	// useBatchWithInterval(ctx)
	useSharedSemaphore(ctx)
}
