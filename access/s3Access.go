package access

import (
	"adnw.cloudent-tmi.com/arch/samples/aws-sample/s3mightiness"
	"github.com/aws/aws-sdk-go/aws/session"
)

var prop s3mightiness.Prop

//Init s3へのアクセス
func S3Init() *session.Session {
	s3Config := s3mightiness.Prop{
		AccessKey:       "hogehoge",
		SecretKey:       "hogehoge",
		ServiceEndpoint: "http://127.0.0.1:9001",
		Region:          "ut1",
	}
	s3Client := session.New(s3Config.S3Access())

	return s3Client
}
