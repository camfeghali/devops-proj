package database

import (

	"hackernews-clone/graph/model"

)

func AllUsers() []model.User {
	
	var users []model.User
	Db.Find(&users)

	return users
}