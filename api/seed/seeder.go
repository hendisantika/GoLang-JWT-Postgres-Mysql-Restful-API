package seed

import (
	"github.com/victorsteven/fullstack/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Itadori Yuji",
		Email:    "yuji@yopmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Gojo satoru",
		Email:    "gojo@yopmail.com",
		Password: "password",
	},
}
