package services

// import (
// 	"context"
// 	"mime/multipart"
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// 

// 	"github.com/minio/minio-go/v7"
// 	// "github.com/minio/minio-go/v7"

// 	// "maranatha_web/datasources/minio" //nolint:goimports
// )

// var (
// 	MinioService minioServiceInterface = &minioService{
// 		bucketName: "mono",
// 	}
// )

// type minioService struct {
// 	bucketName string
// }

// type minioServiceInterface interface {
// 	UploadFile(objectName string, fileBuffer multipart.File, fileSize int64, contentType string) (minio.UploadInfo, error)
// }

// func (s *minioService) UploadFile(objectName string, fileBuffer multipart.File, fileSize int64, contentType string) (minio.UploadInfo, error) {
// 	ctx := context.Background()
// 	// Upload the zip file with PutObject
// 	info, err := minioClient.Client.PutObject(ctx, s.bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
// 	if err != nil {
// 		return info, err
// 	}
// 	return info, nil
// }
