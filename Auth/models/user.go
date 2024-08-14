package models

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

type SignIn struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type UpdatePassword struct {
	OldPassword  string `json:"oldpassword"`
	NewPassword1 string `json:"newpassword1"`
	NewPassword2 string `json:"newpassword2"`
}
