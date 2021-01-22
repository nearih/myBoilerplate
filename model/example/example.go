package example

type Example struct {
	ID       string `json:"id" gorm:"PRIMARY_KEY;column:documentID;type:varchar(70) CHARACTER SET ascii"`
	Username string `json:"username"`
	Password string `json:"password"`
}
