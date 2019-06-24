package blogic1

import (
	"encoding/csv"
	"os"

	"adnw.cloudent-tmi.com/arch/samples/aws-sample/s3mightiness"
	"github.com/aws/aws-sdk-go/aws/session"
)

// //s３からユーザリストをダウンロード
func GetUserListFromS3(s3Client *session.Session, bucket string, key string) ([]string, error) {

	err := s3mightiness.S3Download(s3Client, bucket, key, "test.csv")

	if err != nil {
		return nil, err
	}

	file, err := os.Open("test.csv")
	if err != nil {
		return nil, err
	}

	//読み込んだあとローカルファイル削除
	defer os.Remove("test.csv")

	defer file.Close()

	reader := csv.NewReader(file)
	var line []string

	line, err = reader.Read()
	if err != nil {
		return nil, err
	}
	return line, nil

}
