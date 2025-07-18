package services

import (
	"hris_backend/pkg/r2"
	"mime/multipart"
)

type UploadService interface {
	UploadImage(fileHeader *multipart.FileHeader, folder string) (*r2.UploadResult, error)
	DeleteImage(fileName string) error
}

type uploadService struct {
	Uploader *r2.Uploader
}

func NewUploadService(uploader *r2.Uploader) UploadService {
	return &uploadService{Uploader: uploader}
}

func (s *uploadService) UploadImage(fileHeader *multipart.FileHeader, folder string) (*r2.UploadResult, error) {
	return s.Uploader.Upload(fileHeader, folder)
}

func (s *uploadService) DeleteImage(fileName string) error {
	return s.Uploader.Delete(fileName)
}
