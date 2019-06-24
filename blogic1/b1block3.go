package blogic1

import (
	"adnw.cloudent-tmi.com/arch/samples/samplePJ/access"
	"adnw.cloudent-tmi.com/arch/samples/samplePJ/model"
)

//DBからユーザリスト取得
func SearchUserDataInUsers() []model.User {
	db := access.DBInit()
	users := []model.User{}
	db.Raw("SELECT * FROM user where id in (1)").Scan(&users)
	return users
}
