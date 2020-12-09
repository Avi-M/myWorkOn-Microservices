package impl

import (
	"demo_grpc/domain"

	"context"
	"log"
	"strconv"
)

//RepositoryServiceGrpcImpl is a implementation of RepositoryService Grpc Service.
type RepositoryServiceGrpcImpl struct {
	domain.UnimplementedRepositoryServiceServer
}

//NewRepositoryServiceGrpcImpl returns the pointer to the implementation.
func NewRepositoryServiceGrpcImpl() *RepositoryServiceGrpcImpl {
	return &RepositoryServiceGrpcImpl{}
}

//Add function implementation of GRPC Service.
func (serviceImpl *RepositoryServiceGrpcImpl) Add(ctx context.Context, in *domain.Repository) (*domain.AddRepositoryResponse, error) {
	log.Println("Received request for adding repository with id " + strconv.FormatInt(in.Id, 10))

	//Logic to persist to database or storage.
	log.Println("Repository persisted to the storage")

	return &domain.AddRepositoryResponse{
		AddedRepository: in,
		Error:           nil,
	}, nil
}
