package models

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginCred struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DbLoginCred struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}
