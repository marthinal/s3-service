package s3_service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/net/context"
	"io"
	"os"
	"time"
)

type FileSystem interface {
	Save(keyName string, file io.ReadSeeker) error
}

type FS struct {
	svc *s3.S3
	bucket string
	timeout time.Duration
}

func (fs FS) Save(keyName string, file io.ReadSeeker) error {
	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.
	ctx := context.Background()
	//var cancelFn func()
	//if timeout > 0 {
	//	ctx, cancelFn = context.WithTimeout(ctx, timeout)
	//}
	//if cancelFn != nil {
	//	defer cancelFn()
	//}

	_, err := fs.svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(fs.bucket),
		Key:    aws.String(keyName),
		Body:   file,
	})

	if err != nil {
		//if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
		//	// If the SDK can determine the request or retry delay was canceled
		//	// by a context the CanceledErrorCode error code will be returned.
		//	fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		//} else {
		//	fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		//}
		fmt.Fprintf(os.Stderr, "Failed to upload object, %v\n", err)
		//os.Exit(1)
		return err
	}

	fmt.Printf("Successfully uploaded file to %s/%s\n", fs.bucket, keyName)
	return nil
}