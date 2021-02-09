package main

import (
	"os"
	"sync"

	nakamaContext "github.com/challenge-league/nakama-go/context"
	log "github.com/micro/go-micro/v2/logger"
)

var (
	nakamaCtx     *nakamaContext.Context
	nakamaCtxOnce sync.Once
)

func NewNakamaContextBase() {
	var err error
	nakamaCtx, err = nakamaContext.NewCustomAuthenticatedAdminAPIClient()
	if err != nil {
		log.Errorf("Error %+v", err)
		os.Exit(1)
	}
}

func NewNakamaContextSingleton() *nakamaContext.Context {
	nakamaCtxOnce.Do(func() {
		NewNakamaContextBase()
	})
	return nakamaCtx
}

func NewNakamaContext() *nakamaContext.Context {
	NewNakamaContextBase()
	return nakamaCtx
}

func init() {
	NewNakamaContextSingleton()
}
