package main

import (
	"fmt"

	"adnw.cloudent-tmi.com/arch/samples/samplePJ/access"
	"adnw.cloudent-tmi.com/arch/samples/samplePJ/blogic1"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//日時取得
	fmt.Println(blogic1.GetNow())

	//s3アクセス
	s3Access := access.S3Init()

	//userIDList ユーザIdリストを取得
	userIDList, err := blogic1.GetUserListFromS3(s3Access, "test", "user_list.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(userIDList)
	fmt.Println(err)

	//DBからユーザ情報を探す
	userList := blogic1.SearchUserDataInUsers(userIDList)
	fmt.Println(userList)

}
