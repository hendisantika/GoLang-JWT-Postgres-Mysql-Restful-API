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

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}
