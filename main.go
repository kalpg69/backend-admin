package main

import (
	"context"

	"github.com/kalpg69/backend-admin/src/app"
	"github.com/kalpg69/backend-admin/src/config"
)

func main() {
	ctx := context.Background()
	if err := app.StartApp(ctx, config.GetNetworkConfig(), config.GetSQLConfig()); err != nil {
		return
	}
}
