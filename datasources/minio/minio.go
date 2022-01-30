package minio

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Client *minio.Client

// GetMinioConnection MinioConnection func for opening minio connection.
func GetMinioConnection() (*minio.Client, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	minioAccessKeyID := os.Getenv("MinioAccessKeyID")
	minioEndPoint := os.Getenv("MinioEndPoint")
	minioAccessKey := os.Getenv("MinioAccessKey")
	minioBucketName := os.Getenv("MinioBucketName")

	log.Println(minioEndPoint)
	ctx := context.Background()

	// Initialize minio client object.
	minioClient, errInit := minio.New(minioEndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKeyID, minioAccessKey, ""),
		Secure: false,
	})
	if errInit != nil {
		log.Fatalln(errInit)
	}
	// Make a new bucket called dev-minio.
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

	return minioClient, errInit
}
