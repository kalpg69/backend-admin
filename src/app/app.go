package app

import (
	"context"
	"net"

	v1 "github.com/kalpg69/backend-admin/src/api/v1"
	"github.com/kalpg69/backend-admin/src/clients/sqlclients"
	"github.com/kalpg69/backend-admin/src/config"
	"github.com/kalpg69/backend-admin/src/services"
	"google.golang.org/grpc"
)

func StartApp(ctx context.Context, networkConfig *config.NetworkConfig, sqlConfig *config.SQLConfig) error {

	db, err := sqlclients.NewMSSQLClient(ctx, sqlConfig.UserName, sqlConfig.Password, sqlConfig.ServerName, sqlConfig.DatabaseName, sqlConfig.Timeout)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	v1.RegisterCustomerServiceServer(grpcServer, services.NewCustomerServiceServer(db))

	lis, err := net.Listen("tcp", networkConfig.GRPCServerAddress+":"+networkConfig.GRPCPort)
	if err != nil {
		return err
	}

	grpcServer.Serve(lis)
	return nil

}
