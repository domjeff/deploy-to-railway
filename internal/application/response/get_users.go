package response

type GetUsers struct {
	Users []GetUsersUser `json:"users"`
}

type GetUsersUser struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
