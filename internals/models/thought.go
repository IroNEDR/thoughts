package models

type Thought struct {
	CreatedBy   User       `json:"created_by"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Categories  []Category `json:"categories"`
	Tags        []string   `json:"tags,omitempty"`
	Comments    []Comments `json:"comments"`
	Public      bool       `json:"public"`
}
