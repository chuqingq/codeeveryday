package util

import (
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// sudo docker pull bitnami/minio:2021.10.23
// sudo docker run -p 9000:9000 -v (pwd)/data:/data -e "MINIO_ROOT_USER=minioadmin" -e "MINIO_ROOT_PASSWORD=minioadmin" -d bitnami/minio:2021.10.23 minio server /data
// curl -F "fileupload=@server.go" http://127.0.0.1:9090/api/file/upload -v

const endpoint = "127.0.0.1:9000"
const accessKeyID = "minioadmin"
const secretAccessKey = "minioadmin"
const useSSL = false
const bucketName = "image-check-bucket"
const location = "image-check-location"

var minioClient *minio.Client

// InitFile 初始化文件服务
func InitFile() error {
	// Initialize minio client object.
	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return err
	}

	// Make a new bucket called mymusic.

	ctx := context.Background()
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	return nil
}

// PutFile 上传文件
func PutFile(fileid string, file io.Reader, filesize int64) error {
	_, err := minioClient.PutObject(context.Background(), bucketName, fileid, file, filesize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	return err
}

// GetFile 下载文件
func GetFile(fileid string) (io.ReadCloser, error) {
	reader, err := minioClient.GetObject(context.Background(), bucketName, fileid, minio.GetObjectOptions{})
	return reader, err
}
