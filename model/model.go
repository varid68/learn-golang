package model

type User struct {
	IdUser   int64  `gorm:"primary_key" json:"id_user"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
