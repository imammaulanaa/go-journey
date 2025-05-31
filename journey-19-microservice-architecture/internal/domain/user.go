package domain

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	GetByID(id string) (*User, error)
}
