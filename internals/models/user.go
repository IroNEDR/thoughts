package models

type User struct {
	Username    string   `json:"username"`
	Description string   `json:"description"`
	Interests   []string `json:"interests"`
}
