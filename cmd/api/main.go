package main

import (
	"context"
	"go-rest-webserver-template/internal/boot"
	"go-rest-webserver-template/internal/helpers"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(boot.NewContext(nil))
	defer cancel()

	environ, err := helpers.GetEnv()
	if err != nil {
		log.Fatalf("failed to init api: %v", err)
	}

	if err = boot.InitAPI(ctx, environ); err != nil {
		log.Fatal("application boot failed", "error", err)
	}

}
