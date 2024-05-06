package service

import (
	"context"

	pb "kratos/api/dep"
)

type DepServiceService struct {
	pb.UnimplementedDepServiceServer
}

func NewDepServiceService() *DepServiceService {
	return &DepServiceService{}
}

func (s *DepServiceService) CreateDep(ctx context.Context, req *pb.CreateDepRequest) (*pb.CreateDepReply, error) {
	return &pb.CreateDepReply{}, nil
}
func (s *DepServiceService) UpdateDep(ctx context.Context, req *pb.UpdateDepRequest) (*pb.UpdateDepReply, error) {
	return &pb.UpdateDepReply{}, nil
}
func (s *DepServiceService) DeleteDep(ctx context.Context, req *pb.DeleteDepRequest) (*pb.DeleteDepReply, error) {
	return &pb.DeleteDepReply{}, nil
}
func (s *DepServiceService) GetDep(ctx context.Context, req *pb.GetDepRequest) (*pb.GetDepReply, error) {
	return &pb.GetDepReply{}, nil
}
func (s *DepServiceService) ListDep(ctx context.Context, req *pb.ListDepRequest) (*pb.ListDepReply, error) {
	return &pb.ListDepReply{}, nil
}
