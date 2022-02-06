package filestorage

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"maranatha_web/internal/config"
)

var Client *minio.Client

type MinioDao interface {
	UploadFile(objectName string, fileBuffer multipart.File, fileSize int64, contentType string) (minio.UploadInfo, error)
}

type MinioRepo struct {
	App          *config.AppConfig
	MinioStorage *minio.Client
}

func NewMinoRepo(conn *minio.Client, a *config.AppConfig) MinioDao {
	return &MinioRepo{
		App:          a,
		MinioStorage: conn,
	}
}

// GetMinioConnection MinioConnection func for opening filestorage connection.
func GetMinioConnection() (*minio.Client, string, error) {

	err := godotenv.Load()
	if err != nil {

		log.Println(err)
		log.Fatal("Error  .env file")
	}
	minioAccessKeyID := os.Getenv("MinioAccessKeyID")
	minioEndPoint := os.Getenv("MinioEndPoint")
	minioAccessKey := os.Getenv("MinioAccessKey")
	minioBucketName := os.Getenv("MinioBucketName")

	log.Println(minioEndPoint)
	ctx := context.Background()

	// Initialize filestorage client object.
	minioClient, errInit := minio.New(minioEndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKeyID, minioAccessKey, ""),
		Secure: false,
	})
	if errInit != nil {
		log.Fatalln(errInit)
	}
	// Make a new bucket called dev-filestorage.
	bucketName := minioBucketName
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket %s is already created ...Skipping creation", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	Client = minioClient

	return minioClient, bucketName, errInit
}

//
//type fileStorageServiceInterface interface {
//	UploadFile(objectName string, fileBuffer multipart.File, fileSize int64, contentType string) (minio.UploadInfo, error)
//}

//TODO:Add this upload file to services
func (s *MinioRepo) UploadFile(objectName string, fileBuffer multipart.File, fileSize int64, contentType string) (minio.UploadInfo, error) {
	ctx := context.Background()
	// Upload the zip file with PutObject
	info, err := s.MinioStorage.PutObject(ctx, "clj", objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return info, err
	}
	return info, nil
}
