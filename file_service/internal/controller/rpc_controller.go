package controller

import (
	"context"
	"errors"
	"github.com/catness812/e-petitions-project/file_service/internal/model"
	"github.com/catness812/e-petitions-project/file_service/internal/pb"
)

type IFileService interface {
	SendFile(file *model.File) error
	SendUserPic(file *model.File) error
	GetUserPic(userID uint32) (*model.File, error)
	GetFile(fileID uint32) (*model.File, error)
	GetPetitionFiles(petitionID uint32) ([]model.File, error)
}

type FileRPCServer struct {
	fileSvc IFileService
}

func NewFileRPCServer(fileSvc IFileService) *FileRPCServer {
	return &FileRPCServer{fileSvc: fileSvc}
}

func (s *FileRPCServer) UploadFile(ctx context.Context, req *pb.FileRequest) (*pb.FileResponse, error) {
	if err := validateFileType(req.Type, "petition"); err != nil {
		return nil, err
	}
	file := &model.File{
		UserID:     req.Uid,
		PetitionID: req.Pid,
		Data:       req.FileData,
		Type:       req.Type,
	}
	if err := s.fileSvc.SendFile(file); err != nil {
		return nil, err
	}
	return &pb.FileResponse{Message: "Successfully uploaded file"}, nil
}

func (s *FileRPCServer) UploadUserProfilePicture(ctx context.Context, req *pb.FileRequest) (*pb.FileResponse, error) {
	if err := validateFileType(req.Type, "profilepic"); err != nil {
		return nil, err
	}
	file := &model.File{
		UserID: req.Uid,
		Data:   req.FileData,
		Type:   req.Type,
	}
	if err := s.fileSvc.SendUserPic(file); err != nil {
		return nil, err
	}
	return &pb.FileResponse{Message: "Profile photo uploaded successfully"}, nil
}

func (s *FileRPCServer) DownloadFile(ctx context.Context, req *pb.DownloadFileRequest) (*pb.FileRequest, error) {
	file, err := s.fileSvc.GetFile(req.FileID)
	if err != nil {
		return nil, err
	}
	return &pb.FileRequest{
		FileData: file.Data,
		Uid:      file.UserID,
		Pid:      file.PetitionID,
		Type:     file.Type,
	}, nil
}

func (s *FileRPCServer) DownloadUserProfilePic(ctx context.Context, req *pb.DownloadFileRequest) (*pb.FileRequest, error) {
	pic, err := s.fileSvc.GetUserPic(req.UserID)
	if err != nil {
		return nil, err
	}
	return &pb.FileRequest{
		FileData: pic.Data,
		Uid:      pic.UserID,
		Pid:      pic.PetitionID,
		Type:     pic.Type,
	}, nil
}

func (s *FileRPCServer) DownloadPetitionFiles(ctx context.Context, req *pb.DownloadFileRequest) ([]*pb.FileRequest, error) {
	files, err := s.fileSvc.GetPetitionFiles(req.PetitionID)
	if err != nil {
		return nil, err
	}

	var fileRequests []*pb.FileRequest
	for _, file := range files {
		fileRequest := &pb.FileRequest{
			FileData: file.Data,
			Uid:      file.UserID,
			Pid:      file.PetitionID,
			Type:     file.Type,
		}
		fileRequests = append(fileRequests, fileRequest)
	}

	return fileRequests, nil
}

func validateFileType(extension, purpose string) error {
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
