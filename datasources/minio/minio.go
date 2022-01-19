package minio

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

// MinioConnection func for opening minio connection.
func MinioConnection() (*minio.Client, error) {
	ctx := context.Background()
	endpoint := "localhost:9000"
	accessKeyID := "admin"
	secretAccessKey := "password"
	useSSL := false
	// Initialize minio client object.
	minioClient, errInit := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if errInit != nil {
		log.Fatalln(errInit)
	}

	// Make a new bucket called dev-minio.
	bucketName := "mono"
	location := "us-east-1"

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
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
	MinioClient = minioClient

	return minioClient, errInit
}
