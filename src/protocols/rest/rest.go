package rest

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	v1 "github.com/kalpg69/backend-admin/src/api/v1"
	"github.com/kalpg69/backend-admin/src/config"
	"google.golang.org/grpc"
)

func StartRestServer(ctx context.Context, networkConfig *config.NetworkConfig) error {

	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := v1.RegisterCustomerServiceHandlerFromEndpoint(ctx, mux, networkConfig.GRPCServerAddress+":"+networkConfig.GRPCPort, opts); err != nil {
		log.Fatalf("Error: Failed to start HTTP gateway: %v", err)
	}

	if err := v1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, networkConfig.GRPCServerAddress+":"+networkConfig.GRPCPort, opts); err != nil {
		log.Fatalf("Error: Failed to start HTTP gateway: %v", err)
	}

	server := &http.Server{
		Addr:    networkConfig.HTTPServerAddress + ":" + networkConfig.HTTPPort,
		Handler: mux,
	}

	return server.ListenAndServe()

}
