package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Actions stores all actions/variables for the project
type Actions struct {
	session    *session.Session
	s3         *s3.S3
	s3Uploader *s3manager.Uploader
	buckerName string
}

// SetUpS3 will intance all objects needed for the project
func SetUpS3(bucketName string) (Actions, error) {
	session, err := session.NewSession()
	if err != nil {
		return Actions{}, err
	}
	s3 := s3.New(session)
	s3Uploader := s3manager.NewUploader(session)
	return Actions{session, s3, s3Uploader, bucketName}, nil
}

// UploadFile uploads a file to S3 and returns the remote URL of the uploaded file
func (a *Actions) UploadFile(filename string, reader io.Reader) (string, error) {
	resp, err := a.s3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(a.buckerName),
		Key:    aws.String(filename),
		Body:   reader,
	})
	if err != nil {
		return "", err
	}
	return resp.Location, nil
}
