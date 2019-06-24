package main

import (
	"fmt"

	"adnw.cloudent-tmi.com/arch/samples/samplePJ/blogic1"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//日時取得
	fmt.Println(blogic1.GetNow())

	//s3アクセス
	// s3Access := access.S3Init()

	//userListを取得
	// userList, err := blogic1.GetUserListFromS3(s3Access, "test", "user_list.csv")
	// fmt.Println(userList)
	// fmt.Println(err)

	fmt.Println(blogic1.SearchUserDataInUsers())

}
