package provider

// OrderItem -
type Dev struct {
	Engineer Engineer_Api `json:"engineer"`
	Id       string       `json:"id"`
	Name     string       `json:"name"`
}

type Engineer_Api struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Email string `json:"email"`
}
