package models

// List list struct
type List struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// User user struct
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TableName return a table name
func (b *List) TableName() string {
	return "list"
}

// TableName return a table name
func (b *User) TableName() string {
	return "User"
}
