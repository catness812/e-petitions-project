package controller

import (
	"context"
	"github.com/catness812/e-petitions-project/file_service/internal/pb"
)

type IFileService interface {
	UploadFile(fileID int, data []byte) error
}

type FileRPCServer struct {
	fileSvc IFileService
}

func NewFileRPCServer(fileSvc IFileService) *FileRPCServer {
	return &FileRPCServer{fileSvc: fileSvc}
}

func (s *FileRPCServer) UploadFile(ctx context.Context, req *pb.FileRequest) (*pb.FileResponse, error) {
	return nil, nil
}