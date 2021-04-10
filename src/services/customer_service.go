package services

import (
	"context"
	"database/sql"
	"log"

	v1 "github.com/kalpg69/backend-admin/src/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion                      = "v1"
	queryCreateUpdateDeleteCustomer = "CreateUpdateDeleteCustomer"
	createOperationType             = "I"
	updatedOperationType            = "U"
	deleteOperationType             = "D"
)

type customerServiceServer struct {
	v1.UnimplementedCustomerServiceServer
	db *sql.DB
}

func NewCustomerServiceServer(db *sql.DB) v1.CustomerServiceServer {
	return &customerServiceServer{db: db}
}

func (c *customerServiceServer) CreateCustomer(ctx context.Context, req *v1.CreateCustomerRequest) (*v1.CreateCustomerResponse, error) {

	if apiVersion != req.GetApi() {
		log.Printf("Error: API version for Create Customer request not implemented")
		return nil, status.Error(codes.Unimplemented, "Error: API version not supported")
	}

	//var returnValue, returnMessage string
	newCustomer := req.GetCustomer()

	_, err := c.db.ExecContext(ctx, queryCreateUpdateDeleteCustomer,
		sql.Named("CustomerCode", newCustomer.GetCustomerCode()),
		sql.Named("CustomerName", newCustomer.GetCustomerName()),
		sql.Named("CustomerAddress", newCustomer.GetCustomerAddress()),
		sql.Named("CustomerEmail", newCustomer.GetCustomerEmail()),
		sql.Named("CustomerPhone", newCustomer.GetCustomerPhone()),
		sql.Named("OperationType", createOperationType),
		sql.Named("LastUpdatedBy", newCustomer.GetLastUpdatedBy()))
	if err != nil {
		log.Printf("Error: Failed to create customer; %v", err)

		return nil, status.Errorf(codes.Internal, "Error: Failed to create customer")
	}

	return &v1.CreateCustomerResponse{
		Response: &v1.Response{
			Api:           "v1",
			Status:        "OK",
			StatusMessage: "Customer created successfully",
		},
	}, nil
}

func (c *customerServiceServer) UpdateCustomer(context.Context, *v1.UpdateCustomerRequest) (*v1.UpdateCustomerResponse, error) {
	return nil, nil
}

func (c *customerServiceServer) DeleteCustomer(context.Context, *v1.DeleteCustomerRequest) (*v1.DeleteCustomerResponse, error) {
	return nil, nil
}
