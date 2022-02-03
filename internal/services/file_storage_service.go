package services

import (
	"context"
	"maranatha_web/internal/datasources/filestorage"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

var (
	FileStorageService fileStorageServiceInterface = &fileStorageService{
		bucketName: "mono",
	}
)

type fileStorageService struct {
	bucketName string
}

type fileStorageServiceInterface interface {
	UploadFile(objectName string, fileBuffer multipart.File, fileSize int64, contentType string) (minio.UploadInfo, error)
}

func (s *fileStorageService) UploadFile(objectName string, fileBuffer multipart.File, fileSize int64, contentType string) (minio.UploadInfo, error) {
	ctx := context.Background()
	// Upload the zip file with PutObject
	info, err := filestorage.Client.PutObject(ctx, s.bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return info, err
	}
	return info, nil
}
