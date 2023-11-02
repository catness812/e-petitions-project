package controller

import (
	"context"
	"errors"
	"github.com/catness812/e-petitions-project/file_service/internal/pb"
)

type IFileService interface {
	UploadFileData(data []byte, fileType, id string) error
}

type FileRPCServer struct {
	fileSvc IFileService
}

func NewFileRPCServer(fileSvc IFileService) *FileRPCServer {
	return &FileRPCServer{fileSvc: fileSvc}
}

func (s *FileRPCServer) UploadFile(ctx context.Context, req *pb.FileRequest) (*pb.FileResponse, error) {
	if err := validateFileType(req.Type); err != nil {
		return nil, err
	}
	if err := s.fileSvc.UploadFileData(req.FileData, req.Type, req.Uid); err != nil {
		return nil, err
	}
	return &pb.FileResponse{Message: "Successfully uploaded file"}, nil
}

func validateFileType(extension string) error {
	switch extension {
	case "png":
		return nil
	case "jpg":
		return nil
	case "jpeg":
		return nil
	case "mp4":
		return nil
	default:
		return errors.New("invalid file type")
	}
}
