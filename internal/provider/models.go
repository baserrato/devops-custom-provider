package provider

type Dev struct {
	Engineers map[string]Engineer_Api `json:"engineers"`
	Id        string                  `json:"id"`
	Name      string                  `json:"name"`
}

type Ops struct {
	Engineers map[string]Engineer_Api `json:"engineers"`
	Id        string                  `json:"id"`
	Name      string                  `json:"name"`
}

type Engineer_Api struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Email string `json:"email"`
}
