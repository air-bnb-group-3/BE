package awss3

import (
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/gommon/log"
)

func InitS3(key, secret, region string) *session.Session {
	conect, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				key, secret, "o3T3ozzKzrdIfiDTPMVFMgP7NWfpFm75hxtX2Cww",
			),
		},
	)
	if err != nil {
		log.Error("aws S3 Config error", err)
	}

	return conect
}

func Upload(sess *session.Session, file multipart.FileHeader) string {
	manager := s3manager.NewUploader(sess)
	src, err := file.Open()
	if err != nil {
		log.Info(err)
	}

	defer src.Close()
	buffer := make([]byte, file.Size)
	src.Read(buffer)
	body, _ := file.Open()

	res, err := manager.Upload(
		&s3manager.UploadInput{
			Bucket:      aws.String("airbnb-app"),
			ACL:         aws.String("public-read"),
			ContentType: aws.String(http.DetectContentType(buffer)),
			Key:         aws.String(file.Filename),
			Body:        body,
		},
	)
	if err != nil {
		log.Info(res)
		log.Error("Upload error : ", err)
	}

	return ""
}
