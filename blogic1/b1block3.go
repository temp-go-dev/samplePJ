package blogic1

import (
	"adnw.cloudent-tmi.com/arch/samples/samplePJ/access"
	"adnw.cloudent-tmi.com/arch/samples/samplePJ/model"
)

//DBからユーザリスト取得
func SearchUserDataInUsers(userId []string) []model.User {
	db := access.DBInit()

	users := []model.User{}

	//ユーザIDごとDBでユーザ情報を探す
	for _, userid := range userId {
		user := model.User{}
		db.Raw("SELECT * FROM user where id = ?", userid).Scan(&user)
		if user.ID != "" {
			users = append(users, user)
		}
	}

	return users
}
