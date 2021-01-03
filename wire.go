//+build wireinject

package s3_service

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/wire"
	"os"
)

func GetFS() FS {
	// Load the session.
	s3Session := session.Must(session.NewSession())
	svc := s3.New(s3Session)

	return FS {
		svc: svc,
		bucket: os.Getenv("AWS_BUCKET_NAME"),
	}
}

func InitializeFS() FS {
	wire.Build(GetFS)
	return FS{}
}
