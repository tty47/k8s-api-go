package user

// User struct use to define the fields
type User struct {
	ID       int32  `json:"id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}

// type users is a users array
type Users []User
